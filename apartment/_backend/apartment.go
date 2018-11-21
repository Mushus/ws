package main

import (
	"log"

	"github.com/Mushus/apartment/backend/component"
	"github.com/Mushus/apartment/backend/handler"
	"github.com/Mushus/apartment/backend/middleware"
	"github.com/Mushus/apartment/backend/model"
	"github.com/labstack/echo"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo-contrib/session"
	_ "github.com/mattn/go-sqlite3"
	"github.com/michaeljs1990/sqlitestore"
)

var (
	devMode = true
)

func main() {
	db := initDB()
	defer db.Close()

	auth := middleware.NewAuth()

	indexHandler := handler.Index{}
	sessionHandler := handler.Session{
		Auth: auth,
		DB:   db,
	}
	apartmentAPIHandler := handler.ApartmentAPI{
		DB: db,
	}
	frontendHandler := handler.Frontend{}
	rentalHandler := handler.Rental{}

	// TODO: キー情報を外出し
	store, err := sqlitestore.NewSqliteStore("./session.db", "sessions", "/", 60*60*24, []byte("hogehoge"))
	if err != nil {
		log.Fatalf("failed to create session store: %v", err)
	}
	sess := session.Middleware(store)

	renderer, err := component.NewTemplateRenderer(component.RendererDevMode(true))
	if err != nil {
		log.Fatal(err)
	}

	validator := component.NewValidator()

	r := echo.New()
	// dev mode
	if devMode {
		r.Use(middleware.DevMode)
	}
	r.Renderer = renderer
	r.Validator = validator
	r.GET("/admin/login", sessionHandler.LoginPage, sess)
	r.POST("/admin/login", sessionHandler.Login, sess)
	r.GET("/admin/logout", sessionHandler.Logout, sess)

	authRequired := r.Group("/admin")
	authRequired.Use(sess, auth.AuthRequired)
	authRequired.Static("/public", "./public")
	authRequired.GET("/print", rentalHandler.Print)
	authRequired.GET("/api/apartment", apartmentAPIHandler.List)
	authRequired.POST("/api/apartment", apartmentAPIHandler.Create)
	authRequired.PUT("/api/apartment/:id", apartmentAPIHandler.Update)
	authRequired.GET("/api/apartment/:id", apartmentAPIHandler.Show)
	authRequired.GET("", indexHandler.RedirectToIndex)
	authRequired.GET("/", indexHandler.Index)

	r.GET("/*", frontendHandler.Index)
	r.Logger.Fatal(r.Start(":8080"))
}

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("failed to create application database: %v", err)
	}

	db.AutoMigrate(model.Apartment{})
	db.AutoMigrate(model.Admin{})
	admin := model.Admin{
		Login: "admin",
	}
	(&admin).SetPassword("admin")
	db.Where("login = ?", admin.Login).Attrs(admin).FirstOrCreate(&model.Admin{})

	return db
}
