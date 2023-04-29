package miro

import "time"

type ItemType string

const (
	AppCard    ItemType = "app_card"
	Card       ItemType = "card"
	Document   ItemType = "document"
	Embed      ItemType = "embed"
	Frame      ItemType = "frame"
	Image      ItemType = "image"
	Shape      ItemType = "shape"
	StickyNote ItemType = "sticky_note"
	Text       ItemType = "text"
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

type Parent struct {
	ID    string          `json:"id,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
}

type Geometry struct {
	Height   float64 `json:"height,omitempty"`
	Rotation float64 `json:"rotation,omitempty"`
	Width    float64 `json:"width,omitempty"`
}

type Position struct {
	Origin     Origin  `json:"origin"`
	RelativeTo string  `json:"relativeTo"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
}

type Style struct {
	BorderColor       string `json:"borderColor,omitempty"`
	BorderOpacity     string `json:"borderOpacity,omitempty"`
	BorderStyle       string `json:"borderStyle,omitempty"`
	BorderWidth       string `json:"borderWidth,omitempty"`
	Color             string `json:"color,omitempty"`
	FillColor         string `json:"fillColor,omitempty"`
	FillOpacity       string `json:"fillOpacity,omitempty"`
	FontFamily        string `json:"fontFamily,omitempty"`
	FontSize          string `json:"fontSize,omitempty"`
	TextAlign         string `json:"textAlign,omitempty"`
	TextAlignVertical string `json:"textAlignVertical,omitempty"`
}

type Item struct {
	CreatedAt time.Time       `json:"createdAt"`
	CreatedBy BasicEntityInfo `json:"createdBy"`
	Data      ItemData        `json:"data"`
	Geometry  Geometry        `json:"geometry"`
	ID        string          `json:"id"`
	Links     struct {
		Related string `json:"related,omitempty"`
		Self    string `json:"self"`
	} `json:"links"`
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

type Origin string

const Center Origin = "center"

type PositionUpdate struct {
	// Origin Area of the item that is referenced by its x and y coordinates. For example, an item with a center origin will
	// have its x and y coordinates point to its center. The center point of the board has x: 0 and y: 0 coordinates.
	// Currently, only one option is supported: center
	Origin Origin `json:"origin,omitempty"`
	// X-axis coordinate of the location of the item on the board.
	// By default, all items have absolute positioning to the board, not the current viewport.
	// The center point of the board has x: 0 and y: 0 coordinates.
	X float64 `json:"x,omitempty"`
	// Y-axis coordinate of the location of the item on the board.
	// By default, all items have absolute positioning to the board, not the current viewport.
	// The center point of the board has x: 0 and y: 0 coordinates.
	Y float64 `json:"y,omitempty"`
}

type ParentUpdate struct {
	ID string `json:"id"`
}

type ItemUpdate struct {
	// Parent Contains information about the parent this item attached to.
	// Passing null for ID will attach widget to the canvas directly.
	Parent *ParentUpdate `json:"parent,omitempty"`
	// Position Contains location information about the item, such as its x coordinate, y coordinate,
	// and the origin of the x and y coordinates.
	Position *PositionUpdate `json:"position,omitempty"`
}
