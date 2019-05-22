package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"io/ioutil"
	
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

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

//InitializeRoutes method
func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/news", a.getMultipleNews).Methods("GET")
	a.Router.HandleFunc("/news", a.createNews).Methods("POST")
	a.Router.HandleFunc("/news/{id:[0-9]+}", a.getNews).Methods("GET")
	a.Router.HandleFunc("/news/{id:[0-9]+}", a.updateNews).Methods("PUT")
	a.Router.HandleFunc("/news/{id:[0-9]+}", a.deleteNews).Methods("DELETE")
	a.Router.HandleFunc("/prettified/{id:[0-9]+}", a.prettifiedNews).Methods("GET")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) prettifiedNews(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/single_news.html")  // Parse template file.
	n := news{ID: 1, Link: "google.com", ReturnLink: "notGoogle.com"}
    t.Execute(w, n)  // merge.
}

func (a *App) getNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid news ID")
		return
	}
	n := news{ID: id}
	if err := n.getNews(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "News not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, n)
}

func (a *App) getMultipleNews(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	multipleNews, err := getMultipleNews(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, multipleNews)
}

func (a *App) updateNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid news ID")
		return
	}

	var n news
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	n.ID = id

	if err := n.updateNews(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, n)
}

func (a *App) deleteNews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid news ID")
		return
	}

	n := news{ID: id}
	if err := n.deleteNews(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) createNews(w http.ResponseWriter, r *http.Request) {
	var n news
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&n); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	resp, err := http.Get(n.Link)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)
	if err := n.createNews(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, n)
}