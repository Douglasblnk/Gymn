package schemas

type RegisterTrainingSheet struct {
	Name string `json:"name" validate:"required"`
}

type UpdateTrainingSheet struct {
	Name   string `json:"name"`
	Active string `json:"active"`
}
