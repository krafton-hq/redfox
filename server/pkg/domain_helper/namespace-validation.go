package domain_helper

import (
	"github.com/krafton-hq/red-fox/apis/namespaces"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/krafton-hq/red-fox/server/pkg/validation"
)

func ValidationNamespaceSpec(spec *namespaces.NamespaceSpec) error {
	if spec == nil {
		return nil
	}

	for _, object := range spec.ApiObjects {
		if errs := validation.IsKind(object.Kind); len(errs) > 0 {
			return errors.NewInvalidField("spec.apiObjects.kind", "kind value", object.Kind)
		}
		if errs := validation.IsGroup(object.Group); len(errs) > 0 {
			return errors.NewInvalidField("spec.apiObjects.group", "group value", object.Group)
		}
		if errs := validation.IsVersion(object.Version); len(errs) > 0 {
			return errors.NewInvalidField("spec.apiObjects.version", "version value", object.Version)
		}
	}
	return nil
}
