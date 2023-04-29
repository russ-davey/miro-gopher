package miro

// BasicEntityInfo info type for different entities (i.e. users, teams & organizations)
type BasicEntityInfo struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type PaginationLinks struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Next  string `json:"next,omitempty"`
	Prev  string `json:"prev,omitempty"`
	Self  string `json:"self,omitempty"`
}

type Geometry struct {
	Height   float64 `json:"height,omitempty"`
	Rotation float64 `json:"rotation,omitempty"`
	Width    float64 `json:"width,omitempty"`
}

type Position struct {
	Origin     string  `json:"origin,omitempty"`
	RelativeTo string  `json:"relativeTo"`
	X          float64 `json:"x,omitempty"`
	Y          float64 `json:"y,omitempty"`
}

type (
	BorderStyle       string
	TextAlign         string
	TextAlignVertical string
)

const (
	BorderStyleNormal BorderStyle = "normal"
	BorderStyleDotted BorderStyle = "dotted"
	BorderStyleDashed BorderStyle = "dashed"

	TextAlignCenter TextAlign = "center"
	TextAlignLeft   TextAlign = "left"
	TextAlignRight  TextAlign = "right"

	TextAlignVerticalTop    TextAlignVertical = "top"
	TextAlignVerticalMiddle TextAlignVertical = "middle"
	TextAlignVerticalBottom TextAlignVertical = "bottom"
)

type Style struct {
	BorderColor       string            `json:"borderColor,omitempty"`
	BorderOpacity     string            `json:"borderOpacity,omitempty"`
	BorderStyle       BorderStyle       `json:"borderStyle,omitempty"`
	BorderWidth       string            `json:"borderWidth,omitempty"`
	Color             string            `json:"color,omitempty"`
	FillColor         string            `json:"fillColor,omitempty"`
	FillOpacity       string            `json:"fillOpacity,omitempty"`
	FontFamily        string            `json:"fontFamily,omitempty"`
	FontSize          string            `json:"fontSize,omitempty"`
	TextAlign         TextAlign         `json:"textAlign,omitempty"`
	TextAlignVertical TextAlignVertical `json:"textAlignVertical,omitempty"`
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

type Links struct {
	Related string `json:"related,omitempty"`
	Self    string `json:"self,omitempty"`
}
