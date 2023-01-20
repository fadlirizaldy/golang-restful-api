package middleware

import (
	"time"

	"project_alterra/constants"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId int, name string) (string, error) {
	// Untuk menyimpan payload yang dikirim
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	// exp untuk ngasih batas waktu tokennya
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// ini utk ngegenerate token menggunakan method signingmethodHS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// jadi utk bagian ketiga jwt ditambahin signedString dari Secret JWT tadi (tentuin sendiri)
	return token.SignedString([]byte(constants.SECRET_JWT))
}