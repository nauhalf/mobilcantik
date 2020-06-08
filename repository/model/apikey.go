package model

import (
	"time"
)

type APIKey struct {
	APIKeyID     string
	CreatedDate  time.Time
	ModifiedDate *time.Time
}
