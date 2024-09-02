package object

type ObjType int

const (
	OBJ_BOUND_METHOD ObjType = iota
)

type Obj struct {
	Type ObjType 
	isMarked bool
	next *Obj
}

type ObjFunction struct {
	Obj Obj 
	arity int
}

