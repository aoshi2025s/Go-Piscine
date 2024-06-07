package piscine

import "syscall"

func PrintRevAlpha() {
	fd := syscall.Stdout
	syscall.Write(fd, []byte("zyxwvutsrqponmlkjihgfedcba\n"))
}
