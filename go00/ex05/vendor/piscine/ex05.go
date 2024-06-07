package piscine

import "syscall"

func PrintComb2() {
	fd := syscall.Stdout
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			syscall.Write(fd, []byte{byte(i/10+'0'), byte(i%10+'0'), ' ', byte(j/10+'0'), byte(j%10+'0')})
			if i != 98 {
				syscall.Write(fd, []byte(", "))
			}
		}
	}
	syscall.Write(fd, []byte("\n"))
}
