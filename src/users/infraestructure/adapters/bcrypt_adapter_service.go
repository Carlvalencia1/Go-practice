package adapters

import "golang.org/x/crypto/bcrypt"

type BcryptAdapter struct {
	cost int
}

func NewBcryptAdapter() *BcryptAdapter {
	return &BcryptAdapter{cost: 12}
}

func (ba *BcryptAdapter) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), ba.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (ba *BcryptAdapter) ComparePassword(hashedPassword string, providePassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providePassword))
	if err != nil {
		return false, err
	}
	return true, nil
}
