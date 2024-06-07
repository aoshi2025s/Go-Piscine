package piscine

import "syscall"

func PrintDigit() {
	fd := syscall.Stdout
	syscall.Write(fd, []byte("0123456789\n"))
}
