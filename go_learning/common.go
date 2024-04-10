package common

import (
	"fmt"
	"strings"
)

const FLAG = "#####"

func PrintLine(str string) {
	var build strings.Builder
	var build2 strings.Builder
	build.WriteString("\n")
	for i := 0; i < 10; i++ {
		build.WriteString(FLAG)
		build2.WriteString(FLAG)
	}
	build.WriteString("\n\t\t\t\t")
	build.WriteString(str)
	build.WriteString("\n")
	build.WriteString(build2.String())
	build.WriteString("\n")
	fmt.Println(build.String())
}
