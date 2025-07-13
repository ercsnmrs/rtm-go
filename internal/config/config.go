package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
	"github.com/ercsnmrs/rtm-go/internal/channeldata"
	"github.com/ercsnmrs/rtm-go/internal/driver"
	"github.com/ercsnmrs/rtm-go/internal/models"
	"github.com/robfig/cron/v3"
)

// AppConfig holds application configuration
type AppConfig struct {
	DB            *driver.DB
	Session       *scs.SessionManager
	InProduction  bool
	Domain        string
	MonitorMap    map[int]cron.EntryID
	PreferenceMap map[string]string
	Scheduler     *cron.Cron
	WsClient      models.WSClient
	PusherSecret  string
	TemplateCache map[string]*template.Template
	MailQueue     chan channeldata.MailJob
	Version       string
	Identifier    string
}
