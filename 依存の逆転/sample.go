package main

type PC struct {
	PCParts PCPartsInterface
}

type PCParts struct {
	Name string
}

func (p PCParts) Exec() {
	println("this parts is: " + p.Name)
}

type PCPartsV2 struct {
	Name string
}

func (p PCPartsV2) Exec() {
	println("this parts is: " + p.Name)
}

type PCPartsV3 struct {
	Name string
}

func (p PCPartsV3) ExecV() {
	println("this parts is: " + p.Name)
}

type PCPartsInterface interface {
	Exec()
}

func NewPC(pcParts PCPartsInterface) *PC {
	return &PC{
		PCParts: pcParts,
	}
}

func main() {
	// PC は PCParts に依存していない

	pcParts := &PCParts{Name: "CPU"}
	pc := NewPC(pcParts)
	pc.PCParts.Exec()

	pcPartsV2 := &PCPartsV2{Name: "CPUV2"}
	pc = NewPC(pcPartsV2)
	pc.PCParts.Exec()

	// cannot use pcPartsV3 (variable of type *PCPartsV3) as PCPartsInterface value in argument to NewPC: missing method Exec
	// PCPartsV3 は Exec メソッドを持っていない
	// pcPartsV3 := &PCPartsV3{Name: "CPUV3"}
	// pc = NewPC(pcPartsV3)
	// pc.PCParts.Exec()
}
