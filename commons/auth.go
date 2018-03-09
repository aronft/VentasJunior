package commons

import (
	"crypto/rsa"
	"io/ioutil"
	"log"

	"github.com/blazte/VentasJunior/models"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	// PublicKey se usa para validar el token
	PublicKey *rsa.PublicKey
)

func init() {
	//ioutil es para leer archivos del sistema
	privateBytes, err := ioutil.ReadFile("./keys/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}
	publicBytes, err := ioutil.ReadFile("./keys/public.rsa")
	if err != nil {
		log.Fatal("No se puede leer el archivo público")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a privateKey")
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se pudo hacer el parse a PublicKey")
	}
}

// GenerateJWT genera el token para el cliente
func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			// ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer: "ventasjunior",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}
	return result
}
