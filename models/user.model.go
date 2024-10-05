package models

type User struct {
	ID       int32  `json:"id"`
	USERNAME string `json:"username"`
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
	ROLE_ID  int32  `json:"role_id"`
}

type Login struct {
	EMAIL    string `json:"email"`
	PASSWORD string `json:"password"`
}
