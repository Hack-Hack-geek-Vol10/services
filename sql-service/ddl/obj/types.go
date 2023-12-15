package obj

type Object struct {
	RawData   string      `json:"raw_data"`
	Tables    []*Table    `json:"tables"`
	Relations []*Relation `json:"relations"`
	Enums     []*Enum     `json:"enums"`
}

type Table struct {
	Name    string    `json:"name"`
	Columns []*Column `json:"columns"`
}

type Column struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
}

type Relation struct {
	FromCol string `json:"from_col"`
	ToCol   string `json:"to_col"`
}

type Enum struct {
	Name   string   `json:"name"`
	Fields []string `json:"fields"`
}
