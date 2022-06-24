package domain_helper

import (
	"fmt"

	"github.com/krafton-hq/red-fox/server/pkg/errors"
	"github.com/krafton-hq/red-fox/server/pkg/validation"
)

const subdomainMaxLength int = 63
const nameMaxLength int = subdomainMaxLength
const labelKeyMaxLength int = subdomainMaxLength
const labelValueMaxLength int = 4095

const fieldName = "metadata.name"
const fieldLabelKey = "metadata.label[key]"

func ValidationMetadatable(m Metadatable) error {
	if m == nil {
		return errors.NewInvalidField("$", "Should not be null", "null")
	}

	metadata := m.GetMetadata()
	if errs := validation.IsApiVersion(m.GetApiVersion()); len(errs) > 0 {
		return errors.NewInvalidField("apiVersion", "RFC1123 Dns Label/Version", m.GetApiVersion())
	}

	if len(metadata.Name) == 0 || len(metadata.Name) > nameMaxLength {
		return errors.NewInvalidField(fieldName, fmt.Sprintf("Length Should be [1, %d]", nameMaxLength), metadata.Name)
	}
	if errs := validation.IsDiscoveryName(metadata.Name); len(errs) > 0 {
		return errors.NewInvalidField(fieldName, "RFC1123 Dns Label", metadata.Name)
	}

	for key, value := range metadata.Labels {
		if len(key) == 0 || len(key) > labelKeyMaxLength {
			return errors.NewInvalidField(fieldLabelKey, fmt.Sprintf("Key Length Should be [1, %d]", labelKeyMaxLength), key)
		}
		if len(value) > labelValueMaxLength {
			return errors.NewInvalidField(fieldLabelKey, fmt.Sprintf("Value Length Should be [1, %d]", labelValueMaxLength), value)
		}

		if errs := validation.IsDiscoveryName(key); len(errs) > 0 {
			return errors.NewInvalidField(fieldLabelKey, "RFC1123 Dns Label", metadata.Name)
		}
	}

	for key := range metadata.Annotations {
		if errs := validation.IsAnnotationName(key); len(errs) > 0 {
			return errors.NewInvalidField(fieldLabelKey, fmt.Sprintf("%v", errs), key)
		}
	}
	if errs := validation.IsValidAnnotationsSize(metadata.Annotations); len(errs) > 0 {
		return errors.NewInvalidField("metadata.annotations", fmt.Sprintf("%v", errs), fmt.Sprintf("%v", metadata.Annotations))
	}
	return nil
}
