package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/DarkSoul94/metamemorys_backend/auth"
	authhttp "github.com/DarkSoul94/metamemorys_backend/auth/delivery/http"
	authrepo "github.com/DarkSoul94/metamemorys_backend/auth/repo/postgres"
	authusecase "github.com/DarkSoul94/metamemorys_backend/auth/usecase"
	"github.com/DarkSoul94/metamemorys_backend/metamemory"
	metamemoryhttp "github.com/DarkSoul94/metamemorys_backend/metamemory/delivery/http"
	metamemoryrepo "github.com/DarkSoul94/metamemorys_backend/metamemory/repo/postgres"
	metamemoryusecase "github.com/DarkSoul94/metamemorys_backend/metamemory/usecase"
	"github.com/DarkSoul94/metamemorys_backend/models"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App ...
type App struct {
	authRepo auth.AuthRepo
	authUC   auth.AuthUC

	metamemoryRepo metamemory.MetaRepository
	metamemoryUC   metamemory.MetaUsecase

	httpServer *http.Server
}

// NewApp ...
func NewApp() *App {
	db := initGormDB()
	runGormMigrations(db)

	authRP := authrepo.NewAuthRepo(db)
	authUC := authusecase.NewAuthUC(
		authRP,
		[]byte(viper.GetString("app.auth.signing_key")),
		viper.GetDuration("app.auth.ttl"),
	)

	metaRP := metamemoryrepo.NewMetaRepository(db)
	metaUC := metamemoryusecase.NewMetaUsecase(metaRP)

	return &App{
		authRepo: authRP,
		authUC:   authUC,

		metamemoryRepo: metaRP,
		metamemoryUC:   metaUC,
	}
}

// Run run metamemorylication
func (a *App) Run(port string) error {
	if _, err := os.Stat(viper.GetString("app.store.path")); os.IsNotExist(err) {
		os.Mkdir(viper.GetString("app.store.path"), 0777)
	}

	router := gin.New()
	if viper.GetBool("app.release") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		router.Use(gin.Logger())
	}

	apiRouter := router.Group("/api")

	authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	authhttp.RegisterHTTPEndpoints(apiRouter, a.authUC)

	metamemoryhttp.RegisterHTTPEndpoints(apiRouter, a.metamemoryUC, authMiddleware)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	var l net.Listener
	var err error
	l, err = net.Listen("tcp", a.httpServer.Addr)
	if err != nil {
		panic(err)
	}

	go func(l net.Listener) {
		if err := a.httpServer.Serve(l); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}(l)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initGormDB() *gorm.DB {
	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Kiev",
		viper.GetString("app.db.host"),
		viper.GetString("app.db.login"),
		viper.GetString("app.db.pass"),
		viper.GetString("app.db.name"),
		viper.GetString("app.db.port"),
	)

	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func runGormMigrations(db *gorm.DB) {
	// Migrate the schema
	// Add links to needed models
	err := db.AutoMigrate(
		&models.User{},
		&models.FamilyMember{},
		&models.File{},
	)
	if err != nil {
		panic(err)
	}
}
