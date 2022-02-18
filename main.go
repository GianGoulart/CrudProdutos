package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/GianGoulart/CrudProdutos/api"
	"github.com/GianGoulart/CrudProdutos/app"
	"github.com/GianGoulart/CrudProdutos/model"
	"github.com/GianGoulart/CrudProdutos/store"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// main configure swagger
//
// method of use bearer token in requests
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	startedAt := time.Now()
	time.Sleep(8 * time.Second)

	model.Watch(func(c model.Config, quit chan bool) {
		e := echo.New()
		e.Validator = model.New()
		e.Debug = c.GetString("crudProdutos") != "prod"
		e.HideBanner = true

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))

		e.Use(middleware.Logger())
		e.Use(middleware.BodyLimit("2M"))
		e.Use(middleware.Recover())
		e.Use(middleware.RequestID())

		fmt.Println(c.GetString("database.writer.url"))
		dbWriter, err := gorm.Open(mysql.Open(c.GetString("database.writer.url")), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		// criação dos stores com a injeção do banco de escrita e leitura
		stores := store.New(store.Options{
			DB: dbWriter,
		})

		// criação dos serviços
		apps := app.New(app.Options{
			Stores:    stores,
			Version:   c.GetString("version"),
			StartedAt: startedAt,
		})

		// registros dos handlers
		api.Register(api.Options{
			Group: e.Group(""),
			Apps:  apps,
		})

		port := c.GetString("server.port")
		// if e.Debug {
		// 	swagger.Register(swagger.Options{
		// 		Port:      port,
		// 		Group:     e.Group("/swagger"),
		// 		AccessKey: c.GetString("docs.key"),
		// 	})
		// }

		// funcão padrão pra tratamento de erros da camada http
		e.HTTPErrorHandler = func(err error, c echo.Context) {
			if c.Response().Committed {
				return
			}

			if err := c.JSON(http.StatusInternalServerError, model.Response{Err: err.Error()}); err != nil {
				logrus.Error(c.Request().Context(), err)
			}
		}

		// função para fechar as conexões
		go func() {
			<-quit

			e.Close()
		}()

		go e.Start(port)

		logrus.Info("Microservice started!")
	})
}
