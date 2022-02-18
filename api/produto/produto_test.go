package produto

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/GianGoulart/CrudProdutos/app"
	"github.com/GianGoulart/CrudProdutos/mocks"
	"github.com/GianGoulart/CrudProdutos/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	layout = "02-01-2006T15:04:05"

	res = []model.Produto{{
		Codigo:            "908a9f80dv-dv9s080v-dv90d90",
		Nome:              "Televisao SAMSUNG",
		EstoqueTotal:      100,
		EstoqueCorte:      10,
		EstoqueDisponivel: 90,
		PrecoDe:           float64(2500),
		PrecoPor:          float64(2200),
		CriadoEm:          time.Now().Format(layout),
		UltimaAlteracao:   time.Now().Format(layout),
	}}
	erro = errors.New("ocorreu um erro")
)

func Test_getProdutos(t *testing.T) {
	e := echo.New()
	startedAt := time.Now()
	ctx := context.Background()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData int

		InputVersion  string
		InputDatetime time.Time

		PrepareMock func(mock *mocks.IProdutoApp)
	}{
		"deve retornar sucesso": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusOK, PrepareMock: func(mock *mocks.IProdutoApp) {
			mock.On("GetProdutos", ctx).Return(&res, nil)

		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusInternalServerError, PrepareMock: func(mock *mocks.IProdutoApp) {
			mock.On("GetProdutos", ctx).Return(nil, erro)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoApp)

			cs.PrepareMock(mock)

			request, err := http.NewRequest(http.MethodGet, "/produtos", nil)
			if err != nil {
				t.Fatal(err)
			}
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			h := handler{
				apps: &app.Container{Produto: mock},
			}

			c := e.NewContext(request, rr)

			if assert.NoError(t, h.getProdutos(c)) {
				assert.Equal(t, cs.ExpectedData, rr.Code)
			}

		})
	}
}

func Test_getProdutoByCodigo(t *testing.T) {
	e := echo.New()
	startedAt := time.Now()
	ctx := context.Background()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData int

		InputVersion  string
		InputDatetime time.Time

		PrepareMock func(mock *mocks.IProdutoApp)
	}{
		"deve retornar sucesso": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusOK, PrepareMock: func(mocks *mocks.IProdutoApp) {
			mocks.On("GetProdutoByCodigo", ctx, mock.Anything).Return(&res[0], nil)

		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusInternalServerError, PrepareMock: func(mocks *mocks.IProdutoApp) {
			mocks.On("GetProdutoByCodigo", ctx, mock.Anything).Return(nil, erro)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoApp)

			cs.PrepareMock(mock)

			request, err := http.NewRequest(http.MethodGet, "/produtos/908a9f80dv-dv9s080v-dv90d90", nil)
			if err != nil {
				t.Fatal(err)
			}
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			h := handler{
				apps: &app.Container{Produto: mock},
			}

			c := e.NewContext(request, rr)

			if assert.NoError(t, h.getProdutoByCodigo(c)) {
				assert.Equal(t, cs.ExpectedData, rr.Code)
			}

		})
	}
}

func Test_getProdutoByNome(t *testing.T) {
	e := echo.New()
	startedAt := time.Now()
	ctx := context.Background()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData int

		InputVersion  string
		InputDatetime time.Time

		PrepareMock func(mock *mocks.IProdutoApp)
	}{
		"deve retornar sucesso": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusOK, PrepareMock: func(mocks *mocks.IProdutoApp) {
			mocks.On("GetProdutoByNome", ctx, mock.Anything).Return(&res, nil)

		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusInternalServerError, PrepareMock: func(mocks *mocks.IProdutoApp) {
			mocks.On("GetProdutoByNome", ctx, mock.Anything).Return(nil, erro)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoApp)

			cs.PrepareMock(mock)

			request, err := http.NewRequest(http.MethodGet, "/produtos/produtosByNome", nil)
			if err != nil {
				t.Fatal(err)
			}
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			h := handler{
				apps: &app.Container{Produto: mock},
			}

			c := e.NewContext(request, rr)

			if assert.NoError(t, h.getProdutoByNome(c)) {
				assert.Equal(t, cs.ExpectedData, rr.Code)
			}

		})
	}
}

