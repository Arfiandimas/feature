package routes

import (
	"mandiri_payment_gateway_service/base"
	"mandiri_payment_gateway_service/routes/api"
	"mandiri_payment_gateway_service/routes/web"
	"os"
)

type RouteInterface interface {
	Do(router *base.Router)
}

func Init() {
	router := base.NewRouter()
	api.Init().Do(router)
	web.Init().Do(router)

	router.Run(":" + os.Getenv("APP_PORT"))
}
