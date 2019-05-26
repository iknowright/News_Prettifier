package app
 
import (
    "net/http"
 
	"github.com/gorilla/mux"
	
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
 
    common "../common"
)
 
//App struct
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

const (
	host="news-prettifier.cxo5rl1pvafb.us-east-2.rds.amazonaws.com"
	port="5432"
	user="news_manager"
	password="password"
	dbname="news_prettifier"
)

//Initialize method
func (a *App) Initialize() {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

//Run method
func (a *App) Run(addr string) {
	// ListenAndServer needs a port string and Handler which requires ServeHTTP(ResponseWriter, *Request) method
	// The mux.Router implements ServeHTTP(response, *request)
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}
 
func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/", common.HomePageHandler) // GET
	
	
    a.Router.HandleFunc("/index", common.IndexPageHandler) // GET
	
	a.Router.HandleFunc("/login", common.LoginPageHandler).Methods("GET") // GET
    a.Router.HandleFunc("/login", common.LoginHandler).Methods("POST")
 
    a.Router.HandleFunc("/register", common.RegisterPageHandler).Methods("GET")
    a.Router.HandleFunc("/register", common.RegisterHandler).Methods("POST")
 
    a.Router.HandleFunc("/logout", common.LogoutHandler).Methods("POST")
 
    http.Handle("/", a.Router)
    http.ListenAndServe(":8000", nil)
}