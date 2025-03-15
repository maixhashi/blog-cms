package model

import "time"

// データベースモデル
type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"-" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

// リクエスト用の構造体
type TaskRequest struct {
	Title  string `json:"title" validate:"required,max=100"`
	UserId uint   `json:"-"` // クライアントからは送信されず、JWTから取得
}

// レスポンス用の構造体
type TaskResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TaskからTaskResponseへの変換メソッド
func (t *Task) ToResponse() TaskResponse {
	return TaskResponse{
		ID:        t.ID,
		Title:     t.Title,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

// TaskRequestからTaskへの変換メソッド
func (tr *TaskRequest) ToModel() Task {
	return Task{
		Title:  tr.Title,
		UserId: tr.UserId,
	}
}