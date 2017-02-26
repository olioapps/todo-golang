package models

import (
	olioModels "github.com/rachoac/service-skeleton-go/olio/common/models"
)

type AccessContext struct {
	SystemAccess bool
	User         *User
	RequestID    string
	Permissions  []*olioModels.Permission
}
