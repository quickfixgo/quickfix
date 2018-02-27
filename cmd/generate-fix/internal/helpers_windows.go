// +build windows

package internal

import (
	"os"
	"path/filepath"
	"strings"
)

// getImportPathRoot returns the root path to use in import statements.
// The root path is determined by stripping "$GOPATH/src/" from the current
// working directory.  For example, when generating code within the QuickFIX/Go
// source tree, the returned root path will be "github.com/quickfixgo/quickfix".
func getImportPathRoot() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	//goSrcPath := `C:\Users\Administrator\go\src\`
	goSrcPath := filepath.Join(os.Getenv("USERPROFILE"), "Go", "src")
	importPathRoot := filepath.ToSlash(strings.Replace(pwd, goSrcPath, "", 1))
	return strings.TrimLeft(importPathRoot, "/")
}
