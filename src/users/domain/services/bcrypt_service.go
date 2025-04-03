package services

type  IBcryptService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) (bool, error)
}