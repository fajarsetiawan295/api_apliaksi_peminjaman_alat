package controllers

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres

	"github.com/tokoumat/api/middlewares"
	"github.com/tokoumat/api/models"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize connect to the database and wire up routes
func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	a.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n Cannot connect to database %s", DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s", DbName)
	}

	a.DB.Debug().AutoMigrate(&models.User{}, &models.DataMedia{}, &models.DataAlat{}, &models.DataPenelitian{}) //database migration

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	var dir string
	a.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json
	flag.StringVar(&dir, "img", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	a.Router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./img/"))))

	a.Router.HandleFunc("/register", a.UserSignUp).Methods("POST")
	a.Router.HandleFunc("/login", a.Login).Methods("POST")

	s := a.Router.PathPrefix("/api").Subrouter() // subrouter to add auth middleware
	s.Use(middlewares.SetContentTypeMiddleware)  // setting content-type to json
	s.Use(middlewares.AuthJwtVerify)

	s.HandleFunc("/users", a.GetAllUsers).Methods("GET")
	s.HandleFunc("/getProfile", a.getProfile).Methods("GET")
	s.HandleFunc("/getbyrole", a.getbyrole).Methods("GET")

	s.HandleFunc("/store-data-media", a.StoreDataMedia).Methods("POST")
	s.HandleFunc("/update-data-media", a.UpdateDataMedia).Methods("POST")
	s.HandleFunc("/get-data-media", a.AllDataMedia).Methods("GET")
	s.HandleFunc("/delete-data-media", a.DeleteDataMedia).Methods("GET")

	s.HandleFunc("/store-data-Alat", a.StoreDataAlat).Methods("POST")
	s.HandleFunc("/update-data-Alat", a.UpdateDataAlat).Methods("POST")
	s.HandleFunc("/get-data-Alat", a.AllDataAlat).Methods("GET")
	s.HandleFunc("/delete-data-Alat", a.DeleteDataAlat).Methods("GET")

	s.HandleFunc("/store-data-penelitian", a.StoreDataPenelitian).Methods("POST")
	s.HandleFunc("/update-data-penelitian", a.UpdateDataPenelitian).Methods("POST")
	s.HandleFunc("/get-data-penelitian", a.AllDataPenelitian).Methods("GET")
	s.HandleFunc("/get-user-data-penelitian", a.UserDataPenelitian).Methods("GET")
	s.HandleFunc("/delete-data-penelitian", a.DeleteDataPenelitian).Methods("GET")

}

func (a *App) RunServer() {
	var port = ":8090"
	log.Printf("\nServer starting on port " + port)
	// log.Fatal(http.ListenAndServe(port, a.Router))
	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(a.Router)))
}
