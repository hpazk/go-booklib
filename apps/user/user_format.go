package user

func userLoginResponseFormatter(user User, authToken string) loginResponse {
	formatter := loginResponse{
		ID:        user.ID,
		Name:      user.Name,
		Address:   user.Address,
		Photo:     user.Photo,
		Email:     user.Email,
		AuthToken: authToken,
	}

	return formatter
}

func userResponseFormatter(user User, authToken string) response {
	formatter := response{
		ID:              user.ID,
		Name:            user.Name,
		Address:         user.Address,
		Photo:           user.Photo,
		Email:           user.Email,
		Role:            user.Role,
		EmailVerifiedAt: user.EmailVerifiedAt,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		DeletedAt:       user.DeletedAt,
		AuthToken:       authToken,
	}

	return formatter
}
