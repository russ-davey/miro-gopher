package miro

import "time"

type ListConnectors struct {
	Cursor string          `json:"cursor"`
	Data   []Connector     `json:"data"`
	Limit  int             `json:"limit"`
	Links  PaginationLinks `json:"links"`
	Size   int             `json:"size"`
	Total  int             `json:"total"`
}

type ConnectorSearchParams struct {
	// Limit The maximum number of results to return per call. If the number of connectors in the response is greater
	// than the limit specified, the response returns the cursor parameter with a value.
	// Default: 10
	Limit string `query:"limit,omitempty"`
	// Cursor A cursor-paginated method returns a portion of the total set of results based on the limit specified and a
	// cursor that points to the next portion of the results. To retrieve the next portion of the collection, set the
	// cursor parameter equal to the cursor value you received in the response of the previous request.
	Cursor string `query:"cursor,omitempty"`
}

type Connector struct {
	Captions    []Caption       `json:"captions"`
	CreatedAt   time.Time       `json:"createdAt"`
	CreatedBy   BasicEntityInfo `json:"createdBy"`
	EndItem     ConnectorItem   `json:"endItem"`
	ID          string          `json:"id"`
	IsSupported bool            `json:"isSupported"`
	Links       Links           `json:"links"`
	ModifiedAt  time.Time       `json:"modifiedAt"`
	ModifiedBy  BasicEntityInfo `json:"modifiedBy"`
	Shape       ConnectorShape  `json:"shape"`
	StartItem   ConnectorItem   `json:"startItem"`
	Style       ConnectorStyle  `json:"style"`
	Type        string          `json:"type"`
}

type SetConnector struct {
	// StartItem The end point of the connector. endItem.id must be different from startItem.id
	StartItem SetConnectorItem `json:"startItem"`
	// EndItem The end point of the connector. endItem.id must be different from startItem.id
	EndItem SetConnectorItem `json:"endItem"`
	// Captions Contains information about the style of a connector, such as the color or caption font size
	Captions []Caption `json:"captions"`
	// Style Contains information about the style of a connector, such as the color or caption font size
	Style ConnectorStyle `json:"style"`
	// Shape The path type of the connector line, defines curvature. Default: curved.
	Shape ConnectorShape `json:"shape"`
}

type Caption struct {
	// Content The text you want to display on the connector. Supports inline HTML tags. (Required)
	Content string `json:"content"`
	// Position The relative position of the text on the connector, in percentage, minimum 0%, maximum 100%. With 50% value,
	// the text will be placed in the middle of the connector line. Default: 50%
	Position string `json:"position"`
	// TextAlignVertical The vertical position of the text on the connector. Default: middle
	TextAlignVertical TextAlignVertical `json:"textAlignVertical"`
}

type ConnectorStyle struct {
	// Color Hex value representing the color for the captions on the connector. Default: #1a1a1a
	Color string `json:"color"`
	// EndStrokeCap The decoration cap of the connector end, like an arrow or circle. Default: stealth.
	EndStrokeCap StrokeCap `json:"endStrokeCap"`
	// FontSize Defines the font size, in dp, for the captions on the connector. Default: 14
	FontSize string `json:"fontSize"`
	// StartStrokeCap The decoration cap of the connector end, like an arrow or circle. Default: none.
	StartStrokeCap StrokeCap `json:"startStrokeCap"`
	// strokeColor Hex value of the color of the connector line. Default: #000000.
	StrokeColor string `json:"strokeColor"`
	// strokeStyle The stroke pattern of the connector line. Default: normal.
	StrokeStyle StrokeStyle `json:"strokeStyle"`
	// strokeWidth The thickness of the connector line, in dp. Default: 1.0.
	StrokeWidth string `json:"strokeWidth"`
	// textOrientation The captions orientation relatively to the connector line curvature. Default: aligned.
	TextOrientation TextOrientation `json:"textOrientation"`
}

type ConnectorPosition struct {
	X string `json:"x"`
	Y string `json:"y"`
}

type ConnectorItem struct {
	ID       string            `json:"id"`
	Links    Links             `json:"links"`
	Position ConnectorPosition `json:"position"`
}

type (
	SnapTo         string
	ConnectorShape string
	StrokeCap      string
)

const (
	SnapToAuto   SnapTo = "auto"
	SnapToTop    SnapTo = "top"
	SnapToRight  SnapTo = "right"
	SnapToLeft   SnapTo = "left"
	SnapToBottom SnapTo = "bottom"

	ConnectorShapeCurved   ConnectorShape = "curved"
	ConnectorShapeStraight ConnectorShape = "straight"
	ConnectorShapeElbowed  ConnectorShape = "elbowed"

	StrokeCapStealth        StrokeCap = "stealth"
	StrokeCapDiamond        StrokeCap = "diamond"
	StrokeCapDiamondFilled  StrokeCap = "diamond_filled"
	StrokeCapOval           StrokeCap = "oval"
	StrokeCapOvalFilled     StrokeCap = "oval_filled"
	StrokeCapArrow          StrokeCap = "arrow"
	StrokeCapTriangle       StrokeCap = "triangle"
	StrokeCapTriangleFilled StrokeCap = "triangle_filled"
	StrokeCapErdOne         StrokeCap = "erd_one"
	StrokeCapErdMany        StrokeCap = "erd_many"
	StrokeCapErdOnlyOne     StrokeCap = "erd_only_one"
	StrokeCapErdZeroOrOne   StrokeCap = "erd_zero_or_one"
	StrokeCapErdOneOrMany   StrokeCap = "erd_one_or_many"
)

type SetConnectorItem struct {
	// Position The relative position of the point on an item where the connector is attached. Position
	// with x=0% and y=0% correspond to the top-left corner of the item, x=100% and y=100% correspond to the right-bottom corner.
	Position ConnectorPosition `json:"position"`
	// ID Unique identifier (ID) of the item to which you want to attach the connector. Note that Frames are not supported at the moment.
	ID string `json:"id"`
	// SnapTo The side of the item connector should be attached to, the connection point will be placed in the middle of that side.
	// Option auto allows to pick a connection point automatically. Only either position or snapTo parameter is allowed to be set, if neither provided
	// snapTo: auto will be used by default.
	SnapTo SnapTo `json:"snapTo"`
}
