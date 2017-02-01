package classfile

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

//简化版, 如果字符串中不包含null字符或补充字符,可以正常工作
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}