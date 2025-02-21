package news

type News struct {
	ID      int     `json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
	Image   *string `json:"image"`
}

type RequestNews struct {
	Title   string  `json:"title" binding:"required"`
	Content string  `json:"content" binding:"required"`
	Image   *string `json:"image"`
}

type EndpointInfo struct {
	Name     string            `json:"name"`
	Settings map[string]string `json:"settings"`
	Metadata Metadata          `json:"metadata"`
	Actions  map[string]Action `json:"actions"`
}

type Metadata struct {
	RoutingMap []RoutingMapEntry `json:"routingMap"`
}

type RoutingMapEntry struct {
	Action string `json:"action"`
	Route  string `json:"route"`
	Method string `json:"method"`
}

type Action struct {
	Params  Params `json:"params"`
	Handler string `json:"handler"`
}

type Params struct {
	PathParams  *ParamSchema `json:"pathParams,omitempty"`
	QueryParams *ParamSchema `json:"queryParams,omitempty"`
	Body        *ParamSchema `json:"body,omitempty"`
}

type ParamSchema struct {
	Type     string          `json:"type,omitempty"`
	Optional bool            `json:"optional,omitempty"`
	Props    map[string]Prop `json:"props,omitempty"`
}

type Prop struct {
	Type     string          `json:"type"`
	Optional bool            `json:"optional,omitempty"`
	Props    map[string]Prop `json:"props,omitempty"`
}
