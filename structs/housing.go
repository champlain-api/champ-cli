package structs

type House struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Students uint16 `json:"students"`
	Distance string `json:"distance"`
	Address  string `json:"address"`
	ImageURL string `json:"imageURL"`
}
