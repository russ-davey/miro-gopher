package miro

import "io"

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
	Height   float64 `json:"height"`
	Rotation float64 `json:"rotation"`
	Width    float64 `json:"width"`
}

type Position struct {
	Origin     string  `json:"origin,omitempty"`
	RelativeTo string  `json:"relativeTo"`
	X          float64 `json:"x,omitempty"`
	Y          float64 `json:"y,omitempty"`
}

type Parent struct {
	ID    string          `json:"id,omitempty"`
	Links PaginationLinks `json:"links,omitempty"`
}

type ParentSet struct {
	ID string `json:"id"`
}

type (
	BorderStyle       string
	StrokeStyle       string
	TextAlign         string
	TextAlignVertical string
	TextOrientation   string
)

const (
	BorderStyleNormal BorderStyle = "normal"
	BorderStyleDotted BorderStyle = "dotted"
	BorderStyleDashed BorderStyle = "dashed"

	StrokeStyleNormal StrokeStyle = "normal"
	StrokeStyleDotted StrokeStyle = "dotted"
	StrokeStyleDashed StrokeStyle = "dashed"

	TextAlignCenter TextAlign = "center"
	TextAlignLeft   TextAlign = "left"
	TextAlignRight  TextAlign = "right"

	TextAlignVerticalTop    TextAlignVertical = "top"
	TextAlignVerticalMiddle TextAlignVertical = "middle"
	TextAlignVerticalBottom TextAlignVertical = "bottom"

	TextOrientationHorizontal TextOrientation = "horizontal"
	TextOrientationAligned    TextOrientation = "aligned"
)

type Style struct {
	BorderColor   string      `json:"borderColor,omitempty"`
	BorderOpacity string      `json:"borderOpacity,omitempty"`
	BorderStyle   BorderStyle `json:"borderStyle,omitempty"`
	BorderWidth   string      `json:"borderWidth,omitempty"`
	Color         string      `json:"color,omitempty"`
	FillColor     string      `json:"fillColor,omitempty"`
	FillOpacity   string      `json:"fillOpacity,omitempty"`
	// FontFamily Defines the font type for the text. Default: arial.
	FontFamily string `json:"fontFamily,omitempty"`
	// FontSize Defines the font size, in dp, for the text. Default: 14, minimum: 10, maximum 288.
	FontSize          string            `json:"fontSize,omitempty"`
	TextAlign         TextAlign         `json:"textAlign,omitempty"`
	TextAlignVertical TextAlignVertical `json:"textAlignVertical,omitempty"`
}

type Origin string

const Center Origin = "center"

type PositionSet struct {
	// Origin Area of the item that is referenced by its x and y coordinates. For example, an item with a center origin will
	// have its x and y coordinates point to its center. The center point of the board has x: 0 and y: 0 coordinates.
	// Currently, only one option is supported: center
	Origin Origin `json:"origin,omitempty"`
	// X-axis coordinate of the location of the item on the board.
	// By default, all items have absolute positioning to the board, not the current viewport.
	// The center point of the board has x: 0 and y: 0 coordinates.
	X float64 `json:"x"`
	// Y-axis coordinate of the location of the item on the board.
	// By default, all items have absolute positioning to the board, not the current viewport.
	// The center point of the board has x: 0 and y: 0 coordinates.
	Y float64 `json:"y"`
}

type GeometrySet struct {
	Height float64 `json:"height,omitempty"`
	Width  float64 `json:"width,omitempty"`
}

type Links struct {
	Related string `json:"related,omitempty"`
	Self    string `json:"self,omitempty"`
}

type UploadFileItem struct {
	// Title URL where the document is hosted. (required)
	Title string `json:"title"`
	// Position Contains location information about the item, such as its x coordinate, y coordinate, and the origin of the x and y coordinates.
	Position PositionSet `json:"position"`
	// Geometry Contains geometrical information about the item, such as its width or height.
	Geometry Geometry `json:"geometry"`
	// Parent Contains information about the parent this item attached to. Passing null for ID will attach widget to the canvas directly.
	Parent *ParentSet `json:"parent,omitempty"`
}

type MultiPart struct {
	Reader      io.Reader
	FileName    string
	ContentType string
}

type MultiParts map[string]MultiPart
