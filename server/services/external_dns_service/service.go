package external_dns_service

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/server/services/services"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"sigs.k8s.io/external-dns/endpoint"
)

type Service struct {
	natIpNameTemplate      string
	natIpLabelsTemplate    string
	endpointNameTemplate   string
	endpointLabelsTemplate string

	endpointCache []*endpoint.Endpoint
	interval      time.Duration
	nextRunAt     time.Time

	natIpService services.NamespacedService[*documents.NatIp]
}

func NewService(natIpNameTemplate, natIpLabelsTemplate, endpointNameTemplate, endpointLabelsTemplate string, interval time.Duration, natIpService services.NamespacedService[*documents.NatIp]) *Service {
	zap.S().Infow("Initialize External-Dns Service with", "interval", interval.String())
	return &Service{
		natIpNameTemplate:      natIpNameTemplate,
		natIpLabelsTemplate:    natIpLabelsTemplate,
		endpointNameTemplate:   endpointNameTemplate,
		endpointLabelsTemplate: endpointLabelsTemplate,
		interval:               interval,
		natIpService:           natIpService,
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

	byLabel := map[LabelKey][]*documents.NatIp{}
	for _, natIp := range natIps {
		for label := range natIp.Metadata.Labels {
			key := LabelKey{Label: label, Namespace: natIp.Metadata.Namespace}
			byLabel[key] = append(byLabel[key], natIp)
		}
	}

	natIpNameTmpl, err := NewTemplateHelper(s.natIpNameTemplate)
	if err != nil {
		return err
	}
	var endpoints []*endpoint.Endpoint
	for _, natIp := range natIps {
		srvTargets := lo.Map[string, string](natIp.Spec.Cidrs, func(s string, _ int) string {
			return fmt.Sprintf("1 1 0 %s", s)
		})
		srvTargets = lo.Uniq[string](srvTargets)

		dnsName, err := natIpNameTmpl.Execute(natIp.Metadata.Name, natIp.Metadata.Namespace)
		if err != nil {
			continue
		}

		endpoints = append(endpoints, &endpoint.Endpoint{
			DNSName:    dnsName,
			Targets:    srvTargets,
			RecordType: recordType,
			RecordTTL:  recordTTL,
		})
	}

	natIpLabelsTmpl, err := NewTemplateHelper(s.natIpLabelsTemplate)
	if err != nil {
		return err
	}
	for key, ips := range byLabel {
		srvTargets := lo.FlatMap[*documents.NatIp, string](ips, func(ip *documents.NatIp, _ int) []string {
			return lo.Map[string, string](ip.Spec.Cidrs, func(s string, _ int) string {
				return fmt.Sprintf("1 1 0 %s", s)
			})
		})
		srvTargets = lo.Uniq[string](srvTargets)

		dnsName, err := natIpLabelsTmpl.Execute(key.Label, key.Namespace)
		if err != nil {
			continue
		}

		endpoints = append(endpoints, &endpoint.Endpoint{
			DNSName:    dnsName,
			Targets:    srvTargets,
			RecordType: recordType,
			RecordTTL:  recordTTL,
		})
	}

	s.endpointCache = endpoints
	zap.S().Infow("External-Dns Finish SyncCache", "# of NatIps", len(natIps), "# of Labels", len(byLabel), "# of Records", len(endpoints))
	return nil
}

type LabelKey struct {
	Label     string
	Namespace string
}

type TemplateHelper struct {
	tmpl    *template.Template
	tmplStr string
}

func NewTemplateHelper(tmplStr string) (*TemplateHelper, error) {
	tmpl, err := template.New("gotmpl").Funcs(sprig.TxtFuncMap()).Parse(tmplStr)
	if err != nil {
		zap.S().Infow("External-Dns DnsName Parse Template Failed", "template", tmplStr, "error", err)
		return nil, err
	}
	return &TemplateHelper{tmpl: tmpl, tmplStr: tmplStr}, nil
}

func (h *TemplateHelper) Execute(name, namespace string) (string, error) {
	buf := &bytes.Buffer{}
	values := &TemplateValues{Name: name, Namespace: namespace}
	err := h.tmpl.Execute(buf, values)
	if err != nil {
		zap.S().Infow("External-Dns DnsName Execute Template Failed, This error would skip record template", "template", h.tmplStr, "values", values, "error", err)
		return "", err
	}
	return buf.String(), nil
}

type TemplateValues struct {
	Name      string
	Namespace string
}

const recordTTL = endpoint.TTL(30)
const recordType = "SRV"

// SRV Record Example
// [priority] [weight] [port] [server host name]
// 1 1 8080 example.com

func (s *Service) Records() ([]*endpoint.Endpoint, error) {
	return s.endpointCache, nil
}
