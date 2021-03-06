package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/solrac97gr/twitter-fake/models"
)

/*GenerateJWT : Generate token with user info*/
func GenerateJWT(user models.User) (string, error) {
	myKey := []byte("thisisnotwitter")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastname":  user.Lastname,
		"birthdate": user.Birthdate,
		"biography": user.Biography,
		"location":  user.Location,
		"website":   user.Website,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
