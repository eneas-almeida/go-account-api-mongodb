# Go API MongoDB

## Pacotes utilizados

-   github.com/gofiber/fiber/v2
-   go.mongodb.org/mongo-driver/mongo
-   go get -u github.com/cosmtrek/air
-   go get github.com/pilu/fresh

## Como utilizar o reaload com o fresh

```bash
# Instala o executável
go install github.com/pilu/fresh

# Para verificar a pasta que foi instalado
# Fica dentro da bin
echo $GOPATH
```

### Configurações do fresh

```bash
# Cria o arquivo .fresh.conf e insere o código abaixo:

root: .

command: go run main.go

build_name: main

build_log: ./tmp/main.log

log_color: true

exclude_dirs:
  - tmp
  - vendor
  - .git

include_files:
  - .go
```

### Como executar o fresh

```bash
# Apenas execute com o comando abaixo, irá carregar o arquivo de conf
fresh
```
