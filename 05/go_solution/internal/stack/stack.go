package stack

type Stack []string

func (s *Stack) Push(item string) {
	*s = append(*s, item)
}

func (s *Stack) PushMany(items []string) {
	*s = append(*s, items...)
}

func (s *Stack) Pop() string {
	idx := len(*s) - 1
	item := (*s)[idx]
	*s = (*s)[:idx]
	return item
}

func (s *Stack) PopMany(n int) []string {
	idx := len(*s)
	items := (*s)[idx-n : idx]
	*s = (*s)[:idx-n]
	return items
}
