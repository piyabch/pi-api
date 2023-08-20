package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/piyabch/pi-api/config"
	"github.com/piyabch/pi-api/model"
	"github.com/stretchr/testify/assert"
)

var host string
var token string

func init() {
	webAddr := config.App.WebAddress
	hostname := "localhost"
	port := "8080"
	if webAddr != "" {
		addr := strings.Split(webAddr, ":")
		if len(addr) > 1 {
			if addr[0] != "" {
				host = addr[0]
			}
			if addr[1] != "" {
				port = addr[1]
			}
		}
	}
	host = "http://" + hostname + ":" + port
}

func TestAuth(t *testing.T) {
	authData := model.AuthData{
		Email:    "admin@example.com",
		Password: "defaultpassword",
	}
	jsonValue, _ := json.Marshal(authData)
	rsp, body, err := httpRequest(http.MethodPost, host+"/auth", jsonValue, "")
	// check status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	// check token result
	var authResult model.AuthResult
	err = json.Unmarshal(body, &authResult)
	assert.NoError(t, err)
	assert.NotEmpty(t, authResult.Token)
	// keep token for the next call
	token = authResult.Token
}

func TestCreateUser(t *testing.T) {
	user := model.User{
		FirstName: "Nattapol",
		LastName:  "Mankongprasit",
		Email:     "nat.man@gmail.com",
	}
	jsonValue, _ := json.Marshal(user)
	rsp, body, err := httpRequest(http.MethodPost, host+"/users", jsonValue, token)
	// check status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	// check inserted id
	var resultUser model.User
	err = json.Unmarshal(body, &resultUser)
	assert.NoError(t, err)
	assert.Greater(t, resultUser.ID, 0)
}

func TestUpdateUser(t *testing.T) {
	user := model.User{
		ID:        1,
		FirstName: "-",
		LastName:  "-",
		Email:     "-",
	}
	jsonValue, _ := json.Marshal(user)
	rsp, body, err := httpRequest(http.MethodPut, host+"/users", jsonValue, token)
	// check status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	// check update result
	var resultUser model.User
	err = json.Unmarshal(body, &resultUser)
	assert.NoError(t, err)
	assert.Equal(t, resultUser.FirstName, "-")
}

func TestFindUserById(t *testing.T) {
	userId := 1
	param := strconv.Itoa(userId)
	rsp, body, err := httpRequest(http.MethodGet, host+"/users/"+param, nil, token)
	// check status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	// check find result
	var resultUser model.User
	err = json.Unmarshal(body, &resultUser)
	assert.NoError(t, err)
	assert.Equal(t, resultUser.ID, userId)
}

func TestFindUsersByName(t *testing.T) {
	searchName := "wee"
	rsp, body, err := httpRequest(http.MethodGet, host+"/users/search/?name="+searchName, nil, token)
	// check status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	// check find result
	var resultUsers []model.User
	err = json.Unmarshal(body, &resultUsers)
	assert.NoError(t, err)
	assert.Greater(t, len(resultUsers), 0)
}

func httpRequest(method string, url string, data []byte, authToken string) (*http.Response, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if authToken != "" {
		req.Header.Set("Authorization", authToken)
	}
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err == nil {
		body, _ := io.ReadAll(rsp.Body)
		return rsp, body, err
	}
	defer rsp.Body.Close()
	return nil, nil, err
}
