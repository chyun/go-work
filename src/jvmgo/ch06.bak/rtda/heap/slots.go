package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot
