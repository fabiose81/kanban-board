package aws

import (
	"fmt"
	"os"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"

	"github.com/joho/godotenv"

	"kanban-board/utils"
)

func ValidateJWT(tokenString string) (bool, string, error) {
	godotenv.Load()

	url := os.Getenv("AWS_COGNITO_ISSUER_URL")
	region := os.Getenv("AWS_REGION")
	userPoolId := os.Getenv("AWS_COGNITO_USERPOOL_ID")
	clientId := os.Getenv("AWS_COGNITO_CLIENT_ID")

	jwksURL := fmt.Sprintf(url+"/.well-known/jwks.json", region, userPoolId)

	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval: time.Hour,
	})
	if err != nil {
		return false, "", fmt.Errorf(utils.MSG_FAILED_JWKS, err)
	}

	token, err := jwt.Parse(tokenString, jwks.Keyfunc)
	if err != nil {
		return false, "", fmt.Errorf(utils.MSG_TOKEN_ERROR, err)
	}

	if !token.Valid {
		return false, "", fmt.Errorf(utils.MSG_INVALID_TOKEN)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, "", fmt.Errorf(utils.MSG_INVALID_TOKEN_CLAIMS)
	}

	expectedIss := fmt.Sprintf(url, region, userPoolId)
	if claims["iss"] != expectedIss {
		return false, "", fmt.Errorf(utils.MSG_INVALID_ISSUER)
	}

	if claims["client_id"] != clientId {
		return false, "", fmt.Errorf(utils.MSG_INVALID_AUDIENCE)
	}

	return true, claims["sub"].(string), nil
}
