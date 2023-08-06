package response

type SignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func ToTokenResponse(accessToken string, refreshToken string) SignInResponse {
	response := SignInResponse{}
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken
	return response
}
