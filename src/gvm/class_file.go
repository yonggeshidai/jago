package gvm

import (
	"fmt"
)

type u1 uint8
type u2 uint16
type u4 uint32
type u8 uint64
/*
ClassFile {
	u4				magic;
	u2 				minor_version;
	u2 				major_version;
	u2 				constant_pool_count;
	cp_info 		constant_pool[constant_pool_count-1];
	u2 				access_flags;
	u2 				this_class;
	u2 				super_class;
	u2 				interfaces_count;
	u2 				interfaces[interfaces_count];
	u2 				fields_count;
	field_info 		fields[fields_count];
	u2 				methods_count;
	method_info 	methods[methods_count];
	u2 				attributes_count;
	attribute_info 	attributes[attributes_count];
}
*/
type ClassFile struct {
	size 	            int
	magic               u4
	minorVersion        u2
	majorVersion        u2
	constantPoolCount   u2
	constantPool        []ConstantPoolInfo
	accessFlags         u2
	thisClass           u2
	superClass          u2
	interfaces          []u2
	fieldsCount         u2
	fields              []FieldInfo
	methodsCount        u2
	methods             []MethodInfo
	attributes          []AttributeInfo
}

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type FieldInfo struct {
	accessFlags     u2
	nameIndex       u2
	descriptorIndex u2
	attributeCount  u2
	attributes      []AttributeInfo
}



/*
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type MethodInfo struct {
	accessFlags     u2
	nameIndex       u2
	descriptorIndex u2
	attributeCount  u2
	attributes      []AttributeInfo
}

func NewClassFile() *ClassFile {
	return &ClassFile{}
}

func (this *ClassFile) Print(){
	fmt.Printf("Size: %d bytes\n", this.size)
	fmt.Printf("magic: 0x%X\n", this.magic)
	fmt.Printf("minor version: %d\n", this.minorVersion)
	fmt.Printf("major version: %d\n", this.majorVersion)

	fmt.Printf("accessFlags: 0x%04x\n", this.accessFlags)
	fmt.Printf("thisClass: #%d\n", this.thisClass)
	fmt.Printf("superClass: #%d\n", this.superClass)
	fmt.Printf("interfaces: %d\n", len(this.interfaces))
	for i := 0; i < len(this.interfaces); i++  {
		fmt.Printf("\t#%d", this.interfaces[i])
	}

	fmt.Printf("fields: %d\n", len(this.fields))
	for i := 0; i < len(this.fields); i++  {
		fieldInfo := this.fields[i]
		fmt.Printf("\t%s", this.cpUtf8(fieldInfo.nameIndex))
	}

	fmt.Printf("method: %d\n", len(this.methods))
	for i := 0; i < len(this.methods); i++  {
		methodInfo := this.methods[i]
		fmt.Printf("\t%s\n", this.cpUtf8(methodInfo.nameIndex))
		for j :=0; j < len(methodInfo.attributes); j++ {
			attribute := methodInfo.attributes[j]
			this.printAttribute(attribute)
		}
	}

	fmt.Printf("attributes: %d\n", len(this.attributes))
	for j :=0; j < len(this.attributes); j++ {
		attribute := this.attributes[j]
		this.printAttribute(attribute)
	}
}

func (this *ClassFile) printAttribute(attribute AttributeInfo)  {
	switch attribute.(type) {
	case *CodeAttribute:
		codeAttribute := attribute.(*CodeAttribute)
		fmt.Printf("\t\tCode: %v\n", codeAttribute.code)
		fmt.Printf("\t\tMax locals: %d\n", codeAttribute.maxLocals)
		fmt.Printf("\t\tMax stack: %d\n", codeAttribute.maxStack)
	case *SourceFileAttribue:
		sourceFileAttribute := attribute.(*SourceFileAttribue)
		fmt.Printf("\t\tSourceFile: %v\n", this.cpUtf8(sourceFileAttribute.sourceFileIndex))
	case *LineNumberTableAttribute:
		lineNumberTableAttribute := attribute.(*LineNumberTableAttribute)
		fmt.Printf("\t\tlineNumberTableAttribute: %v\n", lineNumberTableAttribute)
	}
}


func (this *ClassFile) cpUtf8(index u2) string  {
	return u2s(this.constantPool[index].(*ConstantUtf8Info).bytes)
}

//func (this *ClassFile) cpClass(index u2) string  {
//	classInfo := this.constantPool[index].(*ConstantClassInfo)
//	return this.cpUtf8(classInfo.nameIndex)
//}
//
//func (this *ClassFile) cpNameAndType(index u2) (string, string)  {
//	nameAndTypeInfo := this.constantPool[index].(*ConstantNameAndTypeInfo)
//	return this.cpUtf8(nameAndTypeInfo.nameIndex), this.cpUtf8(nameAndTypeInfo.descriptorIndex)
//}
//
//// FieldRef, MethodRef, InterfaceMethodRef
//func (this *ClassFile) cpMemberRef(index u2) (string, string, string)  {
//	memberRefInfo := this.constantPool[index].(*ConstantFieldrefInfo)
//	name, descriptor := this.cpNameAndType(memberRefInfo.nameAndTypeIndex)
//	return this.cpClass(memberRefInfo.classIndex), name, descriptor
//}
//
//
//func (this *ClassFile) cpString(index u2) string  {
//	stringInfo := this.constantPool[index].(*ConstantStringInfo)
//	return this.cpUtf8(stringInfo.stringIndex)
//}
//
//func (this *ClassFile) cpInteger(index u2) int32  {
//	integerInfo := this.constantPool[index].(*ConstantIntegerInfo)
//	return int32(integerInfo.bytes)
//}
//
//func (this *ClassFile) cpLong(index u2) int64  {
//	longInfo := this.constantPool[index].(*ConstantLongInfo)
//	return int64((longInfo.highBytes << 32) | longInfo.lowBytes)
//}
//
//func (this *ClassFile) cpFloat(index u2) float32  {
//	floatInfo := this.constantPool[index].(*ConstantFloatInfo)
//	return float32(floatInfo.bytes)
//}
//
//func (this *ClassFile) cpDouble(index u2) float64  {
//	doubleInfo := this.constantPool[index].(*ConstantDoubleInfo)
//	return float64((doubleInfo.highBytes << 32) | doubleInfo.lowBytes)
//}


const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

/*
cp_info {
    u1 tag;
    u1 info[];
}
 */
