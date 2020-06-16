// Copyright © 2020. Drew Lee. All rights reserved.

package contacts

import (
	"KarlMalone/api/routes/v1/chats"
	"KarlMalone/internal/models"
	"KarlMalone/internal/service"
	"KarlMalone/pkg/app"
	"KarlMalone/pkg/e"
	"KarlMalone/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

// add friend
func addFriend(c *gin.Context) {
	r := app.Gin{C: c}

	data := map[string]string{
		"owner_id": c.PostForm("owner_id"),
		"dst_id":   c.PostForm("dst_id"),
	}
	ownerId, _ := strconv.ParseInt(data["owner_id"], 10, 64)
	dstId, _ := strconv.ParseInt(data["dst_id"], 10, 64)

	friend := models.FindMember(dstId)
	if friend.ID == 0 {
		r.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_MEMBER, nil)
	}

	contact := service.Contact{
		OwnerId: ownerId,
		DstId:   dstId,
	}
	err := contact.AddFriend()
	if err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR_ADD_FRIEND, err)
	} else {
		r.Response(http.StatusOK, e.SUCCESS, nil)
	}
}

// load friend list
func loadFriend(c *gin.Context) {
	r := app.Gin{C: c}
	ownerId, _ := strconv.ParseInt(c.Query("owner_id"), 10, 64)
	contact := service.Contact{
		OwnerId: ownerId,
	}
	members := contact.FindAllFriend()
	r.Response(http.StatusOK, e.SUCCESS, members)
}

// create group
func createGroup(c *gin.Context) {
	r := app.Gin{C: c}
	group := service.Group{}
	if err := c.ShouldBind(&group); err != nil {
		logger.Error("routes v1 contacts bind group error", zap.Any("error", err))
	}

	resGroup, err := group.CreateGroup()
	if err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR_ADD_GROUP_FAIL, err)
	} else {
		r.Response(http.StatusOK, e.SUCCESS, resGroup)
	}
}

// join group
func joinGroup(c *gin.Context) {
	r := app.Gin{C: c}
	groupName := c.PostForm("name")
	ownerId, _ := strconv.ParseInt(c.PostForm("owner_id"), 10, 64)
	group := service.Group{
		Name: groupName,
	}
	resGroup := group.FindGroupByName()
	if resGroup.ID == 0 {
		r.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST_GROUP, nil)
		return
	}

	contact := service.Contact{
		OwnerId: ownerId,
		DstId:   resGroup.ID,
	}
	err := contact.JoinGroup()
	if err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR_JOIN_GROUP_FAIL, err)
	} else {
		chats.AddGroupId(ownerId, resGroup.ID) // refresh member group_set
		r.Response(http.StatusOK, e.SUCCESS, nil)
	}
}

// load group list
func loadGroup(c *gin.Context) {
	r := app.Gin{C: c}
	ownerId, _ := strconv.ParseInt(c.Query("owner_id"), 10, 64)
	contact := service.Contact{
		OwnerId: ownerId,
	}
	groups := contact.FindGroup()
	r.Response(http.StatusOK, e.SUCCESS, groups)
}
