module echo-skelton

go 1.14

replace echo-skelton/datamodels => ./datamodels

replace echo-skelton/datasources => ./datasources

replace echo-skelton/services => ./services

replace echo-skelton/routers => ./routers

replace echo-skelton/oauth => ./oauth

require (
	echo-skelton/datasources v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/flosch/pongo2 v0.0.0-20190707114632-bbf5a6c351f4 // indirect
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gorilla/sessions v1.2.0
	github.com/h3poteto/pongo2echo v0.1.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/joho/godotenv v1.3.0
	github.com/juju/errors v0.0.0-20190930114154-d42613fe1ab9 // indirect
	github.com/khon-kaen-university/echo-context v0.0.4
	github.com/labstack/echo-contrib v0.8.0
	github.com/labstack/echo/v4 v4.1.15
	github.com/lib/pq v1.3.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rbcervilla/redisstore v1.1.0
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a // indirect
	golang.org/x/sys v0.0.0-20200301040627-c5d0d7b4ec88 // indirect
)