type ConstantPoolInfo interface {

}

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
 */
type ConstantClassInfo struct {
	tag       u1
	nameIndex u2
}

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type ConstantFieldrefInfo struct {
	tag              u1
	classIndex       u2
	nameAndTypeIndex u2
}

/*
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type ConstantMethodrefInfo struct {
	tag              u1
	classIndex       u2
	nameAndTypeIndex u2
}

/*
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
 */
type ConstantInterfaceMethodrefInfo struct {
	tag              u1
	classIndex       u2
	nameAndTypeIndex u2
}

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
 */
type ConstantStringInfo struct {
	tag         u1
	stringIndex u2
}

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
 */
type ConstantIntegerInfo struct {
	tag   u1
	bytes u4
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
 */
type ConstantFloatInfo struct {
	tag   u1
	bytes u4
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
 */
type ConstantLongInfo struct {
	tag       u1
	highBytes u4
	lowBytes  u4
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
 */
type ConstantDoubleInfo struct {
	tag       u1
	highBytes u4
	lowBytes  u4
}

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
 */
type ConstantNameAndTypeInfo struct {
	tag             u1
	nameIndex       u2
	descriptorIndex u2
}

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
 */
type ConstantUtf8Info struct {
	tag     u1
	length  u2
	bytes   []u1 //u2 length
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
 */
type ConstantMethodHandleInfo struct {
	tag            u1
	referenceKind  u1
	referenceIndex u2
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
 */
type ConstantMethodTypeInfo struct {
	tag             u1
	descriptorIndex u2
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
 */
type ConstantInvokeDynamicInfo struct {
	tag                      u1
	bootstrapMethodAttrIndex u2
	nameAndTypeIndex         u2
}

