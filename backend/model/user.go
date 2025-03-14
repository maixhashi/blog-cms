package model

import "time"

// User はデータベースのユーザーモデル
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserResponse はクライアントに返すユーザー情報
type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

// UserLoginRequest はログインリクエスト用の構造体
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserSignupRequest はサインアップリクエスト用の構造体
type UserSignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ToUser はUserSignupRequestからUserへの変換メソッド
func (r *UserSignupRequest) ToUser() User {
	return User{
		Email:    r.Email,
		Password: r.Password,
	}
}

// ToUserResponse はUserからUserResponseへの変換メソッド
func (u *User) ToUserResponse() UserResponse {
	return UserResponse{
		ID:    u.ID,
		Email: u.Email,
	}
}