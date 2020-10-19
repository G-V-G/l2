package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	lab2 "github.com/G-V-G/l2"
)

var (
	// -e key used to enter expression through the command line
	inputExpression = flag.String("e", "", "Expression to compute")
	// -f key used to enter expression through the file
	inputFilename = flag.String("f", "", "input file")
	// -e and -f can not be combined!

	// -o key used to output the result to the file mentioned
	outputFilename = flag.String("o", "", "output file")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()

	switch {
	case *inputFilename != "" && *inputExpression != "":
		err := "error: only one source of expression needed\n"
		fmt.Printf(err)
		return
	// If both filename and expression are full-filled
	case *inputFilename != "":
		{
			if *outputFilename != "" {
				input, _ := os.Open(*inputFilename)
				nf, _ := os.Create(*outputFilename)
				handler := lab2.ComputeHandler{Input: input, Output: nf}
				err := handler.Compute()
				if err != nil {
					panic(err)
				}
				// If we have the output filename, we create a file by the name, and open file for input
			} else {
				input, _ := os.Open(*inputFilename)
				handler := lab2.ComputeHandler{Input: input, Output: os.Stdout}
				err := handler.Compute()
				if err != nil {
					panic(err)
				}
				// Exact the same, but we don't create the output file
			}

		}
	// If filename full-filled
	default:
		{
			if *outputFilename != "" {
				nf, _ := os.Create(*outputFilename)
				handler := lab2.ComputeHandler{Input: strings.NewReader(*inputExpression), Output: nf}
				err := handler.Compute()
				if err != nil {
					panic(err)
				}
			} else {
				handler := lab2.ComputeHandler{Input: strings.NewReader(*inputExpression), Output: os.Stdout}
				err := handler.Compute()
				if err != nil {
					panic(err)
				}
			}
		}
		// If expression full-filled
	}

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()
}
