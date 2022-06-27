package schemas

type User struct {
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate:"required"`
	IsPersonal bool   `json:"isPersonal"`
	Photo      string `json:"photo"`
}
