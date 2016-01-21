package models

type Knot struct {
	Text string `json:"text"`
	Title string `json:"title"`
	Created int64 `json:"created"`
}

func (knot Knot) Validate() bool {
	return knot.Text != "" && knot.Title != ""
}
