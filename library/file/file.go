package file

import "os"

func FileExists(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	}
	return false
}
