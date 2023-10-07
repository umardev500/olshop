package model

type UserDetailModel struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type UserModel struct {
	Username  string          `json:"username" bson:"username"`
	Password  string          `json:"password" bson:"password"`
	Status    string          `json:"status" bson:"status"`
	Level     string          `json:"level" bson:"level"`
	Detail    UserDetailModel `json:"detail" bson:"detail"`
	CreatedAt int64           `json:"created_at" bson:"created_at"`
	UpdatedAt int64           `json:"updated_at" bson:"updated_at"`
	DeletedAt int64           `json:"deleted_at" bson:"deleted_at"`
}
