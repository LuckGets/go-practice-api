package domain

type User struct {
	ID        int64  `json:"id"`
	Fullname  int64  `json:"fullname"`
	Lastname  int64  `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}
