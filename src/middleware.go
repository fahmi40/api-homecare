package src

import (
	"home-care/app"
	"home-care/middleware"
)

func Middleware() *middlewareUtil {
	if mdlwr == nil {
		mdlwr = &middlewareUtil{}
		mdlwr.Configure()
		mdlwr.isConfigured = true
	}
	return mdlwr
}

var mdlwr *middlewareUtil

type middlewareUtil struct {
	isConfigured bool
}

func (*middlewareUtil) Configure() {
	app.Server().AddMiddleware(middleware.Ctx().New)
	app.Server().AddMiddleware(middleware.DB().New)
}
