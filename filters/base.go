package filters

import (
	"time"

	"github.com/thedataguild/faer/models"
)

type BaseTodoFilter struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Search    string
	SortKey   string
	SortDir   string
	Limit     int64
	Offset    int64
	CountOnly bool

	AccessingUserID            int64
	AccessingUserIsSystemAdmin bool
}

func (btf *BaseTodoFilter) SetAccessContext(accessContext *models.AccessContext) {
	if accessContext.User != nil {
		btf.AccessingUserID = accessContext.User.ID
	}
}
