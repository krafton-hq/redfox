package utils

import (
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
