package piscine

import "syscall"

func PrintComb() {
	fd := syscall.Stdout
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				syscall.Write(fd, []byte{byte(i), byte(j), byte(k)})
				if i != '7' {
					syscall.Write(fd, []byte(", "))
				}
			}
		}
	}
	syscall.Write(fd, []byte("\n"))
}
