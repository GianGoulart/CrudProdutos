package produto

import (
	"net/http"

	"github.com/GianGoulart/CrudProdutos/app"
	"github.com/GianGoulart/CrudProdutos/model"
	"github.com/labstack/echo/v4"
)

// Register group item check
func Register(g *echo.Group, apps *app.Container) {
	h := &handler{
		apps: apps,
	}
	g.GET("", h.getProdutos)
	g.GET("/:codigo", h.getProdutoByCodigo)
	g.POST("/produtosByNome", h.getProdutoByNome)
	g.POST("", h.createProduto)
	g.PUT("", h.updateProduto)
	g.DELETE("/:codigo", h.deleteProduto)

}

type handler struct {
	apps *app.Container
}

func (h *handler) getProdutos(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := h.apps.Produto.GetProdutos(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Data: nil,
			Err:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: resp,
	})
}

func (h *handler) getProdutoByCodigo(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := h.apps.Produto.GetProdutoByCodigo(ctx, c.Param("codigo"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Data: nil,
			Err:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: resp,
	})
}

func (h *handler) getProdutoByNome(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(model.Produto)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err.Error(),
		})
	}
	resp, err := h.apps.Produto.GetProdutoByNome(ctx, payload.Nome)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Data: nil,
			Err:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: resp,
	})
}

func (h *handler) createProduto(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(model.Produto)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err.Error(),
		})
	}

	response, err := h.apps.Produto.CreateProduto(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: response,
	})
}

func (h *handler) updateProduto(c echo.Context) error {
	ctx := c.Request().Context()
	payload := new(model.Produto)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err.Error(),
		})
	}

	response, err := h.apps.Produto.UpdateProduto(ctx, payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: response,
	})
}

func (h *handler) deleteProduto(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := h.apps.Produto.DeleteProduto(ctx, c.Param("codigo"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Data: nil,
			Err:  err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Data: resp,
	})
}
