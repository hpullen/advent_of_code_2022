package buffer

import "fmt"

type Buffer struct {
	chars      []rune
	charCounts map[rune]int
	size       int
}

func NewBuffer(size int) Buffer {
	var b Buffer
	b.charCounts = make(map[rune]int)
	b.size = size
	return b
}

func (b *Buffer) Len() int {
	return len(b.chars)
}

func (b *Buffer) Push(c rune) {
	if b.Len() == b.size {
		oldChar := b.chars[0]
		b.chars = b.chars[1:]
		b.charCounts[oldChar] -= 1
	}
	b.chars = append(b.chars, c)
	b.charCounts[c] += 1
}

func (b *Buffer) AllUnique() bool {
	for _, c := range b.chars {
		if b.charCounts[c] > 1 {
			return false
		}
	}
	return true
}

func (b *Buffer) Print() {
	fmt.Println(b.chars)
}
