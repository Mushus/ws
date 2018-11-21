package model

type ApartmentParam struct {
	Name string `json:"name" form:"name" query:"name" validate:"required"`
}

func (m ApartmentParam) ToDBModel() Apartment {
	return Apartment{
		ID:   0,
		Name: m.Name,
	}
}
