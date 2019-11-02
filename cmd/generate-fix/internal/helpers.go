package internal

// getImportPathRoot returns the root path to use in import statements.
func getImportPathRoot() string {
	return *pkgRoot
}
