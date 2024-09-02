package repl

import (
	"fmt"
	"os"
	"strings"
	"github.com/BriceLucifer/golisp/parse"

	"github.com/chzyer/readline"
)

func Repl() {
	// 创建 readline 配置
	config := readline.Config{
		Prompt:          "GoLisp 🩵 >> ",
		HistoryFile:     "./readline.tmp", 	   // 历史记录文件路径
		InterruptPrompt: "^C",                // 中断提示符
		EOFPrompt:       "exit",              // EOF 提示符

		// 历史搜索模式
		HistorySearchFold: true, // 忽略大小写
	}

	// 创建 readline 实例
	rl, err := readline.NewEx(&config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Readline error: %v\n", err)
		return
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil { // 检测是否为 EOF (Ctrl+D)
			break
		}

		if strings.ToLower(line) == "exit" {
			fmt.Println("Exiting REPL.")
			break
		}

		// Evaluate the input by converting it to uppercase
		expr, err := EvalInput(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}

		fmt.Printf("Parsed expression: %#v\n", expr)
	}
}


func EvalInput(input string) (parse.Expr, error) {
	parser := parse.NewParser(input)
	expr := parser.Parse()
	return expr, nil // TODO: 添加错误处理
}