package gotestlib2

import (
	"fmt"

	"github.com/liranms/gotestlib"
)

func CallMe() {
	fmt.Printf("callme = %d", gotestlib.ExportedFunc())
}
