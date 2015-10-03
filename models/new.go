package models

/*
New model
*/
type New struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

/*
News Multiple news
*/
type News []*New
