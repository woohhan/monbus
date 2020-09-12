package busserver

type SkillRequest struct {
	Action SkillBodyAction `json:"action"`
}

type SkillBodyAction struct {
	Name         string       `json:"name"`
	DetailParams DetailParams `json:"detailParams"`
}

type DetailParams struct {
	IsWeekend isWeekend `json:"IsWeekend"`
}

type isWeekend struct {
	Value bool `json:"value"`
}

type SkillResponse struct {
	Version  string   `json:"version"`
	Template Template `json:"template"`
}

type Template struct {
	Outputs []SkillOutput `json:"outputs"`
}

type SkillOutput struct {
	SimpleText SimpleText `json:"simpleText"`
}

type SimpleText struct {
	Text string `json:"text"`
}
