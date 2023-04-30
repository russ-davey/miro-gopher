package miro

import "time"

type ItemType string

const (
	ItemTypeAppCard    ItemType = "app_card"
	ItemTypeCard       ItemType = "card"
	ItemTypeDocument   ItemType = "document"
	ItemTypeEmbed      ItemType = "embed"
	ItemTypeFrame      ItemType = "frame"
	ItemTypeImage      ItemType = "image"
	ItemTypeShape      ItemType = "shape"
	ItemTypeStickyNote ItemType = "sticky_note"
	ItemTypeText       ItemType = "text"
)

type ItemSearchParams struct {
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

type ItemData struct {
	Format      string `json:"format,omitempty"`
	ShowContent bool   `json:"showContent,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        string `json:"type,omitempty"`
	Content     string `json:"content,omitempty"`
	Shape       string `json:"shape,omitempty"`
}

type Item struct {
	CreatedAt  time.Time        `json:"createdAt"`
	CreatedBy  BasicEntityInfo  `json:"createdBy"`
	Data       ItemData         `json:"data"`
	Geometry   Geometry         `json:"geometry"`
	ID         string           `json:"id"`
	Links      Links            `json:"links"`
	ModifiedAt time.Time        `json:"modifiedAt"`
	ModifiedBy *BasicEntityInfo `json:"modifiedBy"`
	Parent     *Parent          `json:"parent,omitempty"`
	Position   Position         `json:"position"`
	Style      Style            `json:"style"`
	Type       string           `json:"type"`
}

type ListItems struct {
	Data   []Item          `json:"data"`
	Total  int             `json:"total"`
	Size   int             `json:"size"`
	Cursor string          `json:"cursor,omitempty"`
	Limit  int             `json:"limit"`
	Links  PaginationLinks `json:"links"`
	Type   string          `json:"type"`
}

type ItemUpdate struct {
	// Parent Contains information about the parent this item attached to.
	// Passing null for ID will attach widget to the canvas directly.
	Parent ParentSet `json:"parent"`
	// Position Contains location information about the item, such as its x coordinate, y coordinate,
	// and the origin of the x and y coordinates.
	Position PositionSet `json:"position"`
}
