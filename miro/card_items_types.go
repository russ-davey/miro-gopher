package miro

import "time"

type CardItemData struct {
	Title       string    `json:"title"`
	AssigneeId  string    `json:"assigneeId"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
}

type CardItemStyle struct {
	CardTheme string `json:"cardTheme"`
}

type SetCardItem struct {
	Data     CardItemData  `json:"data"`
	Style    CardItemStyle `json:"style"`
	Position PositionSet   `json:"position"`
	Geometry Geometry      `json:"geometry"`
	Parent   ParentSet     `json:"parent"`
}

type CardItem struct {
	ID         string          `json:"id"`
	Data       CardItemData    `json:"data"`
	Style      CardItemStyle   `json:"style"`
	Position   Position        `json:"position"`
	Geometry   Geometry        `json:"geometry"`
	CreatedAt  time.Time       `json:"createdAt"`
	CreatedBy  BasicEntityInfo `json:"createdBy"`
	ModifiedAt time.Time       `json:"modifiedAt"`
	ModifiedBy BasicEntityInfo `json:"modifiedBy"`
	Parent     *Parent         `json:"parent,omitempty"`
	Links      Links           `json:"links"`
	Type       string          `json:"type"`
}
