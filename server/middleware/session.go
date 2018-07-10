package middleware

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	Secret string         = "default-token"
	SessID string         = "default-sessions"
	Store  sessions.Store = nil
)

func init() {
	rand.Seed(time.Now().UnixNano())
	Secret = randStringRunes(10)
	SessID = randStringRunes(10)
	log.Printf("Secret : %v\nSessID : %v\n", Secret, SessID)
	Store = NewStore(Secret)
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Configure(token, name string) {
	SessID = name
	Secret = token
}

func NewStore(secret string) sessions.Store {
	return cookie.NewStore([]byte(secret))
}

func NewHandler(name string, store sessions.Store) gin.HandlerFunc {
	return sessions.Sessions(name, store)
}

func DefaultHandler() gin.HandlerFunc {
	return NewHandler(SessID, Store)
}

func Session() gin.HandlerFunc {
	return DefaultHandler()
}
