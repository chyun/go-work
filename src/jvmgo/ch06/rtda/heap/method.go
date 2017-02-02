package heap

import "jvmgo/ch06/classfile"

type Method struct {
	ClassMember
	maxStack    uint
	maxLocals   uint
	code        []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberIInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		method[i] = &Method{}
		method[i].class = class
		method[i].copyMemberInfo(cfMethod)
		method[i].copyAttributes(cfMethod)
	}
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberIInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack  = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code      = codeAttr.Code()
	}
}