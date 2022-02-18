package api

import (
	"github.com/GianGoulart/CrudProdutos/api/produto"
	"github.com/GianGoulart/CrudProdutos/app"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Options struct de opções para a criação de uma instancia das rotas
type Options struct {
	Group *echo.Group
	Apps  *app.Container
}

// Register api instance
func Register(opts Options) {
	produto.Register(opts.Group.Group("produtos"), opts.Apps)

	logrus.Info("Registered -> Api")
}
