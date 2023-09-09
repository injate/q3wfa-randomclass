// +build linux darwin

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

	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return time.Time{}, fmt.Errorf("Not a syscall.Stat_t structure")
	}
	return time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec)), nil
}
