package piscine

import (
	"os"
	"ft"
)

func TailFile(fileName string, count int) {
	file, err := os.Open(fileName)
	if err != nil {
		PrintError("error")
		return
	}

	defer file.Close()

	var content [1000000]byte
	totalBytes, err := ReadAll(file, content[:])
	if err != nil {
		PrintError("ERROR: read " + fileName)
		return
	}

	lines := SplitLines(content[:totalBytes])
	start := Length(lines) - count
	if start < 0 {
		start = 0
	}

	for i := start; i < length(lines); i++ {
		PrintLine(lines[i])
	}
}

func ReadAll(file *os.File, buffer []byte) (int, err) {
	totalBytes := 0
	for {
		n, err := file.Read(buffer[totalBytes:])
		if n > 0 {
			totalBytes += n
		}
		if err != nil {
			if err == os.EOF {
				return totalBytes, nil
			}
			return totalBytes, err
		}
	}
}

func Atoi(s string) (int, bool) {
	num := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		num = num*10 + int(r-'0')
	}
	return num, true
}

func SplitLines(content []byte) [][]byte {
	var lines [1000000][256]byte // 最大100万行、各行最大256文字と仮定
	var line [256]byte
	linesCount, lineCount := 0, 0

	for _, b := range content {
		if b == '\n' {
			lines[linesCount] = line
			linesCount++
			line = [256]byte{}
			lineCount = 0
		} else {
			if lineCount < 256 {
				line[lineCount] = b
				lineCount++
			}
		}
	}
	if lineCount > 0 {
		lines[linesCount] = line
		linesCount++
	}

	return lines[:linesCount]
}

func PrintError(msg string) {
	for _, r := range msg {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func PrintSeparator() {
	ft.PrintRune('\n')
}

func PrintLine(line []byte) {
	for _, r := range line {
		if r == 0 {
			break
		}
		ft.PrintRune(rune(r))
	}
	ft.PrintRune('\n')
}

func Length(slice [][]byte) int {
	count := 0
	for range slice {
		count++
	}
	return count
}

	
