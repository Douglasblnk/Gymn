package schemas

type RegisterStudent struct {
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	Birth     *string `json:"birth"`
	Color     *string `json:"color"`
	Weight    string  `json:"weight" validate:"required"`
	Height    string  `json:"height" validate:"required"`
}

type UpdateStudent struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Birth     *string `json:"birth"`
	Color     *string `json:"color"`
	Weight    *string `json:"weight"`
	Height    *string `json:"height"`
}
