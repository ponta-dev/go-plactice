package test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before test")
	os.Exit(m.Run())
}

func TestEchoHello(t *testing.T) {
	EchoHello()
}
