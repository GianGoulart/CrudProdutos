package produto_test

import (
	"context"
	"testing"
	"time"

	"github.com/GianGoulart/CrudProdutos/app/produto"
	"github.com/GianGoulart/CrudProdutos/mocks"
	"github.com/GianGoulart/CrudProdutos/model"
	"github.com/GianGoulart/CrudProdutos/store"
	"github.com/google/go-cmp/cmp"
)

var (
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
	layout = "02-01-2006T15:04:05"
)

func Test_GetProdutos(t *testing.T) {
	ctx := context.Background()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *[]model.Produto

		InputVersion string

		PrepareMock func(mock *mocks.IProdutoStore)
	}{
		"deve retornar sucesso": {InputVersion: "1", ExpectedData: &res, PrepareMock: func(mock *mocks.IProdutoStore) {
			mock.On("FindProdutos", ctx).
				Return(&res, nil)
		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", ExpectedErr: nil, PrepareMock: func(mock *mocks.IProdutoStore) {
			mock.On("FindProdutos", ctx).
				Return(nil, nil)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoStore)

			cs.PrepareMock(mock)

			app := produto.NewApp(&store.Container{Produto: mock})

			data, err := app.GetProdutos(ctx)

			if diff := cmp.Diff(data, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_GetProdutoByCodigo(t *testing.T) {
	ctx := context.Background()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Produto

		InputVersion string

		PrepareMock func(mock *mocks.IProdutoStore)
	}{
		"deve retornar sucesso": {InputVersion: "1", ExpectedData: &res[0], PrepareMock: func(mock *mocks.IProdutoStore) {
			mock.On("FindProdutoByCodigo", ctx, res[0].Codigo).
				Return(&res[0], nil)
		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", ExpectedErr: nil, PrepareMock: func(mock *mocks.IProdutoStore) {
			mock.On("FindProdutoByCodigo", ctx, res[0].Codigo).
				Return(nil, nil)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoStore)

			cs.PrepareMock(mock)

			app := produto.NewApp(&store.Container{Produto: mock})

			data, err := app.GetProdutoByCodigo(ctx, res[0].Codigo)

			if diff := cmp.Diff(data, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_GetProdutoByNome(t *testing.T) {
	ctx := context.Background()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *[]model.Produto

		InputVersion string

		PrepareMock func(mock *mocks.IProdutoStore)
	}{
		"deve retornar sucesso": {InputVersion: "1", ExpectedData: &res, PrepareMock: func(mock *mocks.IProdutoStore) {
			mock.On("FindProdutoByNome", ctx, res[0].Nome).
				Return(&res, nil)
		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", ExpectedErr: nil, PrepareMock: func(mock *mocks.IProdutoStore) {
			mock.On("FindProdutoByNome", ctx, res[0].Nome).
				Return(nil, nil)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoStore)

			cs.PrepareMock(mock)

			app := produto.NewApp(&store.Container{Produto: mock})

			data, err := app.GetProdutoByNome(ctx, res[0].Nome)

			if diff := cmp.Diff(data, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_CreateProduto(t *testing.T) {
	ctx := context.Background()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Produto

		InputVersion string

		PrepareMock func(mock *mocks.IProdutoStore)
	}{
		"deve retornar sucesso": {InputVersion: "1", ExpectedData: &res[0], PrepareMock: func(mock *mocks.IProdutoStore) {
			mock.On("CreateProduto", ctx, &res[0]).
				Return(&res[0], nil)
		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", ExpectedErr: nil, PrepareMock: func(mock *mocks.IProdutoStore) {
			mock.On("CreateProduto", ctx, &res[0]).
				Return(nil, nil)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoStore)

			cs.PrepareMock(mock)

			app := produto.NewApp(&store.Container{Produto: mock})

			data, err := app.CreateProduto(ctx, &res[0])

			if diff := cmp.Diff(data, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_UpdateProduto(t *testing.T) {
	ctx := context.Background()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Produto

		InputVersion string

		PrepareMock func(mock *mocks.IProdutoStore)
	}{
		"deve retornar sucesso": {InputVersion: "1", ExpectedData: &res[0], PrepareMock: func(mock *mocks.IProdutoStore) {
			mock.On("UpdateProduto", ctx, &res[0]).
				Return(&res[0], nil)
		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {InputVersion: "1", ExpectedErr: nil, PrepareMock: func(mock *mocks.IProdutoStore) {
			mock.On("UpdateProduto", ctx, &res[0]).
				Return(nil, nil)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoStore)

			cs.PrepareMock(mock)

			app := produto.NewApp(&store.Container{Produto: mock})

			data, err := app.UpdateProduto(ctx, &res[0])

			if diff := cmp.Diff(data, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_DeleteProduto(t *testing.T) {
	ctx := context.Background()

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData *model.Produto

		InputVersion string

		PrepareMock func(mock *mocks.IProdutoStore)
	}{
		"deve retornar sucesso": {InputVersion: "1", ExpectedData: &res[0], PrepareMock: func(mock *mocks.IProdutoStore) {

			mock.On("FindProdutoByCodigo", ctx, res[0].Codigo).
				Return(&res[0], nil)

			mock.On("DeleteProdutoByCodigo", ctx, &res[0]).
				Return(nil)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			mock := new(mocks.IProdutoStore)

			cs.PrepareMock(mock)

			app := produto.NewApp(&store.Container{Produto: mock})

			data, err := app.DeleteProduto(ctx, res[0].Codigo)

			if diff := cmp.Diff(data, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if diff := cmp.Diff(err, cs.ExpectedErr); diff != "" {
				t.Error(diff)
			}
		})
	}
}
