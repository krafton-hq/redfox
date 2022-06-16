package validation

import (
	"strings"

	"k8s.io/apimachinery/pkg/util/validation"
)

func IsApiVersion(apiVersion string) []string {
	segs := strings.Split(apiVersion, "/")
	if len(segs) != 2 {
		return []string{"ApiVersion should contain one '/'"}
	}

	if errors := IsGroup(segs[0]); len(errors) > 0 {
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

func IsVersion(version string) string {
	if len(version) == 0 {
		return validation.EmptyError()
	}
	return ""
}

func IsKind(kind string) string {
	if len(kind) == 0 {
		return validation.EmptyError()
	}
	return ""
}
