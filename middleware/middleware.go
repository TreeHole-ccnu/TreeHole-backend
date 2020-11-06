package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type jwtClaims struct {
	jwt.StandardClaims
	Phone string `json:"phone"`
}

var (
	key        = "miniProject" //salt
	ExpireTime = 3600          //token expire time
)

func JwtAAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("token")
		if tokenStr == "" {
			c.String(401, "token invalid")
			c.Abort()
			//跳转登录界面
			return
		}
		token, err := verifyToken(tokenStr)
		if token == nil || err != nil {
			c.String(401, "token invalid")
			c.Abort()
			//跳转登录页面
			return
		}
		if !token.Valid {
			c.String(401, "token invalid")
			c.Abort()
			//跳转登录页面
			return
		}
		claim := token.Claims
		c.Set("uid", claim.(jwt.MapClaims)["uid"])
		c.Next()
	}
}

func ProduceToken(phone string) string {
	//id, _ := strconv.Atoi(uid)
	claims := &jwtClaims{
		Phone: phone,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	//fmt.Println(singedToken, err)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims jwtClaims) (string, error) {
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func verifyToken(verifyToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(verifyToken, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(key), nil
	})
	if err != nil {
		log.Print("verifyToken err:")
		fmt.Println(err)
		return nil, err
	}
	return token, nil
}