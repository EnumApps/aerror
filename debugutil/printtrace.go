package debugutil

import (
	"fmt"

	"github.com/EnumApps/aerror"
)

func PrintTrace(err error) {
	fmt.Println(err)
	switch a := err.(type) {
	case *aerror.AError:
		for i, s := range a.Trace {
			fmt.Println(" - Call Stack:", i, s.File, s.Line)
		}
	}
}
