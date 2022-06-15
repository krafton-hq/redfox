package natip_service

import (
	"fmt"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/apis/idl_common"
)

type Service struct {
	local map[string]map[string]*documents.NatIp
}

func (s *Service) Get(namespace string, name string) (*documents.NatIp, error) {
	ns, exist := s.local[namespace]
	if !exist {
		return nil, fmt.Errorf("namespace not eixst")
	}

	if nat, exist := ns[name]; exist {
		return nat, nil
	} else {
		return nil, fmt.Errorf("natip not eixst")
	}
}

func (s *Service) List(labelSelectors map[string]string) []*documents.NatIp {
	var ret []*documents.NatIp
	for _, namespaced := range s.local {
		for _, natIp := range namespaced {
			if containsLabels(natIp.Metadata, labelSelectors) {
				ret = append(ret, natIp)
			}
		}
	}
	return ret
}

func (s *Service) ListNamespaced(namespace string, labelSelectors map[string]string) ([]*documents.NatIp, error) {
	ns, exist := s.local[namespace]
	if !exist {
		return nil, fmt.Errorf("namespace not eixst")
	}

	var ret []*documents.NatIp
	for _, natIp := range ns {
		if containsLabels(natIp.Metadata, labelSelectors) {
			ret = append(ret, natIp)
		}
	}
	return ret, nil
}

func containsLabels(metadata *idl_common.ObjectMeta, labels map[string]string) bool {
	for key, value := range labels {
		if !containsLabel(metadata, key, value) {
			return false
		}
	}
	return true
}

func containsLabel(metadata *idl_common.ObjectMeta, tkey, tvalue string) bool {
	for lkey, lvalue := range metadata.Labels {
		if lkey == tkey {
			if tvalue == "" {
				return true
			}
			if lvalue == tvalue {
				return true
			}
		}
	}
	return false
}

func (s *Service) Create(natIp *documents.NatIp) error {
	ns, exist := s.local[natIp.Metadata.Namespace]
	if !exist {
		return fmt.Errorf("namespace not eixst")
	}

	name := natIp.Metadata.Name
	if _, exist := ns[name]; exist {
		return fmt.Errorf("natip alreay eixst")
	}
	ns[name] = natIp
	return nil
}

func (s *Service) Update(natIp *documents.NatIp) error {
	ns, exist := s.local[natIp.Metadata.Namespace]
	if !exist {
		return fmt.Errorf("namespace not eixst")
	}

	name := natIp.Metadata.Name
	if _, exist := ns[name]; !exist {
		return fmt.Errorf("natip not eixst")
	}
	ns[name] = natIp
	return nil
}

func (s *Service) Delete(namespace string, name string) error {
	ns, exist := s.local[namespace]
	if !exist {
		return fmt.Errorf("namespace not eixst")
	}

	if _, exist := ns[name]; !exist {
		return fmt.Errorf("natip not eixst")
	}
	delete(ns, name)
	return nil
}
