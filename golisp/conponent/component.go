package conponent

type I uint
type L float64
type T *string
func A (ceil []byte) string {
	return string(ceil)
}
const N uint = 1024
var hp,sp,ATOM,PRIM,CONS,CLOS,NIL I = 0,I(N),32,48,64,80,96
var cell [N]byte 
var nil,tru,err,env L

func box(t I,i I) L {
	return 1.0
}

func ord(x L) I ;

func num(n L) L ;
func equ(x L, y L) I;
func atom(s []byte);

func cons(x L, y L) L ;

