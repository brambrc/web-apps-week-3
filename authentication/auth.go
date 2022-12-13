package authentication

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
	"github.com/joho/godotenv"
	"os"
)



func LoadEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
  
	return os.Getenv(key)

}


func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte(LoadEnv("secret_key"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}


	return tokenString, nil
}
