module echo-skelton

go 1.14

replace echo-skelton/datamodels => ./datamodels

replace echo-skelton/datasources => ./datasources

replace echo-skelton/services => ./services

replace echo-skelton/routers => ./routers

replace echo-skelton/oauth => ./oauth

require (
	echo-skelton/datamodels v0.0.0-00010101000000-000000000000 // indirect
	echo-skelton/datasources v0.0.0-00010101000000-000000000000 // indirect
	echo-skelton/services v0.0.0-00010101000000-000000000000 // indirect
	github.com/h3poteto/pongo2echo v0.1.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/khon-kaen-university/echo-context v0.0.4 // indirect
	github.com/labstack/echo-contrib v0.8.0 // indirect
	github.com/labstack/echo/v4 v4.1.15 // indirect
	github.com/valyala/fastjson v1.5.0 // indirect
)
