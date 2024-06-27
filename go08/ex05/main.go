package main

import "ft"

type Door struct {
	state int
}

const (
	OPEN  = 1
	CLOSE = 0
)

func main() {
	var door Door

	OpenDoor(&door)
	if IsDoorClose(&door) {
		PrintStr("the Door is closed")
		OpenDoor(&door)
	}
	if IsDoorOpen(&door) {
		PrintStr("the Door is opened")
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func OpenDoor(door *Door) {
	door.state = OPEN
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(door *Door) bool {
	PrintStr("is the Door opened ?")
	return door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

