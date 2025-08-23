package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/tommystone23/plantmanager/handlers"
)

func main() {
	r := mux.NewRouter()

	// api call to upload images
    r.HandleFunc("/api/upload/image", handlers.PlantImageHandler).Methods("POST")

	reactDir := filepath.Join("frontend", "build")
	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(reactDir, "static")))))

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		indexPath := filepath.Join(reactDir, "index.html")
		if _, err := os.Stat(indexPath); os.IsNotExist(err) {
			http.Error(w, "index.html not found", http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, indexPath)
	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
