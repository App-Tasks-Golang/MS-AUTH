package domain


// User representa un usuario en el sistema
type Auth struct {
	ID       int    `json:"user_id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}


type TokenBlacklist struct {
	ID    uint   `gorm:"primaryKey"`
	Token string `gorm:"unique;not null"`
}
