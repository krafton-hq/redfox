package validation

import (
	"strings"

	api_validation "k8s.io/apimachinery/pkg/api/validation"
	"k8s.io/apimachinery/pkg/util/validation"
)

func IsApiVersion(apiVersion string) []string {
	apiGroup, version, found := strings.Cut(apiVersion, "/")
	if !found {
		return []string{"ApiVersion should contain one '/'"}
	}

	if errors := IsGroup(apiGroup); len(errors) > 0 {
		return errors
	}
	if errors := IsVersion(version); len(errors) > 0 {
		return errors
	}
	return nil
}

func IsGroup(group string) []string {
	return validation.IsDNS1123Subdomain(group)
}

func IsDiscoveryName(name string) []string {
	if name == "" {
		return []string{validation.EmptyError()}
	}
	return validation.IsDNS1123Label(name)
}

func IsVersion(version string) []string {
	if len(version) == 0 {
		return []string{validation.EmptyError()}
	}
	return nil
}

func IsKind(kind string) []string {
	if len(kind) == 0 {
		return []string{validation.EmptyError()}
	}
	return nil
}

func IsAnnotationName(value string) []string {
	return validation.IsQualifiedName(value)
}

func IsValidAnnotationsSize(annotations map[string]string) []string {
	err := api_validation.ValidateAnnotationsSize(annotations)
	if err != nil {
		return []string{err.Error()}
	}
	return nil
}
