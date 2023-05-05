package miro

import "time"

type (
	Format string
	Type   string
)

const (
	FormatCustom Format = "custom"
	TypeFreeform Type   = "freeform"
)

type SetFrameItem struct {
	// Data Contains frame item data, such as the title, frame type, or frame format.
	Data FrameItemData `json:"data"`
	// Style Contains information about the style of a frame item, such as the fill color.
	Style Style `json:"style"`
	// Position Contains location information about the item, such as its x coordinate, y coordinate, and the origin of the x and y coordinates.
	Position PositionSet `json:"position"`
	// Geometry Contains geometrical information about the item, such as its width or height.
	Geometry GeometrySet `json:"geometry"`
}

type FrameItemData struct {
	// Format Only custom frames are supported at the moment. (required)
	Format Format `json:"format"`
	// Title of the frame. This title appears at the top of the frame.
	Title string `json:"title"`
	// Only free form frames are supported at the moment. (required)
	Type        Type `json:"type"`
	ShowContent bool `json:"showContent,omitempty"`
}

type FrameItem struct {
	ID         string          `json:"id"`
	Data       FrameItemData   `json:"data"`
	Style      Style           `json:"style"`
	Position   Position        `json:"position"`
	Geometry   Geometry        `json:"geometry"`
	CreatedAt  time.Time       `json:"createdAt"`
	CreatedBy  BasicEntityInfo `json:"createdBy"`
	ModifiedAt time.Time       `json:"modifiedAt"`
	ModifiedBy BasicEntityInfo `json:"modifiedBy"`
	Links      Links           `json:"links"`
	Type       string          `json:"type"`
	Parent     *Parent         `json:"parent,omitempty"`
}
