package produto

import (
	"context"

	"github.com/GianGoulart/CrudProdutos/model"
	"github.com/GianGoulart/CrudProdutos/store"
)

// App interface de health para implementação
type IProdutoApp interface {
	GetProdutos(ctx context.Context) (*[]model.Produto, error)
	GetProdutoByCodigo(ctx context.Context, codigo string) (*model.Produto, error)
	GetProdutoByNome(ctx context.Context, nome string) (*[]model.Produto, error)
	CreateProduto(ctx context.Context, produto *model.Produto) (*model.Produto, error)
	UpdateProduto(ctx context.Context, produto *model.Produto) (*model.Produto, error)
	DeleteProduto(ctx context.Context, codigo string) (*model.Produto, error)
}

// NewApp cria uma nova instancia do serviço de health
func NewApp(store *store.Container) IProdutoApp {
	return &appImpl{
		stores: store,
	}
}

type appImpl struct {
	stores *store.Container
}

func (p *appImpl) GetProdutos(ctx context.Context) (*[]model.Produto, error) {
	return p.stores.Produto.FindProdutos(ctx)
}

func (p *appImpl) GetProdutoByCodigo(ctx context.Context, codigo string) (*model.Produto, error) {
	return p.stores.Produto.FindProdutoByCodigo(ctx, codigo)
}

func (p *appImpl) GetProdutoByNome(ctx context.Context, nome string) (*[]model.Produto, error) {
	return p.stores.Produto.FindProdutoByNome(ctx, nome)
}

func (p *appImpl) CreateProduto(ctx context.Context, produto *model.Produto) (*model.Produto, error) {
	produto.PreSave()

	if err := produto.Validate(); err != nil {
		return nil, err
	}

	produto, err := p.stores.Produto.CreateProduto(ctx, produto)
	if err != nil {
		return nil, err
	}

	return produto, err

}

func (p *appImpl) UpdateProduto(ctx context.Context, produto *model.Produto) (*model.Produto, error) {
	if err := produto.Validate(); err != nil {
		return nil, err
	}
	produto, err := p.stores.Produto.UpdateProduto(ctx, produto)
	if err != nil {
		return nil, err
	}

	return produto, err
}

func (p *appImpl) DeleteProduto(ctx context.Context, codigo string) (*model.Produto, error) {

	produto, err := p.stores.Produto.FindProdutoByCodigo(ctx, codigo)
	if err != nil {
		return nil, err
	}

	if err := p.stores.Produto.DeleteProdutoByCodigo(ctx, produto); err != nil {
		return nil, err
	}

	return produto, nil
}
