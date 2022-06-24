package utils

import (
	"fmt"
	"net/http"

	go_errors "errors"

	"github.com/gofiber/fiber/v2"
	"github.com/krafton-hq/red-fox/apis/idl_common"
	"github.com/krafton-hq/red-fox/server/pkg/errors"
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
	default:
		return http.StatusInternalServerError
	}
}

func HandleGrpcError(c *fiber.Ctx, err error) error {
	return c.Status(http.StatusInternalServerError).SendString(err.Error())
}

func InvalidArguments(err error) *idl_common.CommonRes {
	return &idl_common.CommonRes{
		Status:  idl_common.ResultCode_INVALID_ARGUMENT,
		Message: err.Error(),
	}
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

func CommonResWithErrorTypes(err error) *idl_common.CommonRes {
	var notFound *errors.NotFoundError
	if go_errors.As(err, &notFound) {
		return &idl_common.CommonRes{
			Status:  idl_common.ResultCode_NOT_FOUND,
			Message: err.Error(),
		}
	}

	var invalidArgs *errors.InvalidArgumentsError
	if go_errors.As(err, &invalidArgs) {
		return &idl_common.CommonRes{
			Status:  idl_common.ResultCode_INVALID_ARGUMENT,
			Message: err.Error(),
		}
	}

	var internal *errors.InternalError
	if go_errors.As(err, &internal) {
		return &idl_common.CommonRes{
			Status:  idl_common.ResultCode_INTERNAL,
			Message: err.Error(),
		}
	}

	return &idl_common.CommonRes{
		Status:  idl_common.ResultCode_INTERNAL,
		Message: fmt.Sprintf("Unexpected Error: %s", err.Error()),
	}
}
