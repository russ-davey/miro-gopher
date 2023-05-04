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
	// Format Only custom frames are supported at the moment.
	Format Format `json:"format"`
	// Title of the frame. This title appears at the top of the frame.
	Title string `json:"title"`
	// Only free form frames are supported at the moment.
	Type Type `json:"type"`
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

type FrameSearchParams struct {
	// ParentItemID ID of the frame for which you want to retrieve the list of available items. (required)
	ParentItemID string `query:"parent_item_id"`
	// Limit The maximum number of results to return per call. If the number of items in the response is greater than
	// the limit specified, the response returns the cursor parameter with a value.
	// Default: 20.
	Limit string `query:"limit,omitempty"`
	// Type If you want to get a list of items of a specific type, specify an item type. For example,
	// if you want to retrieve the list of card items, set type to cards. Possible values: app_card, card, document,
	// embed, frame, image, shape, sticky_note, text
	Type ItemType `query:"type,omitempty"`
	// Cursor A cursor-paginated method returns a portion of the total set of results based on the limit specified and a
	// cursor that points to the next portion of the results. To retrieve the next portion of the collection, set the cursor
	// parameter equal to the cursor value you received in the response of the previous request.
	Cursor string `query:"cursor,omitempty"`
}

type ListFrames struct {
	Data   []FrameItem     `json:"data"`
	Total  int             `json:"total"`
	Size   int             `json:"size"`
	Cursor string          `json:"cursor"`
	Limit  int             `json:"limit"`
	Links  PaginationLinks `json:"links"`
}
