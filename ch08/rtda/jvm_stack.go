package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *Frame) {
	if self.size == self.maxSize {
		panic(" this stack is maxSize")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++

}
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("this stack is empty")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top

}
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("this stack is empty")
	}
	return self._top
}

func (self *Stack) isEmpty() bool {
	return self._top == nil
}
