package miro

import "time"

type Shape string

const (
	ShapeRectangle                  Shape = "rectangle"
	ShapeRoundRectangle             Shape = "round_rectangle"
	ShapeCircle                     Shape = "circle"
	ShapeTriangle                   Shape = "triangle"
	ShapeRhombus                    Shape = "rhombus"
	ShapeParallelogram              Shape = "parallelogram"
	ShapeTrapezoid                  Shape = "trapezoid"
	ShapePentagon                   Shape = "pentagon"
	ShapeHexagon                    Shape = "hexagon"
	ShapeOctagon                    Shape = "octagon"
	ShapeWedgeRoundRectangleCallout Shape = "wedge_round_rectangle_callout"
	ShapeStar                       Shape = "star"
	ShapeFlowChartPredefinedProcess Shape = "flow_chart_predefined_process"
	ShapeCloud                      Shape = "cloud"
	ShapeCross                      Shape = "cross"
	ShapeCan                        Shape = "can"
	ShapeRightArrow                 Shape = "right_arrow"
	ShapeLeftArrow                  Shape = "left_arrow"
	ShapeLeftRightArrow             Shape = "left_right_arrow"
	ShapeLeftBrace                  Shape = "left_brace"
	ShapeRightBrace                 Shape = "right_brace"
)

type ShapeItemData struct {
	// Shape Defines the geometric shape of the item when it is rendered on the board.
	Shape Shape `json:"shape"`
	// Content The text you want to display on the shape.
	Content string `json:"content"`
}

type CreateShapeItem struct {
	// Data Contains shape item data, such as the content or shape type of the shape.
	Data ShapeItemData `json:"data"`
	// Style Contains information about the shape style, such as the border color or opacity.
	Style Style `json:"style"`
	// Position Contains location information about the item, such as its x coordinate, y coordinate, and the origin of the x and y coordinates.
	Position PositionUpdate `json:"position"`
	// Geometry Contains geometrical information about the item, such as its width or height.
	Geometry Geometry `json:"geometry"`
	// Parent Contains information about the parent this item attached to. Passing null for ID will attach widget to the canvas directly.
	Parent ParentUpdate `json:"parent"`
}

type ShapeItem struct {
	ID         string          `json:"id"`
	Data       ShapeItemData   `json:"data"`
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
