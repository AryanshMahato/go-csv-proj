package model

import (
	"encoding/json"
)

type User struct {
	ID       string   `json:"_id"`
	Index    int      `json:"index"`
	GUID     string   `json:"guid"`
	IsActive bool     `json:"isActive"`
	Balance  string   `json:"balance"`
	Tags     []string `json:"tags"`
	Friends  []Friend `json:"friends"`
}

// UsersFromJson get []User from json body
func UsersFromJson(body []byte) (users []User, err error) {
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return
}
