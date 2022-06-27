package domain_helper

import (
	"fmt"
	"net"
	"strconv"

	"github.com/krafton-hq/red-fox/apis/documents"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
)

func ValidationNatIpSpec(spec *documents.NatIpSpec) error {
	if spec == nil {
		return errors.NewInvalidField("spec", "Should not be null", "null")
	}

	if spec.Type == documents.IpType_Ipv4 {
		for index, cidr := range spec.Cidrs {
			_, ipMask, err := net.ParseCIDR(cidr)
			field := fmt.Sprintf("spec.cidrs[%d]", index)
			if err != nil {
				return errors.WrapInvalidField(err, field, "IPv4 CIDR", cidr)
			}
			if len(ipMask.IP) != net.IPv4len {
				return errors.NewInvalidField(field, "4 Bytes", strconv.Itoa(len(ipMask.IP)))
			}
		}
	} else {
		return errors.NewInvalidField("spec.type", "Currently support only Ipv4", spec.Type.String())
	}
	return nil
}
