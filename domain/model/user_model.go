package model

// UserDetailModel is a model for use detail on collection
type UserDetailModel struct {
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Email   string `json:"email" bson:"email"`
	Phone   string `json:"phone" bson:"phone"`
}

// UserModel is a model for users collection
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

// UserCreateRequest is a body to create new user
type UserCreateRequest struct {
	Username string `json:"username" validate:"required,min=6"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,min=6"`
}

// UserUpdatePassRequest is a body to update password
type UserUpdatePassRequest struct {
	Password    string `json:"password" validate:"required,min=8"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

type UserUpdateDetailRequest struct {
	Name    string `json:"name" validate:"required,min=3"`
	Address string `json:"address" validate:"required,min=15"`
	Email   string `json:"email" validate:"required,min=6"`
	Phone   string `json:"phone" validate:"required,min=11"`
}
