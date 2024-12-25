package handlers

import (
	"encoding/json"
	"net/http"
	"note-keeper/internal/services"
)

func GetNotes(service *services.NoteService, w http.ResponseWriter, r *http.Request) {
	notes, err := service.GetNotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(notes)
}

func CreateNote(service *services.NoteService, w http.ResponseWriter, r *http.Request) {
	var note struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdNote, err := service.CreateNote(note.Title, note.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdNote)
}

func UpdateNoteByID(service *services.NoteService, w http.ResponseWriter, r *http.Request) {
	var note struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedNote, err := service.UpdateNoteByID(note.ID, note.Title, note.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedNote)
}

func DeleteNoteByID(service *services.NoteService, w http.ResponseWriter, r *http.Request) {
	var note struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	deletedNote, err := service.DeleteNoteByID(note.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(deletedNote)
}
