package members

import (
	"fmt"
	"github.com/spf13/viper"
)

func Testfunc() {
	fmt.Println(viper.Get("MemberZugang"))
}
