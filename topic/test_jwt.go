package main

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var MySecret = []byte("secret")

func GenToken(username string, password string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, // 自定义字段
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // 过期时间
			Issuer:    "laoguo",                             // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	// 注意这个地方一定要是字节切片不能是字符串
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{},
		func(token *jwt.Token) (i interface{}, err error) {
			return MySecret, nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func main() {
	/*  	s, err := GenToken("ghz", "123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("s: %v\n", s)  */

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImdoeiIsInBhc3N3b3JkIjoiMTIzIiwiZXhwIjoxNzA0MDUwMDU0LCJpc3MiOiJsYW9ndW8ifQ.OYJbx6cwNPU0XhJ8qV5hl_c1LB5-bhZInXIB-9_o1mY"

	mc, err := ParseToken(token)
	if err != nil {
		panic(err)
	}

	fmt.Printf("mc.Password: %v\n", mc.Password)
	fmt.Printf("mc.Username: %v\n", mc.Username)

}
