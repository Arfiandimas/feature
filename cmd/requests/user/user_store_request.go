package user_request

type UserStoreRequest struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Alamat     string `json:"alamat" binding:"required"`
	Usia       int8   `json:"usia" binding:"required"`
	Pendidikan string `json:"pendidikan" binding:"required"`
}
