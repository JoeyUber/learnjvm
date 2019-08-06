package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {

}

func (self *Stack) push(frame *Frame) {}
func (self *Stack) pop() *Frame       {}
func (self *Stack) top() *Frame {
	return self._top
}
