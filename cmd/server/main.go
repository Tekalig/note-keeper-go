package main

import (
	"log"
	"net/http"
	"note-keeper/config"
	"note-keeper/internal/graphql"
	"note-keeper/internal/handlers"
	"note-keeper/internal/services"
)

func main() {
	// Load configuration
	config.Load()

	// Initialize GraphQL client
	client := graphql.NewClient(config.HasuraURL, config.HasuraAdminSecret)

	// Initialize services
	noteService := services.NewNoteService(client)

	// Define routes
	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetNotes(noteService, w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateNote(noteService, w, r)
		} else if r.Method == http.MethodPut {
			handlers.UpdateNoteByID(noteService, w, r)
		} else if r.Method == http.MethodDelete {
			handlers.DeleteNoteByID(noteService, w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
