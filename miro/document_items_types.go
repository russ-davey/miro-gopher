package miro

import "time"

type SetDocumentItem struct {
	// Data Contains information about the document URL. (required)
	Data DocumentItemData `json:"data"`
	// Position Contains location information about the item, such as its x coordinate, y coordinate, and the origin of the x and y coordinates.
	Position PositionSet `json:"position"`
	// Geometry Contains geometrical information about the item, such as its width or height.
	Geometry Geometry `json:"geometry"`
	// Parent Contains information about the parent this item attached to. Passing null for ID will attach widget to the canvas directly.
	Parent ParentSet `json:"parent"`
}

type DocumentItemData struct {
	// DocumentURL A short text header to identify the document.
	DocumentURL string `json:"documentUrl"`
	// Title URL where the document is hosted. (required)
	Title string `json:"title"`
}

type DocumentItem struct {
	ID         string           `json:"id"`
	Data       DocumentItemData `json:"data"`
	Position   Position         `json:"position"`
	Geometry   Geometry         `json:"geometry"`
	CreatedAt  time.Time        `json:"createdAt"`
	CreatedBy  BasicEntityInfo  `json:"createdBy"`
	ModifiedAt time.Time        `json:"modifiedAt"`
	ModifiedBy BasicEntityInfo  `json:"modifiedBy"`
	Parent     Parent           `json:"parent"`
	Links      Links            `json:"links"`
	Type       string           `json:"type"`
}
