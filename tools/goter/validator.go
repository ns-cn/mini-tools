package goter

import "os"

func Required[T interface{}](data T, statement func(T) bool, message string, errorAction func()) {
	if !statement(data) {
		_, _ = os.Stderr.WriteString("ERROR! " + message + "\n")
		if errorAction != nil {
			errorAction()
		}
		os.Exit(1)
	}
}
