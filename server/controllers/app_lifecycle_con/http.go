package app_lifecycle_con

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krafton-hq/red-fox/server/controllers/utils"
)

type AppLifecycleHttp struct {
	server *GrpcController
}

func NewAppLifecycleHttp(server *GrpcController) *AppLifecycleHttp {
	return &AppLifecycleHttp{server: server}
}

func (con *AppLifecycleHttp) Register(router fiber.Router) {
	router.Get("/version", con.Version)
	router.Get("/livez", con.Livez)
	router.Get("/readyz", con.Readyz)
}

func (con *AppLifecycleHttp) Version(c *fiber.Ctx) error {
	res, err := con.server.Version(c.Context(), utils.CommonReq)
	if err != nil {
		return utils.HandleGrpcError(c, err)
	}
	return c.Status(utils.ToHttpStatus(res)).JSON(res)
}

func (con *AppLifecycleHttp) Livez(c *fiber.Ctx) error {
	res, err := con.server.Livez(c.Context(), utils.CommonReq)
	if err != nil {
		return utils.HandleGrpcError(c, err)
	}
	return c.Status(utils.ToHttpStatus(res)).JSON(res)
}

func (con *AppLifecycleHttp) Readyz(c *fiber.Ctx) error {
	res, err := con.server.Readyz(c.Context(), utils.CommonReq)
	if err != nil {
		return utils.HandleGrpcError(c, err)
	}
	return c.Status(utils.ToHttpStatus(res)).JSON(res)
}
