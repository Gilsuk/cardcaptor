package lib

import "os"

/*
IsFileExist is
*/
func IsFileExist(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}
