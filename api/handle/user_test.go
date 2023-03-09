package handle

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"testing"

	"fsm/api/req"
)

func TestLogin(t *testing.T) {
	login := req.UserLogin{
		Email:    "232021312@qq.com",
		PassWord: "aaaaaaaaaaaa",
	}

	marshal, err := json.Marshal(login)
	if err != nil {
		log.Printf("err %v", err)
	}
	request, err := http.NewRequest("POST", "http://127.0.0.1:8080/login", bytes.NewBuffer(marshal))
	request.Header.Add("content-type", "application/json")
	if err != nil {
		log.Printf("err %v", err)
	}

	respone, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("err %v", err)
	}

	logcookies := respone.Cookies()

	delResp, err := http.NewRequest("GET", "http://127.0.0.1:8080/delete", nil)
	delResp.AddCookie(logcookies[0])

	derespone, err := http.DefaultClient.Do(delResp)
	if err != nil {
		log.Printf("err %v", err)
	}

	all, err := io.ReadAll(derespone.Body)
	log.Println(string(all))

}
