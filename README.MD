<h1>API REST EM GOLANG COM GIN</h1> 

<img src="/image/golang e docker.png" width=250px>

</div>
<div style="display: inline_block"><br>
  <img align="center" height="30" width="40" img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" />
  <img align="center" height="30" width="40" img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/html5/html5-original.svg" />
  <img align="center" height="30" width="40" img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/css3/css3-original.svg" />
  <img align="center" height="30" width="40" img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original-wordmark.svg" />
</div>

> <p><h2>Status do Projeto: concluido</h2></p>

### Tópicos 

:small_blue_diamond: [Descrição do projeto](#descrição-do-projeto)

:small_blue_diamond: [Funcionalidades](#funcionalidades)

:small_blue_diamond: [Pré-requisitos](#pré-requisitos)

:small_blue_diamond: [Como rodar a aplicação](#como-rodar-a-aplicação-arrow_forward)


## Descrição do projeto 

<p align="justify">
  Codei essa Api usando apenas Golang e usando o Framework Gin. 

Desenvolvi uma API rest do zero, criando rotas, endpoints, modelos etc. Conectei a minha aplicação com um banco de dados Postgres o qual configurei  usando o Docker-compose. Além disso, implementei o CRUD completo(GET, PATCH, POST, DELETE), fiz rotas para busca por ID, CPF e RG.

Criei e apliquei diferentes validações, realizei a criação de testes dos principais endpoints da minha aplicação, testes esses o qual conecta com banco de dados, cria e deleta recursos para mock de testes, entre outros.

Renderização para visualização das informações da API na página HTML, usando também o framework GIN.
</p>

## Funcionalidades

:heavy_check_mark: Crud completo no banco de dados; 

:heavy_check_mark: Renderização da API na página HTML;

:heavy_check_mark: Testes das funçoes e rotas 

## Pré-requisitos

:warning: [Golang](https://go.dev/dl/)
:warning: [Docker](https://www.docker.com/products/docker-desktop/)

## Dependências a serem instaladas no terminal na pasta do projeto:

:heavy_check_mark: go get -u github.com/gin-gonic/gin)

:heavy_check_mark: go get -u gorm.io/gorm

:heavy_check_mark: go get gorm.io/driver/postgres 

:heavy_check_mark:go get github.com/spf13/viper


## Como rodar a aplicação :arrow_forward:

No terminal, clone o projeto: 

```
$ git clone https://github.com/Gabriel-Newton-dev/gin-api-rest.git
```
Ainda no terminal execute:

```
$ go run main.go
```

## Como rodar os testes

# Para executar todos os testes 
```
$ go test 
```
# Para executar um teste específico

````
$ go test -run + nome do teste
$ go test -run TestSearchByCPF (exemplo)
````

## Casos de Uso

Para utilizar a mesma além de fazer as instalações necessárias conforme dito acima, se faz necessário criar e apontar as suas variáveis de ambiente privadas, para que você consiga configurar o docker-compose e subir o seu container com o respectivo banco de dados. 

## JSON :floppy_disk:

### Usuários: 

|ID |created_at |update_at |delete_at |nome |cpf |rg |
| -------- |-------- |-------- |-------- |-------- |-------- |-------- |
|2 |2023-01-05 22:20:34|2023-01-05 22:20:34||Gabriel Newton|123.456.789-00|12.345.678-9|

## Iniciando/Configurando banco de dados

# Para iniciar o banco de dados se faz necessário apontar as variaveis de ambiente, após rodar o comando:

```
$ docker-compose up
```

## Linguagens, dependencias e libs utilizadas :books:

- [Golang](https://go.dev/dl/)
- [Docker](https://www.docker.com/products/docker-desktop/)
- [Gin](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [Gorm](https://gorm.io/)
- [Postgres](https://www.postgresql.org/download/)

## Resolvendo Problemas :exclamation:

Em [issues](https://github.com/Gabriel-Newton-dev/gin-api-rest/issues) foram abertos alguns problemas gerados durante o desenvolvimento desse projeto e como foram resolvidos. 

## Desenvolvedor:

<img src="/image/fotopf2.jpeg" width=125px>

## Licença 

The [MIT License]() (MIT)

Copyright :copyright: Ano - Titulo do Projeto
