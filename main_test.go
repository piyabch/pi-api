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

func TestCreateUser(t *testing.T) {
	user := model.User{
		FirstName: "Nattapol",
		LastName:  "Mankongprasit",
		Email:     "nat.man@gmail.com",
	}
	jsonValue, _ := json.Marshal(user)
	rsp, body, err := httpRequest(http.MethodPost, host+"/users", jsonValue)
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
	rsp, body, err := httpRequest(http.MethodPut, host+"/users", jsonValue)
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
	rsp, body, err := httpRequest(http.MethodGet, host+"/users/"+param, nil)
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
	rsp, body, err := httpRequest(http.MethodGet, host+"/users/search/?name="+searchName, nil)
	// check status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rsp.StatusCode)
	// check find result
	var resultUsers []model.User
	err = json.Unmarshal(body, &resultUsers)
	assert.NoError(t, err)
	assert.Greater(t, len(resultUsers), 0)
}

func httpRequest(method string, url string, data []byte) (*http.Response, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err == nil {
		body, _ := io.ReadAll(rsp.Body)
		return rsp, body, err
	}
	defer rsp.Body.Close()
	return nil, nil, err
}
