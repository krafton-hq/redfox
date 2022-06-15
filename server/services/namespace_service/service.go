package namespace_service

import (
	"fmt"

	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/apis/namespaces"
)

type Service struct {
	local map[string]*namespaces.Namespace
}

func NewService() *Service {
	return &Service{local: map[string]*namespaces.Namespace{}}
}

func (s *Service) GetNamespace(name string) (*namespaces.Namespace, error) {
	if ns, exist := s.local[name]; exist {
		return ns, nil
	} else {
		return nil, fmt.Errorf("namespace not eixsts")
	}
}

func (s *Service) ListNamespaces(labelSelectors map[string]string) []*namespaces.Namespace {
	var ret []*namespaces.Namespace
	for _, namespace := range s.local {
		if containsLabels(namespace.Metadata, labelSelectors) {
			ret = append(ret, namespace)
		}
	}
	return ret
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

func (s *Service) Create(namespace *namespaces.Namespace) error {
	name := namespace.Metadata.Name
	if _, exist := s.local[name]; exist {
		return fmt.Errorf("namespace Already Exists")
	}

	s.local[name] = namespace
	return nil
}

func (s *Service) Update(namespace *namespaces.Namespace) error {
	name := namespace.Metadata.Name
	if _, exist := s.local[name]; !exist {
		return fmt.Errorf("namespace not eixsts")
	}

	s.local[name] = namespace
	return nil
}

func (s *Service) Delete(name string) error {
	if _, exist := s.local[name]; !exist {
		return fmt.Errorf("namespace not eixsts")
	}

	delete(s.local, name)
	return nil
}
