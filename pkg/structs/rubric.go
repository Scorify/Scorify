package structs

type RubricTemplateField struct {
	Name     string `json:"name"`
	MaxScore int    `json:"max_score"`
}

type RubricTemplate struct {
	Fields   []RubricTemplateField `json:"fields"`
	MaxScore int                   `json:"max_score"`
}

type RubricField struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Notes string `json:"notes"`
}

type Rubric struct {
	Fields []RubricField `json:"fields"`
	Notes  string        `json:"notes"`
}
