package piscine

import "syscall"

func PrintAlp() {
	fd := syscall.Stdout
	syscall.Write(fd, []byte("abcdefghijklmnopqrstuvwxyg\n"))
}
