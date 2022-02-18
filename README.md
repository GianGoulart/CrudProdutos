# CrudProdutos
Crud de cadastro de produtos, estoques e preços

# Rodar a aplicação
Buildar a imagem via docker-compose

```docker-compose build```

e depois pra subir a aplicação

```docker-compose up```

# CURL`s

GET- Todos os Produtos
```
curl --location --request GET 'http://localhost:5055/produtos'
```

GET- Busca produto pelo codigo
```
curl --location --request GET 'http://localhost:5055/produtos/:codigo'
```

POST- Busca produto pelo nome
```
curl --location --request POST 'http://localhost:5055/produtos/produtosByNome' \
--header 'Content-Type: application/json' \
--data-raw '{
        "nome": "TV"
    }'
```

POST- Criar produto 
```
curl --location --request POST 'http://localhost:5055/produtos' \
--header 'Content-Type: application/json' \
--data-raw '{
    "nome": "TV SONY",
    "preco_de": 4000,
    "preco_por": 3700,
    "estoque_total": 250,
    "estoque_corte": 10
}'
```

PUT- Alterar produto 
```
curl --location --request PUT 'http://localhost:5055/produtos' \
--header 'Content-Type: application/json' \
--data-raw '{
    "codigo": "pfpugbto17g9zxis6p8916rcge",
    "nome": "TV SAMSUNG 55",
    "preco_de": 2000,
    "preco_por": 1700,
    "estoque_total": 100,
    "estoque_corte": 10
}'
```

POST- Deletar produto 
```
curl --location --request DELETE 'http://localhost:5055/produtos/:codigo'
```