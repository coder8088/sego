package ds

type Stack interface {
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
	Size() int
	IsEmpty() bool
}

type arrayStack struct {
	array []interface{}
	size  int
}

func NewStack(capacity int) *arrayStack {
	return &arrayStack{array: make([]interface{}, 0, capacity)}
}

func (as *arrayStack) Size() int {
	return as.size
}

func (as *arrayStack) IsEmpty() bool {
	return as.size == 0
}

func (as *arrayStack) Push(e interface{}) {
	as.array = append(as.array, e)
	as.size++
}

func (as *arrayStack) Pop() interface{} {
	if as.IsEmpty() {
		return nil
	}
	last := as.array[as.size-1]
	as.array = as.array[:as.size-1]
	as.size--
	return last
}

func (as *arrayStack) Peek() interface{} {
	if as.IsEmpty() {
		return nil
	}
	return as.array[as.size-1]
}
