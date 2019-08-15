package instructions

import (
	"fmt"
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/instructions/comparisons"
	"jvmgo/ch07/instructions/constants"
	"jvmgo/ch07/instructions/control"
	"jvmgo/ch07/instructions/loads"
	"jvmgo/ch07/instructions/math"
	"jvmgo/ch07/instructions/references"
	"jvmgo/ch07/instructions/stack"
	"jvmgo/ch07/instructions/stores"

	"github.com/axgle/mahonia"
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return &constants.NOP{}
	case 0x01:
		return &constants.ACONST_NULL{}
	case 0x03:
		return &constants.ICONST_0{}
	case 0x04:
		return &constants.ICONST_1{}
	case 0x2c:
		return &loads.ALOAD_2{}
	case 0x2d:
		return &loads.ALOAD_3{}
	case 0x3c:
		return &stores.ISTORE_1{}
	case 0x3d:
		return &stores.ISTORE_2{}
	case 0x12:
		return &constants.LDC{}
	case 0x1b:
		return &loads.ILOAD_1{}
	case 0x1c:
		return &loads.ILOAD_2{}
	case 0x10:
		return &constants.BIPUSH{}
	case 0x4d:
		return &stores.ASTORE_2{}
	case 0x4e:
		return &stores.ASTORE_3{}
	case 0x59:
		return &stack.DUP{}
	case 0x60:
		return &math.IADD{}
	case 0x84:
		return &math.IINC{}
	case 0x99:
		return &comparisons.IFEQ{}
	case 0xa2:
		return &comparisons.IF_ICMPGE{}
	case 0xa7:
		return &control.GOTO{}
	case 0xb2:
		return &references.GET_STATIC{}
	case 0xb3:
		return &references.PUT_STATIC{}
	case 0xb4:
		return &references.GET_FIELD{}
	case 0xb5:
		return &references.PUT_FIELD{}
	case 0xb6:
		return &references.INVOKE_VIRTUAL{}
	case 0xb7:
		return &references.INVOKE_SPECIAL{}
	case 0xbb:
		return &references.NEW{}
	case 0xc0:
		return &references.CHECK_CAST{}
	case 0xc1:
		return &references.INSTANCE_OF{}
	default:
		str := "不支持操作码: 0x%x!"
		str = ConvertToString(str, "gbk", "utf-8")
		panic(fmt.Errorf(str, opcode))
	}
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
