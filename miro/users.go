package miro

import "time"

type User struct {
	BasicEntityInfo
	Company   string    `json:"company"`
	Role      string    `json:"role"`
	Industry  string    `json:"industry"`
	Email     string    `json:"email"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"createdAt"`
	Picture   *Picture  `json:"picture"`
}
