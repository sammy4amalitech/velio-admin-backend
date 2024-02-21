package dao

type Menu struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	Price        int          `json:"price"`
	Description  string       `json:"description"`
	Availability Availability `json:"availability"`
	BaseModel
}

type Availability int

const (
	AVAILABLE = iota + 1
	NOT_AVAILABLE
	UNKNOWN
)
