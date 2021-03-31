//ch16/exinit/exinit.go
package exinit

import "fmt"

var (
	a = c + b // ❶ a값은 c와 b가 초기화된 다음 초기화됩니다.
	b = f()   // ❷ b값은 4가 됩니다.
	c = f()   // ❸ c값은 5 입니다.
	d = 3     // ❹ d값은 초기화가 끝난뒤 6이 됩니다.
)

func init() { // ➎
	d++                             // ➏
	fmt.Println("init function", d) // ➐
}

func f() int { // ➑
	d++                      // ➒
	fmt.Println("f() d:", d) // ❿
	return d                 // ⓫
}

func PrintD() {
	fmt.Println("d:", d)
}
