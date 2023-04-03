package miro

import "time"

type BasicTeamInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Team struct {
	BasicTeamInfo
	CreatedAt  time.Time      `json:"createdAt"`
	ModifiedAt time.Time      `json:"modifiedAt"`
	CreatedBy  *BasicUserInfo `json:"createdBy"`
	ModifiedBy *BasicUserInfo `json:"modifiedBy"`
	Picture    *Picture       `json:"picture"`
}
