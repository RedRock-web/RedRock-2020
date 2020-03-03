package jwts

import (
	"RedRock-2020/users"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//Jwt struct 表示一个 json web token
type Jwt struct {
	form      users.LoginForm
	Header    Header
	Payload   Payload
	Signature Signature
}

//NewJwt 返回一个 Jwt struct
func NewJwt() Jwt {
	return Jwt{}
}

//Create 返回 一个 特定的 Jwt
func (j Jwt) Create(form users.LoginForm, key string) (string, error) {
	j.form = form
	j.Signature.key = key

	hAndp, err := j.J2S()
	if err != nil {
		errors.New("Header 和 Payload 拼接失败！")
		return "", err
	}
	signature, err := j.Signature.New()
	if err != nil {
		errors.New("获取 Signature 失败！")
		return "", err
	}

	return hAndp + "." + signature, nil
}

func (j Jwt) Check(token string, key string) {
	flag := 0

	//首先把 token 和划分为 3 部分
	arr := strings.Split(token, ".")
	if len(arr) != 3 {
		flag = 1
		os.Exit(1)
	}

	//对 Header 解密
	_, err := base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		flag = 1
		os.Exit(1)
	}

	//对 payload 解密
	_, err = base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		flag = 1
		os.Exit(1)
	}

	//对 signature 解密
	sign, err := base64.StdEncoding.DecodeString(arr[2])
	if err != nil {
		flag = 1
		os.Exit(1)
	}

	fmt.Println(sign)
	hAndP := arr[0] + "." + arr[1]
	fmt.Println(hAndP)
	s := base64.StdEncoding.EncodeToString(HmacSha256(hAndP, "redrock"))
	fmt.Println(s)
	//if res := bytes.Compare(sign, s); res != 0 {
	//	flag = 1
	//}

	//if flag == 1 {
	//	return users.LoginForm{}, errors.New("token error!")
	//} else {
	//	var payload Payload
	//	//json.Unmarshal(pay, &payload)
	//	return users.LoginForm{payload.Username, payload.Password}, nil
	//}
}

//Header 表示 Jwt 的 header
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

//New 返回一个 Alg 为 HS256, Typ 为 JWT 的 Header 对象
func (h Header) New() Header {
	return Header{
		Alg: "HS256",
		Typ: "JWT",
	}
}

//Payload 表示 Jwt 的 payload
type Payload struct {
	Iss      string `json:"iss"`
	Exp      string `json:"exp"`
	Iat      string `json:"iat"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//New 返回一个特定的 Payload
func (p Payload) New(form users.LoginForm) Payload {
	return Payload{
		Iss:      "redrock",
		Exp:      strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),
		Iat:      strconv.FormatInt(time.Now().Unix(), 10),
		Username: form.Username,
		Password: form.Password,
	}
}

//Signature 表示 Jwt 的 signature
type Signature struct {
	key string
}

//New 返回 一个 signature
func (s Signature) New() (string, error) {
	j := NewJwt()
	str, err := j.J2S()
	if err != nil {
		errors.New("Header 和 Payload 拼接失败！")
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(HmacSha256(str, s.key))
	return signature, nil
}

//J2S() 将 Json 格式 的 Header 和 Payload 转换为 String 后并拼接
func (j Jwt) J2S() (string, error) {
	h, err := json.Marshal(j.Header.New())
	if err != nil {
		errors.New("解析 Header 出错！")
		return "", err
	}
	p, err := json.Marshal(j.Payload.New(j.form))
	if err != nil {
		errors.New("解析 Payload 出错！")
		return "", err
	}
	headerBase64 := base64.StdEncoding.EncodeToString(h)
	payloadBase64 := base64.StdEncoding.EncodeToString(p)

	return strings.Join([]string{headerBase64, payloadBase64}, "."), nil
}

func HmacSha256(str string, key string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(str))
	return mac.Sum(nil)
}
