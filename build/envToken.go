package build

import (
	"fmt"
	"os"
)

type envToken struct{}

func (e envToken) token() string {
	value := os.Getenv("TG_KEY")
	fmt.Println("Env_Token is " + value)
	return value
}
