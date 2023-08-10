package model

import (
	"encoding/json"
)

type User struct {
	ID       string   `json:"_id" csv:"_id"`
	Index    int      `json:"index" csv:"index"`
	GUID     string   `json:"guid" csv:"guid"`
	IsActive bool     `json:"isActive" csv:"isActive"`
	Balance  string   `json:"balance" csv:"balance"`
	Tags     []string `json:"tags" csv:"tags"`
	Friends  []Friend `json:"friends" csv:"friends"`
}

// UsersFromJson get []User from json body
func UsersFromJson(body []byte) (users []User, err error) {
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return
}
