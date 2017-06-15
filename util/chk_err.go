package util

import (
	"fmt"
	"os"
)

func ChkErr(err error, place string) {
	if err != nil {
		fmt.Printf("(%s)", place)
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(0)
	}
}
