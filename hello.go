package rpc

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("welcome, %v", name)
}
