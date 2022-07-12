package domain_helper

import (
	"strings"

	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/krafton-hq/red-fox/server/pkg/validation"
)

func GetGvkName(gvk *idl_common.GroupVersionKindSpec) string {
	return strings.ToLower(gvk.Kind + "." + gvk.Group)
}

func ParseGvkName(name string) (lowerKind string, group string, err error) {
	lowerKind, group, found := strings.Cut(name, ".")
	if !found {
		return "", "", errors.NewInvalidArguments("GvkName should have at least 1 dot '.'")
	}

	if errs := validation.IsGroup(group); len(errs) > 0 {
		return "", "", errors.NewInvalidField("GvkName", "RFC1123 Dns Label/Version", lowerKind)
	}
	return
}

func CreateGvkFromMetadatable(m Metadatable) (*idl_common.GroupVersionKindSpec, error) {
	group, version, found := strings.Cut(m.GetApiVersion(), "/")
	if !found {
		return nil, errors.NewInvalidField("apiVersion", "Should have one '/'", m.GetApiVersion())
	}
	return &idl_common.GroupVersionKindSpec{
		Group:   group,
		Version: version,
		Kind:    m.GetKind(),
	}, nil
}

func EqualsGvk(g1 *idl_common.GroupVersionKindSpec, g2 *idl_common.GroupVersionKindSpec) bool {
	return g1.Group == g2.Group && g1.Version == g2.Version && g1.Kind == g2.Kind
}
