package templates

import (
	"github.com/ercsnmrs/rtm-go/internal/models"
)

// TemplateData defines template data
type TemplateData struct {
	CSRFToken       string
	IsAuthenticated bool
	PreferenceMap   map[string]string
	User            models.User
	Flash           string
	Warning         string
	Error           string
	GwVersion       string
}
