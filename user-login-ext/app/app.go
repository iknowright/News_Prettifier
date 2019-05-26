package app
 
import (
    "net/http"
 
	"github.com/gorilla/mux"
	
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"

	helpers "../helpers"
    repos "../repos"
 
	"github.com/gorilla/securecookie"
	"text/template"
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
	a.Router.HandleFunc("/", a.HomePageHandler) // GET
	
	
    a.Router.HandleFunc("/index", a.IndexPageHandler) // GET
	
	a.Router.HandleFunc("/login", a.LoginPageHandler).Methods("GET") // GET
    a.Router.HandleFunc("/login", a.LoginHandler).Methods("POST")
 
    a.Router.HandleFunc("/register", a.RegisterPageHandler).Methods("GET")
    a.Router.HandleFunc("/register", a.RegisterHandler).Methods("POST")
 
    a.Router.HandleFunc("/logout", a.LogoutHandler).Methods("POST")
 
    http.Handle("/", a.Router)
    http.ListenAndServe(":8000", nil)
}


var cookieHandler = securecookie.New(
    securecookie.GenerateRandomKey(64),
    securecookie.GenerateRandomKey(32))
 

type Todo struct {
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

// Handlers
 
// for GET
func (a *App) HomePageHandler(response http.ResponseWriter, request *http.Request) {
    data := TodoPageData{
        PageTitle: "My TODO list",
        Todos: []Todo{
            {Title: "Task 1", Done: false},
            {Title: "Task 2", Done: true},
            {Title: "Task 3", Done: true},
        },
    }
    tmpl := template.Must(template.ParseFiles("templates/home.html"))
    tmpl.Execute(response, data)
}

// for GET
func (a *App) LoginPageHandler(response http.ResponseWriter, request *http.Request) {
    var body, _ = helpers.LoadFile("templates/login.html")
    fmt.Fprintf(response, body)
}
 
// for POST
func (a *App) LoginHandler(response http.ResponseWriter, request *http.Request) {
    name := request.FormValue("name")
    pass := request.FormValue("password")
    redirectTarget := "/"
    if !helpers.IsEmpty(name) && !helpers.IsEmpty(pass) {
        // Database check for user data!
        _userIsValid := repos.UserIsValid(name, pass)
 
        if _userIsValid {
            a.SetCookie(name, response)
            redirectTarget = "/index"
        } else {
            redirectTarget = "/register"
        }
    }
    http.Redirect(response, request, redirectTarget, 302)
}
 
// for GET
func (a *App) RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
    var body, _ = helpers.LoadFile("templates/register.html")
    fmt.Fprintf(response, body)
}
 
// for POST
func (a *App) RegisterHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
 
    uName := r.FormValue("username")
    email := r.FormValue("email")
    pwd := r.FormValue("password")
	confirmPwd := r.FormValue("confirmPassword")
	
	fmt.Println(uName)
	fmt.Println(email)
	fmt.Println(pwd)
	fmt.Println(confirmPwd)
	
 
    _uName, _email, _pwd, _confirmPwd := false, false, false, false
    _uName = !helpers.IsEmpty(uName)
    _email = !helpers.IsEmpty(email)
    _pwd = !helpers.IsEmpty(pwd)
    _confirmPwd = !helpers.IsEmpty(confirmPwd)
 
    if _uName && _email && _pwd && _confirmPwd {
        fmt.Fprintln(w, "Username for Register : ", uName)
        fmt.Fprintln(w, "Email for Register : ", email)
        fmt.Fprintln(w, "Password for Register : ", pwd)
        fmt.Fprintln(w, "ConfirmPassword for Register : ", confirmPwd)
    } else {
        fmt.Fprintln(w, "This fields can not be blank!")
    }
}
 
// for GET
func (a *App) IndexPageHandler(response http.ResponseWriter, request *http.Request) {
    userName := a.GetUserName(request)
    if !helpers.IsEmpty(userName) {
        var indexBody, _ = helpers.LoadFile("templates/index.html")
        fmt.Fprintf(response, indexBody, userName)
    } else {
        http.Redirect(response, request, "/", 302)
    }
}
 
// for POST
func (a *App) LogoutHandler(response http.ResponseWriter, request *http.Request) {
    a.ClearCookie(response)
    http.Redirect(response, request, "/", 302)
}
 
// Cookie
 
func (a *App) SetCookie(userName string, response http.ResponseWriter) {
    value := map[string]string{
        "name": userName,
    }
    if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
        cookie := &http.Cookie{
            Name:  "cookie",
            Value: encoded,
            Path:  "/",
        }
        http.SetCookie(response, cookie)
    }
}
 
func (a *App) ClearCookie(response http.ResponseWriter) {
    cookie := &http.Cookie{
        Name:   "cookie",
        Value:  "",
        Path:   "/",
        MaxAge: -1,
    }
    http.SetCookie(response, cookie)
}
 
func (a *App) GetUserName(request *http.Request) (userName string) {
    if cookie, err := request.Cookie("cookie"); err == nil {
        cookieValue := make(map[string]string)
        if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
            userName = cookieValue["name"]
        }
    }
    return userName
}