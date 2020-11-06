package msgraph

type List struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	CreatedDateTime      string    `json:"createdDateTime"`
	LastModifiedDateTime string    `json:"lastModifiedDateTime"`
	List                 ListClass `json:"list"`
}

type ListClass struct {
	Hidden   bool   `json:"hidden"`
	Template string `json:"template"`
}
