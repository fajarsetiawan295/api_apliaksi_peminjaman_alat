package controllers

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres

	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/middlewares"
	"github.com/fajars295/api_apliaksi_peminjaman_alat/api/models"
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

	a.DB.Debug().AutoMigrate(&models.User{}, &models.Menu{}, &models.DataMedia{}, &models.PeminjamanBarang{}, &models.PeminjamanAlat{}, &models.DataAlat{}, &models.KerusakanAlat{}) //database migration

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
	s.HandleFunc("/UpdateUser", a.UpdateUser).Methods("POST")
	s.HandleFunc("/getById", a.getById).Methods("GET")
	s.HandleFunc("/UpdateIdUser", a.UpdateIdUser).Methods("POST")

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

	s.HandleFunc("/store-peminjaman-media", a.StorePeminjaman).Methods("POST")
	s.HandleFunc("/list-peminjaman-media", a.MePeminjaman).Methods("GET")
	s.HandleFunc("/delete-peminjaman-media", a.DeletePeminjaman).Methods("GET")
	s.HandleFunc("/status-peminjaman-media", a.StatusPeminjaman).Methods("GET")
	s.HandleFunc("/update-peminjaman-media", a.UpdatePeminjaman).Methods("POST")

	s.HandleFunc("/store-peminjaman-alat", a.StorepeminjamanAlat).Methods("POST")
	s.HandleFunc("/list-peminjaman-alat", a.MePeminjamanAlat).Methods("GET")
	s.HandleFunc("/delete-peminjaman-alat", a.DeletePeminjamanAlat).Methods("GET")
	s.HandleFunc("/list-user-peminjaman-alat", a.UserPeminjamanAlat).Methods("GET")

	s.HandleFunc("/store-kerusakan-alat", a.StoreKerusakan).Methods("POST")
	s.HandleFunc("/list-kerusakan-alat", a.MeKerusakan).Methods("GET")
	s.HandleFunc("/delete-kerusakan-alat", a.DeleteKerusakan).Methods("GET")
	s.HandleFunc("/user-list-kerusakan-alat", a.UserKerusakan).Methods("GET")
	s.HandleFunc("/update-list-kerusakan-alat", a.UpdateKerusakan).Methods("POST")
	s.HandleFunc("/AllKerusakan-list-kerusakan-alat", a.AllKerusakan).Methods("GET")

	s.HandleFunc("/strore-menu", a.StoreMenu).Methods("POST")
	s.HandleFunc("/update-menu", a.UpdateMenu).Methods("POST")
	s.HandleFunc("/All-menu", a.AllMenu).Methods("GET")
	s.HandleFunc("/ByRole-menu", a.ByRoleMenu).Methods("GET")

}

func (a *App) RunServer() {
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
	var port = ":8092"
	log.Printf("\nServer starting on port " + port)
	// log.Fatal(http.ListenAndServe(port, a.Router))
	log.Fatal(http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(a.Router)))
}

func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
