package test

import (
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestUUid(t *testing.T) {
	v := uuid.NewV4()
	fmt.Println(v)
}
