package response

import "time"

type ReissueTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessExp    string `json:"accessExp"`
	RefreshExp   string `json:"refreshExp"`
}

func ToReissueTokenResponse(accessToken string, refreshToken string, accessExp time.Time, refreshExp time.Time) ReissueTokenResponse {
	response := ReissueTokenResponse{}
	response.AccessToken = accessToken
	response.RefreshToken = refreshToken
	response.AccessExp = accessExp.String()
	response.RefreshExp = refreshExp.String()
	return response
}
