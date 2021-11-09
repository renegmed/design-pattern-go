package main

import (
	"fmt"
	"prototype_desing_pattern/prototype"
	"prototype_desing_pattern/prototype/file"
	"prototype_desing_pattern/prototype/folder"
)

func main() {
	file1 := file.NewFile("File1") // &file.File{name: "File1"}
	file2 := file.NewFile("File2") //&file.File{name: "File2"}
	file3 := file.NewFile("File3") //&file.File{name: "File3"}

	folder1 := folder.NewFolder([]prototype.Inode{file1, file3}, "Folder1")

	// folder1 := &folder.Folder{
	// 	children: []prototype.Inode{file1},
	// 	name:     "Folder1",
	// }

	folder2 := folder.NewFolder([]prototype.Inode{folder1, file2, file3}, "Folder2")

	// folder2 := &folder.Folder{
	// 	children: []prototype.Inode{folder1, file2, file3},
	// 	name:     "Folder2",
	// }

	fmt.Println("1. +++++++++++++++++++++++")

	fmt.Println("Printing hierarchy for Folder2")
	folder2.Print("  ")

	fmt.Println("2. +++++++++++++++++++++++")

	cloneFolder := folder2.Clone()
	fmt.Println("Printing hierarchy for clone Folder2")
	cloneFolder.Print("  ")

	fmt.Println("3. +++++++++++++++++++++++")
	folder2.Add(file2)
	folder2.Print("  ")
	fmt.Println("4. +++++++++++++++++++++++")

	cloneFolder = folder2.Clone()
	fmt.Println("Printing hierarchy for clone Folder2")
	cloneFolder.Print("  ")
}
