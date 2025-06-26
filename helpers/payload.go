package helpers

type Payload struct {
	Page     int      `json:"page"`
	Length   int      `json:"length"`
	Orders   []string `json:"orders"`
	SortType string   `json:"sortType"`
	Filters  []Filter `json:"filters"`
}

type Filter struct {
	Field              string `json:"field"`
	Value              string `json:"value"`
	Type               string `json:"type"`
	ComparisonOperator string `json:"comparisonOperator"`
}
