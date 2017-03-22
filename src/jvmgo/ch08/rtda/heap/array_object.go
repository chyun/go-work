package heap

func (self *Object) Bytes() []int8 {
	return self.data.([]int8)
}