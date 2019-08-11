package instructions

import (
	"fmt"
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/instructions/comparisons"
	"jvmgo/ch05/instructions/constants"
	"jvmgo/ch05/instructions/control"
	"jvmgo/ch05/instructions/loads"
	"jvmgo/ch05/instructions/math"
	"jvmgo/ch05/instructions/stores"

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
	case 0x3c:
		return &stores.ISTORE_1{}
	case 0x3d:
		return &stores.ISTORE_2{}
	case 0x1b:
		return &loads.ILOAD_1{}
	case 0x1c:
		return &loads.ILOAD_2{}
	case 0x10:
		return &constants.BIPUSH{}
	case 0x60:
		return &math.IADD{}
	case 0x84:
		return &math.IINC{}
	case 0xa2:
		return &comparisons.IF_ICMPGE{}
	case 0xa7:
		return &control.GOTO{}
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
