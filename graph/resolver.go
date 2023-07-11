package graph

import (
	"sync"

	"github.com/Rashad-Muntar/println/models"
	"github.com/jinzhu/gorm"
	"github.com/nats-io/nats.go"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
    DB *gorm.DB
    Nats *nats.Conn
	CreatedJob  *models.Job
    JobObservers map[string]chan *models.Job
    mu            sync.Mutex
}
