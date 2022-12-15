package dirs

import (
	"fmt"
	"strconv"
	"strings"
)

type Dir struct {
	name  string
	dirs  map[string]*Dir
	files []file
	depth int
	Size  int
}

type file struct {
	name string
	size int
}

func newDir(name string, depth int) *Dir {
	dir := Dir{name: name, depth: depth}
	dir.dirs = make(map[string]*Dir)
	return &dir
}

func FillDirs(console []string) (*Dir, []*Dir) {
	topDir := newDir(strings.Split(console[0], " ")[2], 0)
	var allDirs []*Dir
	var dirStack []*Dir
	currentDir := topDir

	for _, cmd := range console[1:] {
		splitCmd := strings.Split(cmd, " ")
		if splitCmd[0] == "$" {
			switch splitCmd[1] {
			case "ls":
				continue
			case "cd":
				target := splitCmd[2]
				if target == ".." {
					allDirs = append(allDirs, currentDir)
					prevDir := dirStack[len(dirStack)-1]
					prevDir.Size += currentDir.Size
					currentDir = prevDir
					dirStack = dirStack[:len(dirStack)-1]
				} else {
					dirStack = append(dirStack, currentDir)
					currentDir = currentDir.dirs[target]
				}
			}
		} else {
			if splitCmd[0] == "dir" {
				currentDir.dirs[splitCmd[1]] = newDir(splitCmd[1], len(dirStack)+1)
			} else {
				size, _ := strconv.Atoi(splitCmd[0])
				file := file{name: splitCmd[1], size: size}
				currentDir.files = append(currentDir.files, file)
				currentDir.Size += file.size
			}
		}
	}

	// Finish going back up the stack to compute sizes
	nInStack := len(dirStack)
	for i := 0; i < nInStack; i++ {
		prevDir := dirStack[len(dirStack)-1]
		prevDir.Size += currentDir.Size
		allDirs = append(allDirs, currentDir)
		currentDir = prevDir
		dirStack = dirStack[:len(dirStack)-1]
	}

	return topDir, allDirs
}

func (d *Dir) Print() {
	indent := strings.Repeat("  ", d.depth)
	fmt.Printf("%s- %s (dir)\n", indent, d.name)
	for _, subdir := range d.dirs {
		subdir.Print()
	}
	for _, file := range d.files {
		fmt.Printf("%s  - %s (file, size=%d)\n", indent, file.name, file.size)
	}
}
