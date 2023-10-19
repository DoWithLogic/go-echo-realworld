package dtos

type (
	UserDetailResponse struct {
		Email    string `json:"email"`
		Token    string `json:"token"`
		UserName string `json:"username"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
	}
)
