package userService

import (
	"gymn/internal/dto"
	"gymn/internal/exceptions"
	userRepository "gymn/internal/repository/user"
	"gymn/internal/security"
	"gymn/internal/utils"
	"gymn/v1/schemas"
)

func UpdateUser(id string, data *schemas.UpdateUser) (*dto.UserDTO, *utils.Error) {
	user, err := userRepository.FindUserByUID(id)

	if err != nil {
		return nil, err
	}

	if data.Name != nil {
		user.Name = *data.Name
	}

	if data.Email != nil {
		userWithSameEmail, err := userRepository.FindUserByEmail(*data.Email)

		if err != nil && err.Error != exceptions.ErrUserNotFound {
			return nil, err
		}

		if userWithSameEmail != nil && userWithSameEmail.ID != user.ID {
			return nil, utils.Throw(exceptions.ErrUserAlreadyExists, 400)
		}

		user.Email = *data.Email
	}

	if data.Password != nil {
		hashedPassword, err := security.HashPassword(*data.Password)

		if err != nil {
			return nil, err
		}

		user.Password = hashedPassword
	}

	// if data.Photo != nil {
	// 	pictureKey, err := storageProvider.SaveFile(data.Photo, "user")

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	if user.Photo != nil {
	// 		if err = storageProvider.DeleteFile(user.Photo); err != nil {
	// 			return nil, err
	// 		}
	// 	}

	// 	user.Photo = &pictureKey
	// }

	if data.IsPersonal != nil {
		user.Is_personal = *data.IsPersonal
	}

	if err = userRepository.UpdateUser(user); err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		Name:        user.Name,
		Email:       user.Email,
		Is_personal: user.Is_personal,
		Photo:       user.Photo,
	}, nil
}
