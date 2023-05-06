package miro

import "time"

type TextItemData struct {
	Content string `json:"content"`
}

type TextItem struct {
	ID         string          `json:"id"`
	Data       TextItemData    `json:"data"`
	Style      Style           `json:"style"`
	Position   Position        `json:"position"`
	Geometry   Geometry        `json:"geometry"`
	CreatedAt  time.Time       `json:"createdAt"`
	CreatedBy  BasicEntityInfo `json:"createdBy"`
	ModifiedAt time.Time       `json:"modifiedAt"`
	ModifiedBy BasicEntityInfo `json:"modifiedBy"`
	Parent     Parent          `json:"parent"`
	Links      Links           `json:"links"`
	Type       string          `json:"type"`
}

type TextItemSet struct {
	Data     TextItemData     `json:"data"`
	Style    Style            `json:"style"`
	Position PositionSet      `json:"position"`
	Geometry TextItemGeometry `json:"geometry"`
	Parent   *ParentSet       `json:"parent"`
}

type TextItemGeometry struct {
	Rotation float64 `json:"rotation,omitempty"`
	Width    float64 `json:"width,omitempty"`
}
