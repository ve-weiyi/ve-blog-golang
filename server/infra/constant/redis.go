package constant

import "fmt"

const (
	ForgetPassword = "forget_password"
	Register       = "register"
)

func RedisWrapKey(module string, key interface{}) string {
	return fmt.Sprintf("%s:%v", module, key)
}
