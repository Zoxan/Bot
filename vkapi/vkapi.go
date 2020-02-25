package vkapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

const (
	apiURL          = "api.vk.com"
	apiVer          = "5.103"
	apiUsersGet     = "method/users.get"
	apiMessagesSend = "method/messages.send"
)

//User describes vk user
type User struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Deacticated     string `json:"dedicated"`
	IsClosed        bool   `json:"is_closed"`
	CanAccessClosed bool   `json:"can_access_closed"`
}

//UsersResponse ..
type UsersResponse struct {
	Users []*User `json:"response"`
}

var errNoUsers = errors.New("no users")
var accessToken string

//Start ..
func Start(accToken string) {
	accessToken = accToken
}

//RequestUser ..
func RequestUser(userID int) (*User, error) {
	query := url.Values{}
	query.Add("user_ids", strconv.Itoa(userID))

	url := buildURL(apiUsersGet, query)

	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var response UsersResponse
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}

	if response.Users == nil || len(response.Users) == 0 {
		return nil, errNoUsers
	}

	return response.Users[0], nil
}

//RequestSendToGroup ..
func RequestSendToGroup(peerID int, randomID int, message string) error {
	query := url.Values{}
	query.Add("peer_id", strconv.Itoa(peerID))
	query.Add("random_id", strconv.Itoa(randomID))
	query.Add("message", message)

	url := buildURL(apiMessagesSend, query)

	_, err := http.DefaultClient.Get(url)
	return err
}

func buildURL(method string, query url.Values) string {

	query.Add("access_token", accessToken)
	query.Add("v", apiVer)

	url := url.URL{
		Scheme:   "https",
		Host:     apiURL,
		Path:     method,
		RawQuery: query.Encode(),
	}

	return url.String()
}
