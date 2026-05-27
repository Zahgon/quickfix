package internal

import (
	"flag"
	"go/ast"
	"go/printer"
	"go/token"
)

var (
	useFloat    = flag.Bool("use-float", false, "By default, FIX float fields are represented as arbitrary-precision fixed-point decimal numbers.  Set to 'true' to instead generate FIX float fields as float64 values.")
	useUDecimal = flag.Bool("use-udecimal", false, "By default, FIX uses the shopspring/decimal library for fixed-point decimal numbers.  Set to 'true' to instead use the quagmt/udecimal library.")
	pkgRoot     = flag.String("pkg-root", "github.com/quickfixgo", "Set a string here to provide a custom import path for generated packages.")
	tabWidth    = 8
	printerMode = printer.UseSpaces | printer.TabIndent
)

// ParseError indicates generated go source is invalid
type ParseError struct {
	path string
	err  error
}

func (e ParseError) Error() string { _ = "STUB: not implemented"; return "" }

// ErrorHandler is a convenience struct for interpretting generation Errors
type ErrorHandler struct {
	ReturnCode int
}

// Handle interprets the generation error. Proceeds with setting returnCode, or panics depending on error type
func (h *ErrorHandler) Handle(err error) { _ = "STUB: not implemented"; return }

//do nothing

func write(filePath string, fset *token.FileSet, f *ast.File) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteFile parses the generated code in fileOut and writes the code out to filePath.
// Function performs some import clean up and gofmts the code before writing
// Returns ParseError if the generated source is invalid but is written to filePath
func WriteFile(filePath, fileOut string) error { _ = "STUB: not implemented"; return nil }

//write out the file regardless of parseFile errors
