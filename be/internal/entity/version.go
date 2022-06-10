package entity

//Version represents application version
//swagger:response
type Version struct {
	Number string `json:"number"`
	Build  string `json:"build"`
}
