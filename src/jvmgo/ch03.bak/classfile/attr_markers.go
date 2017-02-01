package classfile

type DeprecatedAttribute struct { MarkerAttribute }
type SyntheticAttribute struct { MarkerAttribute }

type MarkerAttribute struct {}
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	//do nothing, 这两个属性都没有数据
}