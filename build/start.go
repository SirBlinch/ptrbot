// Запускаем бота
package build

import "fmt"

func Start() {
	token, myBot := Init()
	myBot.botUnit(token.token())
	fmt.Println("prt_token is " + token.token())
}
