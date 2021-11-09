package main

import "fmt"

func main() {
	diro := Director{}
	nokiaPhone := &Nokia{}
	diro.SetBuilder(nokiaPhone)
	phone := diro.Construct()
	printProduct(phone)

	samsungPhone := &Samsung{}
	diro.SetBuilder(samsungPhone)
	phone = diro.Construct()
	printProduct(phone)
}

func printProduct(phone CellPhone) {
	fmt.Println("Phone has camera? ", getAns(phone.Camera))
	fmt.Println("Phone has Dual Sim? ", getAns(phone.DualSim))
	fmt.Println("Phone has Torch? ", getAns(phone.Torch))
	fmt.Println("Phone has Color Display? ", getAns(phone.ColorDisplay))
}

func getAns(b bool) string {
	if b {
		return "YES"
	}
	return "NO"
}
