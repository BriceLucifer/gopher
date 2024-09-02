package chunk
import(
  "fmt"
  "github.com/BriceLucifer/golox/value"
)

type Opcode int

const (
  OP_CONSTANT Opcode = iota
  OP_NIL  
  OP_TRUE
  OP_FALSE
  OP_POP
  OP_GET_LOCAL
  OP_SET_LOCAL
  OP_GET_GLOBAL
  OP_DEFINE_GLOBAL
  OP_SET_GLOBAL
  OP_GET_UPVALUE
  OP_SET_UPVALUE
  OP_GET_PROPERTY
  OP_SET_PROPERTY
  OP_GET_SUPER
  OP_EQUAL
  OP_GREATER
  OP_LESS
  OP_ADD
  OP_SUBTRACT
  OP_MULTIPLY
  OP_DIVIDE
  OP_NOT
  OP_NEGATE
  OP_PRINT
  OP_JUMP
  OP_JUMP_IF_FALSE
  OP_LOOP
  OP_CALL
  OP_INVOKE
  OP_SUPER_INVOKE
  OP_CLOSURE
  OP_CLOSE_UPVALUE
  OP_RETURN
  OP_CLASS
  OP_INHERIT
  OP_METHOD
)

// 考虑后期加入到value.go

type Chunk struct {
	Count     int         // 计数
	capacity  int         // 容量
	code      *uint8      // 代码块
	lines     *int        // 行数
	constants *value.ValueArray // 常数
}

func initChunk(chunk *Chunk){
	println()
  fmt.Printf("initChunk")
}
func freeChunk(chunk *Chunk){
	println()
}
func writeChunk(chunk *Chunk, Byte uint8, line int){
	println()
}
func addConstant(chunk *Chunk, value value.Value){
	println()
}

