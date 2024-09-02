package debug

import (
	"fmt"
	"github.com/BriceLucifer/golox/chunk"	
)

func disassembleChunk(chunk *chunk.Chunk, name string) {
	fmt.Printf("== %s ==\n", name)
	offset := 0
	for offset < chunk.Count {
		offset = disassembleInstruction(chunk, offset)
	}
}

func disassembleInstruction(chunk *chunk.Chunk, offset int) int {
	return 1
}
