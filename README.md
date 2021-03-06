# address-book-go

此為 Address book 後端 API 專案，使用 GO 語言進行開發。

## GO

### 套件管理 Go Module

專案目前使用 Go Module 進行管理，Go 1.11 版本以上才有支援。

#### Go Module

先下指令 `go env` 確認 go module 環境變數是否為 `on`

如果不等於 `on` 的話，下指令

```
export GO111MODULE=on
```

即可打開 go module 的功能。

原則上專案編譯時會自行安裝相關套件，

但也可以先執行下列指令，安裝 module 套件。

```
go mod tidy
```

### How to set up environment?

安裝 docker 後，使用 `docker-compose.yml` 來建立 mysql、redis

```bash
docker-compose up -d
```

並參考 `.env.sample` 來設置 `.env` 

### How to do DB migration?

我們使用 sql-migrate 套件實作 DB migration 功能，

先進行 cmd 安裝

```bash
go get -v github.com/rubenv/sql-migrate/...
```

指令如下：

執行 Migration

```bash
make migrate-up
```

Rollback Migration

```bash
make migrate-down
```

套件連結： Please refer [sql-migrate](https://github.com/rubenv/sql-migrate)

### How to develop?

```shell
go install 
go run main.go
```

## Swagger API Doc

先進行 cmd 安裝

```bash
go get -u github.com/swaggo/swag/cmd/swag
```

1. 產生文件

可以經由 Makefile 執行

```bash
make doc
```

或者原生指令

```bash
swag init
```

2. And then you take [Swagger document](http://localhost:8080/swagger/index.html)

