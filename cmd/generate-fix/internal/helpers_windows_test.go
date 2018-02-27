package internal

import (
	"fmt"
	"testing"
)

func TestGetImportPathRoot(t *testing.T) {
	fmt.Println("this is import :", getImportPathRoot())
}
