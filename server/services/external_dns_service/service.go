package external_dns_service

import (
	"context"
	"fmt"
	"time"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/server/services/services"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"sigs.k8s.io/external-dns/endpoint"
)

type Service struct {
	natIpsByName     map[string]*documents.NatIp
	natIpsByLabel    map[string][]*documents.NatIp
	endpointsByName  map[string]*documents.Endpoint
	endpointsByLabel map[string][]*documents.Endpoint

	natIpDomain    string
	endpointDomain string

	interval  time.Duration
	nextRunAt time.Time

	natIpService services.NamespacedService[*documents.NatIp]
}

func NewService(natIpDomain, endpointDomain string, interval time.Duration, natIpService services.NamespacedService[*documents.NatIp]) *Service {
	zap.S().Infow("Initialize External-Dns Service with", "natIpDomain", natIpDomain, "interval", interval.String())

	return &Service{
		natIpsByName:     map[string]*documents.NatIp{},
		natIpsByLabel:    map[string][]*documents.NatIp{},
		endpointsByName:  map[string]*documents.Endpoint{},
		endpointsByLabel: map[string][]*documents.Endpoint{},
		natIpDomain:      natIpDomain,
		endpointDomain:   endpointDomain,
		interval:         interval,
		natIpService:     natIpService,
	}
}

func (s *Service) shouldRunOnce(now time.Time) bool {
	if now.Before(s.nextRunAt) {
		return false
	}
	s.nextRunAt = now.Add(s.interval)
	return true
}

func (s *Service) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		if s.shouldRunOnce(time.Now()) {
			if err := s.SyncCache(ctx); err != nil {
				zap.S().Errorw("External-Dns Sync Inmemory Cache Failed", "error", err)
			}
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			zap.S().Info("Terminating main controller loop")
			return
		}
	}
}

func (s *Service) SyncCache(ctx context.Context) error {
	zap.S().Info("External-Dns Start SyncCache")
	natIps, err := s.natIpService.List(ctx, nil)
	if err != nil {
		return err
	}

	byName := map[string]*documents.NatIp{}
	byLabel := map[string][]*documents.NatIp{}

	for _, natIp := range natIps {
		byName[natIp.Metadata.Name] = natIp
		for key := range natIp.Metadata.Labels {
			byLabel[key] = append(byLabel[key], natIp)
		}
	}

	s.natIpsByName = byName
	s.natIpsByLabel = byLabel
	zap.S().Infof("External-Dns Finish SyncCache, Synced %d Records", len(natIps))
	return nil
}

const recordTTL = endpoint.TTL(60)
const recordType = "SRV"

// SRV Record Example
// [priority] [weight] [port] [server host name]
// 1 1 8080 example.com

func (s *Service) Records() ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint
	for name, ip := range s.natIpsByName {
		srvTargets := lo.Map[string, string](ip.Spec.Cidrs, func(s string, _ int) string {
			return fmt.Sprintf("1 1 0 %s", s)
		})

		endpoints = append(endpoints, &endpoint.Endpoint{
			DNSName:    fmt.Sprintf("%s.nat-name.%s", name, s.natIpDomain),
			Targets:    srvTargets,
			RecordType: recordType,
			RecordTTL:  recordTTL,
		})
	}

	for label, ips := range s.natIpsByLabel {

		srvTargets := lo.FlatMap[*documents.NatIp, string](ips, func(ip *documents.NatIp, _ int) []string {
			return lo.Map[string, string](ip.Spec.Cidrs, func(s string, _ int) string {
				return fmt.Sprintf("1 1 0 %s", s)
			})
		})

		srvTargets = lo.Uniq[string](srvTargets)

		endpoints = append(endpoints, &endpoint.Endpoint{
			DNSName:    fmt.Sprintf("%s.nat-label.%s", label, s.natIpDomain),
			Targets:    srvTargets,
			RecordType: recordType,
			RecordTTL:  recordTTL,
		})
	}

	return endpoints, nil
}
