package piscine

import "syscall"

func IsNegative(nb int) {
	fd := syscall.Stdout
	if nb < 0 {
		syscall.Write(fd, []byte("T\n"))
	} else {
		syscall.Write(fd, []byte("F\n"))
	}
}
