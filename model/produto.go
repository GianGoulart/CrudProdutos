package model

import (
	"encoding/json"
	"errors"
	"io"
)

type Produto struct {
	Codigo            string  `json:"codigo,omitempty" gorm:"primary_key"`
Nome              string  `json:"nome,omitempty" gorm:"size:255;not null"`
PrecoDe           float64 `json:"preco_de,omitempty" gorm:"not null"`
PrecoPor          float64 `json:"preco_por,omitempty" gorm:"not null"`
CriadoEm          string  `json:"criado_em,omitempty" gorm:"not null"`
UltimaAlteracao   string  `json:"ultima_alteracao,omitempty" gorm:"not null"`
EstoqueTotal      int64   `json:"estoque_total,omitempty" gorm:"not null"`
EstoqueCorte      int64   `json:"estoque_corte,omitempty" gorm:"not null"`
EstoqueDisponivel int64   `json:"estoque_disponivel,omitempty" gorm:"not null"`
}

func (me *Produto) PreSave() {
	me.Codigo = NewId()
	me.EstoqueDisponivel = me.EstoqueTotal - me.EstoqueCorte
}

func (me *Produto) Validate() error {
	if me.PrecoDe < me.PrecoPor {
		return errors.New("preço de não pode ser inferior a Preço por")
	}

	if me.EstoqueTotal < me.EstoqueCorte {
		return errors.New("estoque indisponivel")
	}

	return nil
}

func ProdutoFromJson(data io.Reader) (*Produto, error) {
	decoder := json.NewDecoder(data)
	var o *Produto
	err := decoder.Decode(&o)
	return o, err

}
