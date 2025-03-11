package dataStructure

type Stack struct {
	Data []int
	Size int
}

func StackConstructor() Stack {
	return Stack{
		Data: []int{},
		Size: 0,
	}
}

func (s *Stack) Push(x int) {
	s.Data = append(s.Data, x)
	s.Size++
}

func (s *Stack) Pop() int {
	if s.Size == 0 {
		panic("stack is empty")
	}
	topElement := s.Data[s.Size-1]
	s.Data = s.Data[:s.Size-1]
	s.Size--
	return topElement
}
