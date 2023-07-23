package user_entities

import "time"

type Users struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Alamat     string    `json:"alamat"`
	Usia       int8      `json:"usia"`
	Pendidikan string    `json:"pendidikan"`
	CreatedAt  time.Time `json:"created_at"`
}
