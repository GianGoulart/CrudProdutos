# Imagem base
FROM golang

# Adiciona informações da pessoa que mantem a imagem
LABEL maintainer="Gian Goulart <giangoulart1994@gmail.com>"

# diretoria um diretorio de trabalho
WORKDIR /app/src/teste

# aponta a variavel gopath do go para o diretorio app
ENV GOPATH=/app

# copia os arquivos do projeto para o workdir do container
# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
#COPY *.go ./
COPY . ./

# execulta o main.go e baixa as dependencias do projeto
RUN go build main.go

# Comando para rodar o executavel
ENTRYPOINT ["./main"]

# expõe a pota 5055
EXPOSE 5055

