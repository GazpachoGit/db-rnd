# Get DB data with join

## Prerequisites

1. Go language package installed
2. Docker installed
3. From the project root directory in the prompt run command:

```
docker-compose --project-name="db-rnd" up -d
```

4. Wait and check that all containers are up and running. You can use the command

```
docker container ls
```

## GORM version

### Description

Get data using Go ORM.

### Start

From the project root dir run command:

```
go run ./cmd/gorm_version/main.go
```

## PGX version

### Description

Get data using SQL syntax query

### Start

From the project root dir run command:

```
go run ./cmd/pgx_version/main.go
```
