// Метод получения токена бота телеги из параметра запуска

package build

import (
	"flag"
	"log"
)

type Token interface {
	token() string
}

type flagToken struct{}

func (t flagToken) token() string {
	var token string
	flag.StringVar(&token, "tgtoken", "", "tg token here")
	flag.Parse()
	if token == "" {
		log.Panic("No token!!")
	}
	return token
}
