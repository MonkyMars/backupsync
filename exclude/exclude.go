// Package exclude contains an array and a function of directories that need to be excluded from copying.
package exclude

import (
	"io/fs"
	"slices"
)

var Excluded = []string{".venv", ".git", ".vercel", ".next", ".vscode", "node_modules"}

func Exclude(d fs.DirEntry) bool {
	name := d.Name()
	return slices.Contains(Excluded, name)
}
