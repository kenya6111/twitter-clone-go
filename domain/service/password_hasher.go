package service

// パスワードハッシュ化のインターフェース
type PasswordHasher interface {
	CompareHashAndPassword(hashedPassword, password string) error
	HashPassword(password string) (string, error)
	GenerateSecureToken(n int) (string, error)
}
