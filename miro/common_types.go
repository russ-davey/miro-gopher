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
