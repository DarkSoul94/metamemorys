package http

type UserRegistration struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Pass       string `json:"pass"`
	PassRepeat string `json:"pass_repeat"`
}

type UserAuth struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
