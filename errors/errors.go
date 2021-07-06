package errors

import (
	"fmt"
)

func Error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s", line, where, message)
}
