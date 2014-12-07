package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/justinas/alice"

	"github.com/OPENCBS/server/config"
	"github.com/OPENCBS/server/app"
	"github.com/OPENCBS/server/mssql"
	"github.com/OPENCBS/server/api"
)

type appHandler struct {
	*app.AppContext
	handle func(*app.AppContext, http.ResponseWriter, *http.Request)
}

func (h appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handle(h.AppContext, w, r)
}

type Handler struct {
	*app.AppContext
	handle func(http.ResponseWriter, *api.APIRequest)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	apiRequest := &api.APIRequest{r, h.AppContext.DbProvider}
	h.handle(w, apiRequest)
}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n%s %s %s\n", r.Method, r.URL.String(), r.Proto)
		for k, v := range r.Header {
			fmt.Printf("%s: %s\n", k, v[0])
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func main() {
	db, err := getDb()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}

	dbProvider := &app.DbProvider{Db: db, GetSql: mssql.GetSql}
	context := &app.AppContext{dbProvider}
	commonHandlers := alice.New(loggingHandler)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Handle("/sessions", commonHandlers.Then(Handler{context, api.AddSession})).Methods("POST")
	apiRouter.Handle("/sessions", commonHandlers.Then(Handler{context, api.DeleteSession})).Methods("DELETE")

	apiRouter.Handle("/people", commonHandlers.Then(Handler{context, api.GetPeople})).Methods("GET")

	apiRouter.Handle("/economic-activities", commonHandlers.Then(Handler{context, api.GetEconomicActivities})).Methods("GET")

	//apiRouter.Handle("/users", commonHandlers.Then(appHandler{context, api.GetUsers})).Methods("GET")
	//apiRouter.Handle("/users/{id:[0-9]+}", commonHandlers.Then(appHandler{context, api.GetUser})).Methods("GET")
	//apiRouter.Handle("/sessions", commonHandlers.Then(appHandler{context, api.AddSession})).Methods("POST")
	//apiRouter.Handle("/sessions", commonHandlers.Then(appHandler{context, api.DeleteSession})).Methods("DELETE")
	//apiRouter.Handle("/clients", commonHandlers.Then(appHandler{context, api.GetClients})).Methods("GET")
	//apiRouter.Handle("/people/{id:[0-9]+}", commonHandlers.Then(appHandler{context, api.GetPerson})).Methods("GET")
	//apiRouter.Handle("/activities", commonHandlers.Then(appHandler{context, api.GetActivities})).Methods("GET")
	//apiRouter.Handle("/branches", commonHandlers.Then(appHandler{context, api.GetBranches})).Methods("GET")
	//apiRouter.Handle("/cities", commonHandlers.Then(appHandler{context, api.GetCities})).Methods("GET")
	//apiRouter.Handle("/lookup-data", commonHandlers.Then(appHandler{context, api.GetLookupData})).Methods("GET")
	//apiRouter.Handle("/custom-fields", commonHandlers.Then(appHandler{context, api.GetCustomFields})).Methods("GET")

	http.Handle("/", router)

	log.Println("OPENCBS server is running...")
	http.ListenAndServe(":8080", nil)
}

func getDb() (*sql.DB, error) {
	var conf *config.Configuration
	conf, err := config.Get()
	if err != nil {
		return nil, err
	}

	template := "server=%s;user id=%s;password=%s;database=%s;connection timeout=5"
	connString := fmt.Sprintf(template,
		conf.Database.Host,
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Name)
	db, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
