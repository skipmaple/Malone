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

// @Summary add_friend
// @Description member add friend
// @Tags contacts
// @Accept multipart/form-data
// @Produce  json
// @Param owner_id formData string true "OwnerId"
// @Param dst_id formData string true "DstId"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/contacts/add_friend [post]
func addFriend(c *gin.Context) {
	r := app.Gin{C: c}

	data := map[string]string{
		"owner_id": c.PostForm("owner_id"),
		"dst_id":   c.PostForm("dst_id"),
	}
	ownerId, _ := strconv.ParseInt(data["owner_id"], 10, 64)
	dstId, _ := strconv.ParseInt(data["dst_id"], 10, 64)

	friend, err := models.FindMember(dstId)
	if err != nil {
		r.Response(http.StatusNotFound, e.ERROR_NOT_EXIST_MEMBER, nil)
		return
	}

	contact := service.Contact{
		OwnerId: ownerId,
		DstId:   friend.ID,
	}
	err = contact.AddFriend()
	if err != nil {
		r.Response(http.StatusInternalServerError, e.ERROR_ADD_FRIEND, err)
	} else {
		r.Response(http.StatusOK, e.SUCCESS, nil)
	}
}

// @Summary load_friend
// @Description load friend list
// @Tags contacts
// @Produce  json
// @Param owner_id query string true "OwnerId"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/contacts/load_friend [get]
func loadFriend(c *gin.Context) {
	r := app.Gin{C: c}
	ownerId, _ := strconv.ParseInt(c.Query("owner_id"), 10, 64)
	contact := service.Contact{
		OwnerId: ownerId,
	}
	members := contact.FindAllFriend()
	r.Response(http.StatusOK, e.SUCCESS, members)
}

// @Summary create_group
// @Description create group
// @Tags contacts
// @Accept multipart/form-data
// @Produce  json
// @Param owner_id formData string true "OwnerId"
// @Param name formData string true "GroupName"
// @Param icon formData string false "icon"
// @Param memo formData string false "memo"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/contacts/create_group [post]
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

// @Summary join_group
// @Description join group
// @Tags contacts
// @Accept multipart/form-data
// @Produce  json
// @Param owner_id formData string true "OwnerId"
// @Param name formData string true "GroupName"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/contacts/join_group [post]
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

// @Summary load_group
// @Description load group list
// @Tags contacts
// @Produce  json
// @Param owner_id query string true "OwnerId"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/contacts/load_group [get]
func loadGroup(c *gin.Context) {
	r := app.Gin{C: c}
	ownerId, _ := strconv.ParseInt(c.Query("owner_id"), 10, 64)
	contact := service.Contact{
		OwnerId: ownerId,
	}
	groups := contact.FindGroup()
	r.Response(http.StatusOK, e.SUCCESS, groups)
}

// @Summary Find group members
// @Description Find group members by group_id
// @Tags contacts
// @Produce  json
// @Param group_id query string true "GroupId"
// @Success 200 {object} app.Response
// @Failure 400 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /v1/contacts/find_group_members [get]
func findGroupMembers(c *gin.Context) {
	r := app.Gin{C: c}
	groupId, _ := strconv.ParseInt(c.Query("group_id"), 10, 64)
	group := service.Group{
		ID: groupId,
	}
	members := group.FindGroupMembersByGroupId()
	r.Response(http.StatusOK, e.SUCCESS, members)
}