func Test_createProduto(t *testing.T) {
	e := echo.New()
	startedAt := time.Now()
	ctx := context.Background()

	body, _ := json.Marshal(res[0])

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData int
		BodyReq      io.Reader

		InputVersion  string
		InputDatetime time.Time

		PrepareMock func(mock *mocks.IProdutoApp)
	}{
		"deve retornar sucesso": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusOK, BodyReq: strings.NewReader(string(body)), PrepareMock: func(mocks *mocks.IProdutoApp) {
			mocks.On("CreateProduto", ctx, mock.Anything).Return(&res[0], nil)

		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusBadRequest, BodyReq: strings.NewReader(string(body)), PrepareMock: func(mocks *mocks.IProdutoApp) {
			mocks.On("CreateProduto", ctx, mock.Anything).Return(nil, erro)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoApp)

			cs.PrepareMock(mock)

			request, err := http.NewRequest(http.MethodPost, "/produtos", cs.BodyReq)
			if err != nil {
				t.Fatal(err)
			}
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			h := handler{
				apps: &app.Container{Produto: mock},
			}

			c := e.NewContext(request, rr)

			if assert.NoError(t, h.createProduto(c)) {
				assert.Equal(t, cs.ExpectedData, rr.Code)
			}

		})
	}
}

func Test_updateProduto(t *testing.T) {
	e := echo.New()
	startedAt := time.Now()
	ctx := context.Background()

	body, _ := json.Marshal(res[0])

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData int
		BodyReq      io.Reader

		InputVersion  string
		InputDatetime time.Time

		PrepareMock func(mock *mocks.IProdutoApp)
	}{
		"deve retornar sucesso": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusOK, BodyReq: strings.NewReader(string(body)), PrepareMock: func(mocks *mocks.IProdutoApp) {
			mocks.On("UpdateProduto", ctx, mock.Anything).Return(&res[0], nil)

		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusBadRequest, BodyReq: strings.NewReader(string(body)), PrepareMock: func(mocks *mocks.IProdutoApp) {
			mocks.On("UpdateProduto", ctx, mock.Anything).Return(nil, erro)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoApp)

			cs.PrepareMock(mock)

			request, err := http.NewRequest(http.MethodPut, "/produtos", cs.BodyReq)
			if err != nil {
				t.Fatal(err)
			}
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			h := handler{
				apps: &app.Container{Produto: mock},
			}

			c := e.NewContext(request, rr)

			if assert.NoError(t, h.updateProduto(c)) {
				assert.Equal(t, cs.ExpectedData, rr.Code)
			}

		})
	}
}

func Test_deleteProduto(t *testing.T) {
	e := echo.New()
	startedAt := time.Now()
	ctx := context.Background()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData int

		InputVersion  string
		InputDatetime time.Time

		PrepareMock func(mock *mocks.IProdutoApp)
	}{
		"deve retornar sucesso": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusOK, PrepareMock: func(mocks *mocks.IProdutoApp) {
			mocks.On("DeleteProduto", ctx, mock.Anything).Return(&res[0], nil)

		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", InputDatetime: startedAt, ExpectedData: http.StatusBadRequest, PrepareMock: func(mocks *mocks.IProdutoApp) {
			mocks.On("DeleteProduto", ctx, mock.Anything).Return(nil, erro)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoApp)

			cs.PrepareMock(mock)

			request, err := http.NewRequest(http.MethodDelete, "/produtos/908a9f80dv-dv9s080v-dv90d90", nil)
			if err != nil {
				t.Fatal(err)
			}
			request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			h := handler{
				apps: &app.Container{Produto: mock},
			}

			c := e.NewContext(request, rr)

			if assert.NoError(t, h.deleteProduto(c)) {
				assert.Equal(t, cs.ExpectedData, rr.Code)
			}

		})
	}
}
