# Golang for Junior Middle Profi

## Use OS Linux Ubuntu

#### Helper in world golang

## GO - assembler GOOS your OS goENV.go file go

`GOOS=linux GOARCH=amd64 go tool compile -S goEnv.go`

#### Определение структуры types.Container

`https://godoc.org/github.com/docker/docker/api/types#Container`

#### Определение структуры types.ImageSummary

`https://godoc.org/github.com/docker/docker/api/types#ImageSummary`

## Glot library print graphics and save .png

`https://github.com/Arafatk/glot`

## Docker API

`https://docs.docker.com/develop/sdk/examples/`

## Ошибка docker: Error response from daemon: driver failed programming external connectivity on endpoint confident_torvalds (1937bdb71a50a3312e02cf665075fba81ef1faa95c7ab44d07d046a2f3581d71): Bind for 0.0.0.0:5801 failed: port is already allocated.

### sudo lsof -n -i :5800 <- порт который занят && sudo kill -9 28871 <- id процесса
