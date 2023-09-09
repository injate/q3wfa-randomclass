// +build windows

package main

import (
	"os"
	"syscall"
	"time"
	"fmt"
)

func getLastAccessTime(filename string) (time.Time, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return time.Time{}, err
	}

	stat, ok := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return time.Time{}, fmt.Errorf("Not a syscall.Win32FileAttributeData structure")
	}
	return time.Unix(0, stat.LastAccessTime.Nanoseconds()), nil
}
