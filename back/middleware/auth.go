package middleware

import (
	"apigo/back/models"
	u "apigo/back/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JwtAuthentication : jwt middleware
var JwtAuthentication = func(c *gin.Context) {

	w := c.Writer
	r := c.Request

	notAuth := []string{"/users", "/login"} //List of endpoints that doesn't require auth
	requestPath := r.URL.Path               //current request path

	//check if request does not need authentication, serve the request if it doesn't need it
	for _, value := range notAuth {

		if value == requestPath {
			c.Next()
			return
		}
	}

	response := make(map[string]interface{})
	tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

	if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
		response = u.Message(false, "Missing auth token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}

	splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
	if len(splitted) != 2 {
		response = u.Message(false, "Invalid/Malformed auth token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}

	tokenPart := splitted[1] //Grab the token part, what we are truly interested in
	tk := &models.Token{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})

	if err != nil { //Malformed token, returns with http code 403 as usual
		response = u.Message(false, "Malformed authentication token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}

	if !token.Valid { //Token is invalid, maybe not signed on this server
		response = u.Message(false, "Token is not valid.")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		return
	}

	//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
	fmt.Sprintf("User %", tk.Uuid) //Useful for monitoring
	var tokenMap map[string]string
	tokenMap = make(map[string]string)
	tokenMap["uuid"] = (tk.Uuid).String()
	tokenMap["access_level"] = strconv.Itoa(tk.AccessLevel)
	ctx := context.WithValue(r.Context(), "user", tokenMap)
	r = r.WithContext(ctx)
	c.Request = r
	c.Next() //proceed in the middleware chain!

}
