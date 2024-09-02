package memory

import (
	"fmt"
	//"github.com/BriceLucfier/golox/common"
	"github.com/BriceLucifer/golox/object"
)

func reallocate(pointer interface{}, oldSize uint32) interface{}{
	fmt.Printf("hello")
	return 1
}

func markObject(object *object.Obj){
	fmt.Printf("Obj")
}

func collectGarbage(){
	fmt.Printf("collectGarbage")
}

func freeObject(){
	fmt.Printf("freeObject")
}
