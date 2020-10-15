package main

import (
	"flag"
	"strings"
	"os"
	lab2 "github.com/G-V-G/l2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()
	// handler := lab2.ComputeHandler{Input: strings.NewReader("2 5 +"), Output: os.Stdout}
	nf, _ := os.Create(`go-file.txt`)
	handler := lab2.ComputeHandler{Input: strings.NewReader("2 5 7 ^ +"), Output: nf}
	// nf, _ := os.Open(`go-file.txt`)
	// handler := lab2.ComputeHandler{Input: nf, Output: os.Stdout}
	err := handler.Compute()
	if err != nil {
		panic(err)
	}
	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

	// res, _ := lab2.PrefixToPostfix("+ 2 2")
	// fmt.Println(res)
}
