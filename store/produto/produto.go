package produto

import (
	"context"
	"time"

	"github.com/GianGoulart/CrudProdutos/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Store interface para implementação do health
type IProdutoStore interface {
	FindProdutos(ctx context.Context) (*[]model.Produto, error)
	CreateProduto(ctx context.Context, produto *model.Produto) (*model.Produto, error)
	FindProdutoByCodigo(ctx context.Context, codigo string) (*model.Produto, error)
	FindProdutoByNome(ctx context.Context, nome string) (*[]model.Produto, error)
	UpdateProduto(ctx context.Context, produto *model.Produto) (*model.Produto, error)
	DeleteProdutoByCodigo(ctx context.Context, produto *model.Produto) error
}

// NewProduto cria uma nova instancia do repositorio de produto
func NewProduto(reader *gorm.DB) IProdutoStore {
	return &storeImpl{reader}
}

type storeImpl struct {
	db *gorm.DB
}

const (
	layout = "02-01-2006T15:04:05"
)

func (r *storeImpl) FindProdutos(ctx context.Context) (*[]model.Produto, error) {
	produtos := new([]model.Produto)

	if err := r.db.Find(&produtos).Error; err != nil {

		logrus.Error("store.produtos.FindProdutos", err.Error())
		return produtos, err
	}

	return produtos, nil

}

func (r *storeImpl) FindProdutoByCodigo(ctx context.Context, codigo string) (*model.Produto, error) {

	res := new(model.Produto)

	if err := r.db.WithContext(ctx).Where(&model.Produto{Codigo: codigo}).Find(res).Error; err != nil {

		logrus.Error("store.produto.FindProdutoByCodigo", err.Error())
		return res, err
	}

	return res, nil

}

func (r *storeImpl) FindProdutoByNome(ctx context.Context, nome string) (*[]model.Produto, error) {
	produtos := new([]model.Produto)

	if err := r.db.Where("nome like ?", nome+"%").Find(&produtos).Error; err != nil {

		logrus.Error("store.produtos.FindProdutoByNome", err.Error())
		return produtos, err
	}

	return produtos, nil

}

func (r *storeImpl) CreateProduto(ctx context.Context, produto *model.Produto) (*model.Produto, error) {

	produto.CriadoEm = time.Now().Format(layout)
	produto.UltimaAlteracao = time.Now().Format(layout)

	exec := "INSERT INTO `produtos` (`codigo`,`nome`,`preco_de`,`preco_por`,`criado_em`,`ultima_alteracao`,`estoque_total`,`estoque_corte`,`estoque_disponivel`) VALUES (?,?,?,?,?,?,?,?,?)"

	if err := r.db.Exec(exec, produto.Codigo, produto.Nome, produto.PrecoDe, produto.PrecoPor, produto.CriadoEm, produto.UltimaAlteracao, produto.EstoqueTotal, produto.EstoqueCorte, produto.EstoqueDisponivel).Error; err != nil {
		logrus.Error("store.produto.CreateProduto", err.Error())
		return &model.Produto{}, err
	}

	return produto, nil
}

func (r *storeImpl) UpdateProduto(ctx context.Context, produto *model.Produto) (*model.Produto, error) {

	produto.UltimaAlteracao = time.Now().Format(layout)
	produto.EstoqueDisponivel = produto.EstoqueTotal - produto.EstoqueCorte

	exec := "UPDATE `produtos` SET `nome`=?, `preco_de`=?,`preco_por`=?,`ultima_alteracao`=?,`estoque_total`=?,`estoque_corte`=?,`estoque_disponivel`=? WHERE `codigo`=?"

	if err := r.db.Exec(exec, produto.Nome, produto.PrecoDe, produto.PrecoPor, produto.UltimaAlteracao, produto.EstoqueTotal, produto.EstoqueCorte, produto.EstoqueDisponivel, produto.Codigo).Error; err != nil {
		logrus.Error("store.produto.UpdateProdutoByCodigo", err.Error())
		return &model.Produto{}, err
	}

	return produto, nil
}

func (r *storeImpl) DeleteProdutoByCodigo(ctx context.Context, produto *model.Produto) error {

	exec := "DELETE FROM `produtos` WHERE `codigo`=?"
	if err := r.db.Exec(exec, produto.Codigo).Error; err != nil {
		logrus.Error("store.produto.DeleteProdutoByCodigo", err.Error())
		return err
	}

	return nil
}
