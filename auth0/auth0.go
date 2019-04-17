package auth0

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	input "github.com/krishna/rogerapp"
)

var clientId = "Ey30MV9vtSWAxNx1wPH98P1TbDEEs4ml"
var clientSecret = "hDG2ZnUmzIgTSssyQNre0ogLNiMpoux1ZI5j4-U8kEnGbegxSt1glUkYzwI41_G9"
var audience = "https://epnweb.auth0.com/api/v2/"
var grantType = "client_credentials"

type creds struct {
	Client       string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
}

type user struct {
	Connection string `json:"connection"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
}

type response struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func GetAuth0Token() string {

	url := "https://epnweb.auth0.com/oauth/token"

	payload := creds{
		Client:       clientId,
		ClientSecret: clientSecret,
		Audience:     audience,
		GrantType:    grantType,
	}
	payloadJSON, err := json.Marshal(payload)

	if err != nil {
		panic("error")
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	res1 := response{}
	jerr := json.Unmarshal(body, &res1)
	if jerr != nil {
		panic(err)
	}

	return res1.AccessToken

}

func CreateUser(userInput *input.SignupInfo) {
	url := "https://epnweb.auth0.com/api/v2/users"

	payload := user{
		Connection: "Username-Password-Authentication",
		Email:      *userInput.Email,
		Password:   *userInput.Password,
		Name:       *userInput.Name,
	}
	payloadJSON, err := json.Marshal(payload)

	if err != nil {
		panic("error")
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))

	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "bearer "+GetAuth0Token())

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func JwtVerification(ctx context.Context) error {
	var Authorization string
	tokenString, _ := ctx.Value(Authorization).(string)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(clientSecret), nil
	})
	if err != nil {
		return err
	}
	return nil
}
