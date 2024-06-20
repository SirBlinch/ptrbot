// Запускаем бота
package build

import "fmt"

func Start() {
	token, myBot := Init()
	fmt.Println("prt_token is " + token.token())
	myBot.botUnit(token.token())

}
