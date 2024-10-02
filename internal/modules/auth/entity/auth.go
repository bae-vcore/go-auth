package entity

type (
	LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegisterReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
)
