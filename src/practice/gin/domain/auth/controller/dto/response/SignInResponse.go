package response

import "time"

type SignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessExp    string `json:"accessExp"`
	RefreshExp   string `json:"refreshExp"`
}

func ToTokenResponse(accessToken string, refreshToken string, accessExp time.Time, refreshExp time.Time) SignInResponse {
	response := SignInResponse{}
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken
	response.AccessExp = accessExp.String()
	response.RefreshExp = refreshExp.String()
	return response
}
