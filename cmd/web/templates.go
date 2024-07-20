package main
import "snippetbox.anhnt2001/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}