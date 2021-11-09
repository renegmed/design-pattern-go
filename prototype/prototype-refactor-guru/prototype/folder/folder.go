package folder

import (
	"fmt"
	"prototype_desing_pattern/prototype"
)

type Folder struct {
	children []prototype.Inode
	name     string
}

func NewFolder(children []prototype.Inode, name string) *Folder {
	return &Folder{children: children, name: name}
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() prototype.Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []prototype.Inode
	for _, i := range f.children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func (f *Folder) Add(node prototype.Inode) {
	f.children = append(f.children, node)
}
