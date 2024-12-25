package services

import (
	"note-keeper/internal/graphql"
)

type NoteService struct {
	client *graphql.Client
}

func NewNoteService(client *graphql.Client) *NoteService {
	return &NoteService{client: client}
}

func (s *NoteService) GetNotes() (interface{}, error) {
	query := `query { notes { id title content } }`
	result, err := s.client.Query(query, nil)
	if err != nil {
		return nil, err
	}
	return result["data"], nil
}

func (s *NoteService) CreateNote(title, content string) (interface{}, error) {
	query := `mutation ($title: String!, $content: String!) {
		insert_notes(objects: {title: $title, content: $content}) {
			returning { id title content }
		}
	}`
	variables := map[string]interface{}{
		"title":   title,
		"content": content,
	}
	result, err := s.client.Query(query, variables)
	if err != nil {
		return nil, err
	}
	return result["data"], nil
}

func (s *NoteService) UpdateNoteByID(id, title, content string) (interface{}, error) {
	query := `mutation ($id: Int!, $title: String!, $content: String!) {
		update_notes(where: {id: {_eq: $id}}, _set: {title: $title, content: $content}) {
			returning { id title content }
		}
	}`
	variables := map[string]interface{}{
		"id":      id,
		"title":   title,
		"content": content,
	}
	result, err := s.client.Query(query, variables)
	if err != nil {
		return nil, err
	}
	return result["data"], nil
}

func (s *NoteService) DeleteNoteByID(id string) (interface{}, error) {
	query := `mutation ($id: Int!) {
		delete_notes(where: {id: {_eq: $id}}) {
			returning { id }
		}
	}`
	variables := map[string]interface{}{
		"id": id,
	}
	result, err := s.client.Query(query, variables)
	if err != nil {
		return nil, err
	}
	return result["data"], nil
}
