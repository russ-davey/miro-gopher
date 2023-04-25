package miro

import "time"

type Team struct {
	BasicEntityInfo
	CreatedAt  time.Time        `json:"createdAt"`
	ModifiedAt time.Time        `json:"modifiedAt"`
	CreatedBy  *BasicEntityInfo `json:"createdBy"`
	ModifiedBy *BasicEntityInfo `json:"modifiedBy"`
	Picture    *Picture         `json:"picture"`
}
