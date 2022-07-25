package dto

import "time"

type AccessToken struct {
	Token      string
	Expiration time.Time
}

type SignInResponseDTO struct {
	AccessToken  *AccessToken
	RefreshToken string
	User         *UserDTO
}

type RefreshAccessTokenDTO struct {
	AccessToken  *AccessToken
	RefreshToken string
}
