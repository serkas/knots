package models

type Knot struct {
	Text string `json:"text"`
	Title string `json:"title"`
}

func (knot Knot) Validate() bool {
	return knot.Text != "" && knot.Title != ""
}
