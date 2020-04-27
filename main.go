package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"echo-skelton/datamodels"
	datasources "echo-skelton/datasources"
	services "echo-skelton/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/h3poteto/pongo2echo"
	"github.com/joho/godotenv"
	zercleCTX "github.com/khon-kaen-university/echo-context"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// PrdMode from GO_ENV
var PrdMode bool

func main() {
	// Running flag
	runEnv := flag.String("env", "dev", "A env file name without .env")
	flag.Parse()
	// Load env
	err := godotenv.Load(*runEnv + ".env")
	if err != nil {
		log.Fatalf("error while loading the env:\n %+v", err)
	}

	PrdMode = (os.Getenv("GO_ENV") == "production")
	maxDBConn := 8
	maxDBIdle := 4
	if !PrdMode {
		maxDBConn = 2
		maxDBIdle = 1
	}

	// Init database connection
	// Create connection to database
	datasources.DBMain, datasources.DBMainConnStr, err = datasources.NewMariadbDB(
		os.Getenv("DB_NAME_MAIN"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		"",
		maxDBIdle,
		maxDBConn,
		true,
	)
	if err != nil {
		log.Fatalf("Error Connect to database:\n %+v", err)
	}

	// Init redis
	datasources.RedisStore, err = datasources.NewredisDB(
		"tcp",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASSWORD"),
		0,
		time.Duration(30*time.Second),
		os.Getenv("SESS_PREFIX"),
	)
	if err != nil {
		log.Fatalf("error while connect to redis:\n %+v", err)
	}

	// Init JWT
	err = datasources.NewJWTECKey(os.Getenv("JWT_PUBLIC"), os.Getenv("JWT_PRIVATE"))
	if err != nil {
		log.Fatalf("Error init JWT:\n %+v", err)
	}

	// Close DB connection after exit
	defer datasources.DBMain.Close()
	defer datasources.RedisClient.Close()

	// Init app
	app := echo.New()
	app.Use(middleware.Recover())
	app.Pre(middleware.RemoveTrailingSlash())

	// Init HTTPClient
	services.CreateHTTPClient()

	// Store session in redis
	if PrdMode {
		app.Use(session.MiddlewareWithConfig(session.Config{Store: datasources.RedisStore}))
	} else {
		app.Use(middleware.Logger())
		app.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESS_PREFIX")))))
	}

	// Allow API access by cross site request
	crs := middleware.CORS()

	// Validate JWT
	authJWT := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    datasources.JWTECVerifyKey,
		SigningMethod: "ES256",
		ContextKey:    "jwt",
	})

	//Register favicon
	app.File("favicon.ico", "./public/favicon.ico")

	// Register template
	render := pongo2echo.NewRenderer()
	render.AddDirectory("./templates")
	app.Renderer = render

	// Migrate the schema
	datasources.DBMain.AutoMigrate(&datamodels.MainUsers{})

	// Register static file
	app.Static("/", "./public")

	// Index endpoint
	app.GET("/", func(c echo.Context) error {
		ctx := &zercleCTX.Context{c}
		name := ctx.FormValueDefault("name", "Anonymous")
		return ctx.String(http.StatusOK, "Hello "+name)
	}, crs)

	// Token endpoint
	app.GET("/token", func(ctx echo.Context) (err error) {
		jwtToken := ctx.Get("jwt").(*jwt.Token)
		return ctx.JSON(http.StatusOK, jwtToken)
	}, authJWT)

	log.Printf("Runtime ENV: %s", os.Getenv("GO_ENV"))

	// Start HTTP server
	go func() {
		if err := app.Start(os.Getenv("HTTP_PORT")); err != nil {
			app.Logger.Info("shutting down the HTTP server")
		}
	}()

	// Start HTTPS server
	go func() {
		if err := app.StartTLS(os.Getenv("HTTPS_PORT"), os.Getenv("CERT_PATH"), os.Getenv("PRIV_PATH")); err != nil {
			app.Logger.Info("shutting down the HTTPS server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 8 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}

}
