package entities

type Basket struct {
	Username string `json:"name"`
	Pets     []Pet  `json:"pets"`
}
