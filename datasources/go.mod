module echo-skelton/datasources

go 1.14

replace echo-skelton/datamodels => ../datamodels

replace echo-skelton/services => ../services

replace echo-skelton/routers => ../routers

replace echo-skelton/oauth => ../oauth

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.12
	github.com/jmoiron/sqlx v1.2.0
	github.com/rbcervilla/redisstore v1.1.0
)
