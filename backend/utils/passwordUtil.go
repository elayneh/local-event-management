package utils

import "golang.org/x/crypto/bcrypt"

func ComparePasswords(hashedPassword string, password string) bool {
    hashedPasswordBytes := []byte(hashedPassword)
    err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(password))
    if err != nil {
        return false
    }
    return true
}

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}