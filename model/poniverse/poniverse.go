// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package oauthponiverse

import (
	"encoding/json"
	"github.com/mattermost/platform/einterfaces"
	"github.com/mattermost/platform/model"
	"io"
	"strconv"
)

const (
	USER_AUTH_SERVICE_PONIVERSE = "poniverse"
)

type PoniverseProvider struct {
}

type PoniverseUser struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

func init() {
	provider := &PoniverseProvider{}
	einterfaces.RegisterOauthProvider(USER_AUTH_SERVICE_PONIVERSE, provider)
}

func userFromPoniverseUser(poniverseUser *PoniverseUser) *model.User {
	user := &model.User{}
	username := poniverseUser.Username
	user.Username = model.CleanUsername(username)
	user.FirstName = poniverseUser.DisplayName
	user.Email = poniverseUser.Email
	user.AuthData = strconv.FormatInt(poniverseUser.GetId(), 10)
	user.AuthService = USER_AUTH_SERVICE_PONIVERSE

	return user
}

func poniverseUserFromJson(data io.Reader) *PoniverseUser {
	decoder := json.NewDecoder(data)
	var poniverseUser PoniverseUser
	err := decoder.Decode(&poniverseUser)
	if err == nil {
		return &poniverseUser
	} else {
		return nil
	}
}

func (poniverseUser *PoniverseUser) IsValid() bool {
	if poniverseUser.GetId() == 0 {
		return false
	}

	if len(poniverseUser.Email) == 0 {
		return false
	}

	return true
}

func (poniverseUser *PoniverseUser) getAuthData() string {
	return strconv.FormatInt(poniverseUser.GetId(), 10)
}

func (poniverseUser *PoniverseUser) GetId() int64 {
	if s, err := strconv.ParseInt(poniverseUser.Id, 10, 64); err == nil {
		return s;
	}

	return 0;
}

func (m *PoniverseProvider) GetIdentifier() string {
	return USER_AUTH_SERVICE_PONIVERSE
}

func (m *PoniverseProvider) GetUserFromJson(data io.Reader) *model.User {
	poniverseUser := poniverseUserFromJson(data)
	if poniverseUser.IsValid() {
		return userFromPoniverseUser(poniverseUser)
	}

	return &model.User{}
}

func (m *PoniverseProvider) GetAuthDataFromJson(data io.Reader) string {
	poniverseUser := poniverseUserFromJson(data)

	if poniverseUser.IsValid() {
		return poniverseUser.getAuthData()
	}

	return ""
}
