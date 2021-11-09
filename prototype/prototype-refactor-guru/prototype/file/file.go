package file

import (
	"fmt"
	"prototype_desing_pattern/prototype"
)

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{name: name}
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) Clone() prototype.Inode {
	return &File{name: f.name + "_clone"}
}
