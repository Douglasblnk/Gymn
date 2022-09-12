package schemas

type RegisterEquipment struct {
	Name string `json:"name" validate:"required"`
}

type UpdateEquipment struct {
	Name *string `json:"name"`
}
