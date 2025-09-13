package domains

import "fmt"

// Declarar Struct
type StackArray struct {
	Items []int
}

func (s *StackArray) Push(data int) {
	s.Items = append(s.Items, data)
}

func (s *StackArray) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *StackArray) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	removed := s.Items[len(s.Items)]
	s.Items = s.Items[:len(s.Items)-1]
	return removed
}

func (s *StackArray) Show() {
	for _, item := range s.Items {
		fmt.Print(item, "")
	}
	fmt.Println()
}

func (s *StackArray) Size() int {
	return len(s.Items)
}
