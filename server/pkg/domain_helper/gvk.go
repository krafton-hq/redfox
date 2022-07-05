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

func ParseGvkName(name string) (kind string, group string, err error) {
	args := strings.SplitN(name, ".", 1)
	if len(args) < 1 {
		return "", "", errors.NewInvalidArguments("GvkName should have at least 1 dot '.'")
	}
	kind = args[0]
	group = args[1]

	if errs := validation.IsGroup(group); len(errs) > 0 {
		return "", "", errors.NewInvalidField("GvkName", "RFC1123 Dns Label/Version", kind)
	}
	return
}
