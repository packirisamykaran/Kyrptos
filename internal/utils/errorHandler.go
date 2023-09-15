package utils

import (
	"fmt"
	"runtime"
)

func HandleError(err error) error {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("[%s:%d] %v", file, line, err)
	return fmt.Errorf("[%s:%d] %w", file, line, err)
}
