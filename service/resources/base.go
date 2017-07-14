package resources

import (
	"errors"

	"github.com/olioapps/todo-golang/filters"
	"github.com/gin-gonic/gin"
	olioResources "github.com/olioapps/service-skeleton-go/olio/service/resources"
)

type BaseTodoResource struct {
	olioResources.BaseResource
}

// func (self *BaseSIMResource) getAccessContext(c *gin.Context) *models.AccessContext {
// 	user := self.getCurrentUser(c)
// 	requestId, exists := c.Get("Request-Id")
// 	if !exists {
// 		requestId = ""
// 	}

// 	accessContext := &models.AccessContext{User: user, RequestID: requestId.(string)}
// 	token, exists := c.Get("JWT_TOKEN")
// 	if exists {
// 		accessContext.Permissions = olioUtil.TokenToPermissions(token.(string))
// 	}

// 	return accessContext
// }

// func (self *BaseSIMResource) getCurrentUser(c *gin.Context) *models.User {
// 	user, exists := c.Get("currentUser")
// 	if !exists {
// 		return nil
// 	}

// 	return user.(*models.User)
// }

func (self *BaseTodoResource) parseFilter(c *gin.Context, filter *filters.BaseTodoFilter) error {
	// todo - return error instead of panic if invalid parse
	if id := self.ParseInt(c, "id"); id > 0 {
		filter.ID = id
	}

	limitStr := self.ParseString(c, "limit")
	offsetStr := self.ParseString(c, "offset")
	if limitStr != "" || offsetStr != "" {
		offset := self.ParseInt(c, "offset")
		limit := self.ParseInt(c, "limit")
		if limit > 100 {
			limit = 100
		} else if limit < 1 {
			limit = 20
		}
		if offset < 0 {
			offset = 0
		}
		filter.Limit = limit
		filter.Offset = offset
	}

	if search := self.ParseString(c, "search"); search != "" {
		filter.Search = search
	}
	if sortKey := self.ParseString(c, "sortkey"); sortKey != "" {
		filter.SortKey = sortKey
	}
	if sortDir := self.ParseString(c, "sortdir"); sortDir != "" {
		if !(sortDir == "asc" || sortDir == "desc") {
			return errors.New("Invalid sort direction")
		}
		filter.SortDir = sortDir
	}
	return nil
}
