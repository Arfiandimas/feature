package user_entities

import "time"

type Users struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `json:"name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Alamat     string    `json:"alamat,omitempty"`
	Usia       int8      `json:"usia,omitempty"`
	Pendidikan string    `json:"pendidikan,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
