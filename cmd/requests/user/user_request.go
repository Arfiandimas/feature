package user_request

type UserStoreRequest struct {
	Name       string `json:"name" binding:"required,ascii"`
	Email      string `json:"email" binding:"required,email"`
	Alamat     string `json:"alamat" binding:"required,ascii"`
	Usia       int8   `json:"usia" binding:"required,numeric"`
	Pendidikan string `json:"pendidikan" binding:"required,ascii"`
}

type UserUpdateRequest struct {
	Name       string `json:"name" binding:"ascii"`
	Email      string `json:"email" binding:"email"`
	Alamat     string `json:"alamat" binding:"ascii"`
	Usia       int8   `json:"usia" binding:"numeric"`
	Pendidikan string `json:"pendidikan" binding:"ascii"`
}
