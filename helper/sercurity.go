package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error){
	hasehdPass, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "",err
	}
	return string(hasehdPass),nil
}

func ComparePassword(hashedPassword, plainPassword string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)) == nil
}