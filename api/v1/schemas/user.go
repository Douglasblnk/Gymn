package schemas

type RegisterUser struct {
	Name       string  `json:"name" validate:"required"`
	Email      string  `json:"email" validate:"required"`
	Password   string  `json:"password" validate:"required"`
	IsPersonal *bool   `json:"isPersonal"`
	Photo      *string `json:"photo"`
}

type UpdateUser struct {
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Password   *string `json:"password"`
	IsPersonal *bool   `json:"isPersonal"`
	Photo      *string `json:"photo"`
}
