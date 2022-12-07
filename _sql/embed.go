package sql

import "embed"

//go:embed mssql mysql oracle postgresql sqlite3
var FS embed.FS
