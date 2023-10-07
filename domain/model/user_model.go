package model

type UserDetailModel struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type UserModel struct {
	Username  string          `json:"username"`
	Password  string          `json:"password"`
	Status    string          `json:"status"`
	Level     string          `json:"level"`
	Detail    UserDetailModel `json:"detail"`
	CreatedAt int64           `json:"created_at"`
	UpdatedAt int64           `json:"updated_at"`
	DeletedAt int64           `json:"deleted_at"`
}
