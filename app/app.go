package app

import (
	"time"

	"github.com/GianGoulart/CrudProdutos/app/produto"
	"github.com/GianGoulart/CrudProdutos/store"
	"github.com/sirupsen/logrus"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	Produto produto.IProdutoApp
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Stores *store.Container

	StartedAt time.Time
	Version   string
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {

	container := &Container{
		// Health:            health.NewApp(opts.Stores, opts.Version, opts.StartedAt),
		Produto: produto.NewApp(opts.Stores),
	}

	logrus.Info("Registered -> App")

	return container

}
