package dto

import "gymn/internal/models"

type SignInResponseDTO struct {
	AccessToken  *models.AccessToken
	RefreshToken string
	User         *UserDTO
}
