package main

type Director struct {
	builder BuildProcess
}

func (d *Director) SetBuilder(b BuildProcess) {
	d.builder = b
}

func (d *Director) Construct() CellPhone {
	d.builder.SetCamera().SetDualSim().SetTorch().SetColorDisplay()
	return d.builder.GetCellPhone()
}
