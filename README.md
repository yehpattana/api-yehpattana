cmd create migration exaple
migrate create -dir migrati## Installation



## Running the server

```bash
# development
$ go run main.go .env.dev 

# production mode
$ go run main.go .env.prod

## run air
air -c .air.dev.toml

## gen swagger
swag init -g main.go

## DB Query
DROP TABLE [Size]
DROP TABLE [Colors]
DROP TABLE [Oauth]
DROP TABLE [Sub_Menu]
DROP TABLE [Menu]
DROP TABLE [Stock]
DROP TABLE [Products]
DROP TABLE [Orders]
DROP TABLE [Companies]
DROP TABLE [Customers]
DROP TABLE [Admins]
DROP TABLE [Users]
DROP TABLE [schema_migrations]

```

## Run Test
```bash
$ go test
$ go test -cover
$ go test -cover -coverprofile=c.out
$ go tool cover -html=c.out -o coverage.html
$ open coverage.html



```
# api-yehpattana
