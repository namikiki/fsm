package main

import (
	"io"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func maina() {
	r := gin.Default()

	r.Run(":8000")


	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("sessionid", store))

	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})

}

func TestName(t *testing.T) {
	//http.DefaultClient
	request, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/delete", nil)
	request.Header.Add("Cookie", "sessionId=MTY2NzAzNTk2MHxEdi1CQkFFQ180SUFBUkFCRUFBQVJQLUNBQUVHYzNSeWFXNW5EQWdBQm5WelpYSnBaQVp6ZEhKcGJtY01KZ0FrTm1OaU5USXhZVFl0WldJMFl5MDBNakUzTFRrM1pUSXRNRFUxTkdWbU5ESmtOR1U1fGxAoXwhkPXYe8-WhNA0kkuMs18CoQAHfoVycPdPIPNA; Path=/; Expires=Sun, 29 Oct 2023 09:38:27 GMT;")
	response, _ := http.DefaultClient.Do(request)
	all, _ := io.ReadAll(response.Body)
	log.Printf(string(all))

}

func TestA(t *testing.T) {
	filepath.WalkDir("/Users/zylzyl/go", func(path string, d fs.DirEntry, err error) error {
		log.Println(path)
		return err
	})
}
