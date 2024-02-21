package dao

type Ingredient struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Description string `json:"description"`
	BaseModel
}
