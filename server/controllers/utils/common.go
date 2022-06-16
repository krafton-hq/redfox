package utils

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/krafton-hq/red-fox/apis/idl_common"
)

var CommonReq = &idl_common.CommonReq{}

func ToHttpStatus(res *idl_common.CommonRes) int {
	switch res.Status {
	case idl_common.ResultCode_SUCCESS:
		return http.StatusOK
	case idl_common.ResultCode_NOT_FOUND:
		return http.StatusNotFound
	case idl_common.ResultCode_INTERNAL:
		return http.StatusInternalServerError
	case idl_common.ResultCode_SERVER_INTERNAL:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func HandleGrpcError(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusInternalServerError).SendString(err.Error())
}

func CommonResNotExist(kind string, value string) *idl_common.CommonRes {
	return &idl_common.CommonRes{
		Status:  idl_common.ResultCode_NOT_FOUND,
		Message: fmt.Sprintf("'%s/%s', is Not Exist", kind, value),
	}
}

func CommonResEmpty(kind string, value string) *idl_common.CommonRes {
	return &idl_common.CommonRes{
		Status:  idl_common.ResultCode_NOT_FOUND,
		Message: fmt.Sprintf("'%s/%s', Field Should be Empty", kind, value),
	}
}

func CommonResNotEmpty(field string) *idl_common.CommonRes {
	return &idl_common.CommonRes{
		Status:  idl_common.ResultCode_INVALID_ARGUMENT,
		Message: fmt.Sprintf("'%s' Field Should not be Empty", field),
	}
}

func CommonResDnsLabel(field string, errors []string) *idl_common.CommonRes {
	return &idl_common.CommonRes{
		Status:  idl_common.ResultCode_INVALID_ARGUMENT,
		Message: fmt.Sprintf("'%s' Field Should be RFC1123 Dns Label, See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/#namespaces-and-dns, errors: %v", field, errors),
	}
}

func CommonResFieldValid(field string, errors []string) *idl_common.CommonRes {
	return &idl_common.CommonRes{
		Status:  idl_common.ResultCode_INVALID_ARGUMENT,
		Message: fmt.Sprintf("'%s' Field Validate Failed, errors: %v", field, errors),
	}
}

func CommonResInternalError(err error) *idl_common.CommonRes {
	return &idl_common.CommonRes{
		Status:  idl_common.ResultCode_INTERNAL,
		Message: fmt.Sprintf("Error Occurred %s", err.Error()),
	}
}
