package miro

import "time"

type BasicUserInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type User struct {
	BasicUserInfo
	Company   string    `json:"company"`
	Role      string    `json:"role"`
	Industry  string    `json:"industry"`
	Email     string    `json:"email"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"createdAt"`
	Picture   *Picture  `json:"picture"`
}
