package dtos

type Satellites struct {
	Name     string   `json:"name" validate:"required"`
	Distance float64  `json:"distance" validate:"required"`
	Message  []string `json:"message" validate:"required,min=1,max=3"`
}
