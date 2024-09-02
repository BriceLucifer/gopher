package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

func main(){
	fmt.Printf("GoLisp 👋 Version 0.1\n")
	fmt.Printf("Press Ctrl + C or Type exit to Exit \n")
	repl()
}

func repl() {
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

	fmt.Println("Simple REPL. Type 'exit' to quit.")

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
		result := strings.ToUpper(line)

		// Print the result
		fmt.Println("Output:", result)
	}
}