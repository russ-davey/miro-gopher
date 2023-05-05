package miro

import "time"

type ImageItemData struct {
	// ImageURL URL of the image.
	ImageURL string `json:"imageUrl,omitempty"`
	// Title A short text header to identify the image.
	Title string `json:"title,omitempty"`
}

type ImageItem struct {
	ID         string          `json:"id"`
	Data       ImageItemData   `json:"data"`
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

type ImageItemSet struct {
	// Data Contains information about the document URL. (required)
	Data ItemDataSet `json:"data"`
	// Position Contains location information about the item, such as its x coordinate, y coordinate, and the origin of the x and y coordinates.
	Position PositionSet `json:"position"`
	// Geometry Contains geometrical information about the item, such as its width or height.
	Geometry GeometrySet `json:"geometry"`
	// Parent Contains information about the parent this item attached to. Passing null for ID will attach widget to the canvas directly.
	Parent ParentSet `json:"parent"`
}
