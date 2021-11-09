package main

type BuildProcess interface {
	SetCamera() BuildProcess
	SetDualSim() BuildProcess
	SetTorch() BuildProcess
	SetColorDisplay() BuildProcess
	GetCellPhone() CellPhone
}
