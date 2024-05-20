package api

import (
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	json "github.com/goccy/go-json"
	"github.com/seriouspoop/GoAPI/service/store"
	"github.com/seriouspoop/GoAPI/service/user"
)

type APIServer struct {
	addr string
	db   *mongo.Database
}

func NewAPIServer(addr string, db *mongo.Database) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	v1 := http.NewServeMux()

	router.Handle("/api/v1/", http.StripPrefix("/api/v1", v1))

	//Health Check
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp := map[string]string{
			"service": "healty",
		}
		json.NewEncoder(w).Encode(resp)
	})

	userStore := store.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(v1)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Println("Listening on", s.addr)

	return server.ListenAndServe()
}
