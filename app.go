package main // [ INIT OF THE SERVICE PACKAGE ]

import (
	// "database/sql"
	// "databese/sql"
	"database/sql"
	"log"
	"net/http"

	// "net/http"
	"os"

	// "encoding.json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	ID			int			`json:"id"`
	Name		string	`json:"name"`
	Email		string	`json:"email"`		
}

func main() {
	// [ CONNECT DATABASE ]
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// [ CREATE ROUTER ]
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
	router.HandleFunc("/users", createUse(db)).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser(db)).Methods("DELETE")

	// [ START SERVER ]
	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleWare(router)))
}

func jsonContentTypeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
