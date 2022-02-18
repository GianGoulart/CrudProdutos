package store

import (
	"github.com/GianGoulart/CrudProdutos/model"
	"github.com/GianGoulart/CrudProdutos/store/produto"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	Produto produto.IProdutoStore
}

// Options struct de opções para a criação de uma instancia dos repositórios
type Options struct {
	DB *gorm.DB
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	container := &Container{
		Produto: produto.NewProduto(opts.DB),
	}

	opts.DB.AutoMigrate(model.Produto{})
	logrus.Info("Registered -> Store")

	return container
}
