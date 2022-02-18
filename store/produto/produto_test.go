package produto_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/GianGoulart/CrudProdutos/model"
	"github.com/GianGoulart/CrudProdutos/store/produto"
	"github.com/GianGoulart/CrudProdutos/test"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
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

func Test_FindProdutos(t *testing.T) {

	query := regexp.QuoteMeta("SELECT * FROM `produtos`")

	rows := sqlmock.NewRows([]string{
		"Codigo",
		"Nome",
		"EstoqueTotal",
		"EstoqueCorte",
		"EstoqueDisponivel",
		"PrecoDe",
		"PrecoPor",
		"CriadoEm",
		"UltimaAlteracao",
	}).
		AddRow(
			res[0].Codigo,
			res[0].Nome,
			res[0].EstoqueTotal,
			res[0].EstoqueCorte,
			res[0].EstoqueDisponivel,
			res[0].PrecoDe,
			res[0].PrecoPor,
			res[0].CriadoEm,
			res[0].UltimaAlteracao,
		)

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData interface{}

		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {ExpectedData: &res, PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(query).WillReturnRows(rows)
		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {ExpectedData: new([]model.Produto), PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(query).WillReturnError(nil)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := produto.NewProduto(db)
			ctx := context.Background()

			response, err := store.FindProdutos(ctx)

			if diff := cmp.Diff(response, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if err != nil {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_FindProdutoByCodigo(t *testing.T) {

	query := regexp.QuoteMeta("SELECT * FROM `produtos` WHERE `produtos`.`codigo` = ?")

	rows := sqlmock.NewRows([]string{
		"Codigo",
		"Nome",
		"EstoqueTotal",
		"EstoqueCorte",
		"EstoqueDisponivel",
		"PrecoDe",
		"PrecoPor",
		"CriadoEm",
		"UltimaAlteracao",
	}).
		AddRow(
			res[0].Codigo,
			res[0].Nome,
			res[0].EstoqueTotal,
			res[0].EstoqueCorte,
			res[0].EstoqueDisponivel,
			res[0].PrecoDe,
			res[0].PrecoPor,
			res[0].CriadoEm,
			res[0].UltimaAlteracao,
		)

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData interface{}

		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {ExpectedData: &res[0], PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(query).WillReturnRows(rows)
		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {ExpectedData: new(model.Produto), PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(query).WillReturnError(nil)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := produto.NewProduto(db)
			ctx := context.Background()

			response, err := store.FindProdutoByCodigo(ctx, res[0].Codigo)

			if diff := cmp.Diff(response, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if err != nil {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_FindProdutoByNome(t *testing.T) {

	query := regexp.QuoteMeta("SELECT * FROM `produtos` WHERE nome like ?")

	rows := sqlmock.NewRows([]string{
		"Codigo",
		"Nome",
		"EstoqueTotal",
		"EstoqueCorte",
		"EstoqueDisponivel",
		"PrecoDe",
		"PrecoPor",
		"CriadoEm",
		"UltimaAlteracao",
	}).
		AddRow(
			res[0].Codigo,
			res[0].Nome,
			res[0].EstoqueTotal,
			res[0].EstoqueCorte,
			res[0].EstoqueDisponivel,
			res[0].PrecoDe,
			res[0].PrecoPor,
			res[0].CriadoEm,
			res[0].UltimaAlteracao,
		)

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData interface{}

		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {ExpectedData: &res, PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(query).WillReturnRows(rows)
		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {ExpectedData: new([]model.Produto), PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectQuery(query).WillReturnError(nil)
		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := produto.NewProduto(db)
			ctx := context.Background()

			response, err := store.FindProdutoByNome(ctx, res[0].Nome)

			if diff := cmp.Diff(response, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if err != nil {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_CreateProduto(t *testing.T) {

	query := regexp.QuoteMeta("INSERT INTO `produtos` (`codigo`,`nome`,`preco_de`,`preco_por`,`criado_em`,`ultima_alteracao`,`estoque_total`,`estoque_corte`,`estoque_disponivel`) VALUES (?,?,?,?,?,?,?,?,?)")

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData interface{}

		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {ExpectedData: &res[0], PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec(query).WithArgs(
				res[0].Codigo,
				res[0].Nome,
				res[0].PrecoDe,
				res[0].PrecoPor,
				res[0].CriadoEm,
				res[0].UltimaAlteracao,
				res[0].EstoqueTotal,
				res[0].EstoqueCorte,
				res[0].EstoqueDisponivel,
			).WillReturnResult(sqlmock.NewResult(1, 1))
		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {ExpectedData: new(model.Produto), PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec(query).WillReturnError(nil)

		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := produto.NewProduto(db)
			ctx := context.Background()

			response, err := store.CreateProduto(ctx, &res[0])

			if diff := cmp.Diff(response, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if err != nil {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_UpdateProdutoByCodigo(t *testing.T) {

	query := regexp.QuoteMeta("UPDATE `produtos` SET `nome`=?, `preco_de`=?,`preco_por`=?,`ultima_alteracao`=?,`estoque_total`=?,`estoque_corte`=?,`estoque_disponivel`=? WHERE `codigo`=?")

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData interface{}

		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {ExpectedData: &res[0], PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec(query).WithArgs(
				res[0].Nome,
				res[0].PrecoDe,
				res[0].PrecoPor,
				res[0].UltimaAlteracao,
				res[0].EstoqueTotal,
				res[0].EstoqueCorte,
				res[0].EstoqueDisponivel,
				res[0].Codigo,
			).WillReturnResult(sqlmock.NewResult(1, 1))
		}},
		"deve retornar erro com a mensagem: ocorreu um erro": {ExpectedData: new(model.Produto), PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec(query).WillReturnError(nil)

		}},
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := produto.NewProduto(db)
			ctx := context.Background()

			response, err := store.UpdateProduto(ctx, &res[0])

			if diff := cmp.Diff(response, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if err != nil {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_DeleteProdutoByCodigo(t *testing.T) {

	query := regexp.QuoteMeta("DELETE FROM `produtos` WHERE `codigo`=?")

	cases := map[string]struct {
		ExpectedErr  error
		ExpectedData interface{}

		PrepareMock func(mock sqlmock.Sqlmock)
	}{
		"deve retornar sucesso": {ExpectedData: nil, PrepareMock: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec(query).WithArgs(
				res[0].Codigo,
			).WillReturnResult(sqlmock.NewResult(1, 1))
		}},
	
	}

	for name, cs := range cases {
		t.Run(name, func(t *testing.T) {
			db, mock := test.GetDB()
			cs.PrepareMock(mock)

			store := produto.NewProduto(db)
			ctx := context.Background()

			err := store.DeleteProdutoByCodigo(ctx, &res[0])

			if diff := cmp.Diff(err, cs.ExpectedData); diff != "" {
				t.Error(diff)
			}

			if err != nil {
				assert.NotNil(t, err)
			}
		})
	}
}
