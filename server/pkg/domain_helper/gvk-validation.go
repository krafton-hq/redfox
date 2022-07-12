package domain_helper

import (
	"github.com/krafton-hq/red-fox/apis/crds"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/krafton-hq/red-fox/server/pkg/validation"
)

func ValidationCrdSpec(spec *crds.CustomResourceDefinitionSpec) error {
	if spec == nil {
		return errors.NewInvalidField("spec", "object", "null")
	}

	if spec.Gvk == nil {
		return errors.NewInvalidField("spec.detail", "object", "null")
	}

	if errs := validation.IsKind(spec.Gvk.Kind); len(errs) > 0 {
		return errors.NewInvalidField("spec.apiObjects.kind", "kind value", spec.Gvk.Kind)
	}
	if errs := validation.IsGroup(spec.Gvk.Group); len(errs) > 0 {
		return errors.NewInvalidField("spec.apiObjects.group", "group value", spec.Gvk.Group)
	}
	if errs := validation.IsVersion(spec.Gvk.Version); len(errs) > 0 {
		return errors.NewInvalidField("spec.apiObjects.version", "version value", spec.Gvk.Version)
	}
	return nil
}
