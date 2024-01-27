package yaegi

import (
	"fmt"
	"testing"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func Test(t *testing.T) {
	intp := interp.New(interp.Options{}) // 初始化一个 yaegi 解释器
	intp.Use(stdlib)                     // 允许脚本调用（几乎）所有的 Go 官方 package 代码

	intp.Eval(src) // src 就是上面的 Go 代码字符串
	v, _ := intp.Eval("plugin.Fib")
	fu := v.Interface().(func(int) int)

	fmt.Println("Fib(35) =", fu(35))
}
