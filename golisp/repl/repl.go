package repl

import (
	"fmt"
	"os"
	"strings"

	"github.com/chzyer/readline"
	"github.com/BriceLucifer/golisp/parse"
)

func Repl() {
	config := readline.Config{
		Prompt:          "GoLisp 🩵 >> ",
		HistoryFile:     "./readline.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold: true,
	}

	rl, err := readline.NewEx(&config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Readline error: %v\n", err)
		return
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		if strings.ToLower(line) == "exit" {
			fmt.Println("Exiting REPL.")
			break
		}

		result, err := EvalInput(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}

		fmt.Printf("Result: %v\n", result)
	}
}

func Eval(expr parse.Expr, env map[parse.Symbol]interface{}) (interface{}, error) {
	switch expr := expr.(type) {
	case parse.Symbol:
		if val, ok := env[expr]; ok {
			return val, nil
		}
		return nil, fmt.Errorf("undefined symbol: %s", expr)
	case parse.Number:
		return float64(expr), nil
	case parse.List:
		if len(expr) == 0 {
			return nil, fmt.Errorf("empty list")
		}
		op := expr[0]
		if len(expr) != 3 {
			return nil, fmt.Errorf("wrong number of arguments for %v", op)
		}
		arg1, err := Eval(expr[1], env)
		if err != nil {
			return nil, err
		}
		arg2, err := Eval(expr[2], env)
		if err != nil {
			return nil, err
		}
		switch op {
		case parse.Symbol("add"):
			return arg1.(float64) + arg2.(float64), nil
		case parse.Symbol("sub"):
			return arg1.(float64) - arg2.(float64), nil
		case parse.Symbol("mul"):
			return arg1.(float64) * arg2.(float64), nil
		case parse.Symbol("div"):
			return arg1.(float64) / arg2.(float64), nil
		default:
			return nil, fmt.Errorf("unknown operation: %v", op)
		}
	default:
		return nil, fmt.Errorf("unknown expression type: %T", expr)
	}
}

func EvalInput(input string) (interface{}, error) {
	parser := parse.NewParser(input)
	expr := parser.Parse()
	env := make(map[parse.Symbol]interface{})
	result, err := Eval(expr, env)
	return result, err
}