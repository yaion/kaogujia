package routes

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	v1 "kaogujia/routes/v1"
)

func RegisterRoutes(h *server.Hertz) {
	v1.Register(h)
}
