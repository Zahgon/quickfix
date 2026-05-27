package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"sync"
	"text/template"

	"github.com/quickfixgo/quickfix/cmd/generate-fix/internal"
	"github.com/quickfixgo/quickfix/datadictionary"
)

var (
	waitGroup sync.WaitGroup
	errors    = make(chan error)
)

func usage() { _ = "STUB: not implemented"; return }

func getPackageName(fixSpec *datadictionary.DataDictionary) string {
	_ = "STUB: not implemented"
	return ""
}

func getTransportPackageName(fixSpec *datadictionary.DataDictionary) string {
	_ = "STUB: not implemented"
	return ""
}

type component struct {
	Package          string
	FIXPackage       string
	TransportPackage string
	FIXSpec          *datadictionary.DataDictionary
	Name             string
	*datadictionary.MessageDef
}

func genHeader(pkg string, spec *datadictionary.DataDictionary) { _ = "STUB: not implemented"; return }

func genTrailer(pkg string, spec *datadictionary.DataDictionary) { _ = "STUB: not implemented"; return }

func genMessage(fixPkg string, spec *datadictionary.DataDictionary, msg *datadictionary.MessageDef) {
	_ = "STUB: not implemented"
	return
}

func genTags() { _ = "STUB: not implemented"; return }

func genFields() { _ = "STUB: not implemented"; return }

func genEnums() { _ = "STUB: not implemented"; return }

func gen(t *template.Template, fileOut string, data interface{}) { _ = "STUB: not implemented"; return }

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
	}

	args := flag.Args()
	if len(args) == 1 {
		dictpath := args[0]
		if strings.Contains(dictpath, "FIX50SP1") {
			args = append(args, strings.Replace(dictpath, "FIX50SP1", "FIXT11", -1))
		} else if strings.Contains(dictpath, "FIX50SP2") {
			args = append(args, strings.Replace(dictpath, "FIX50SP2", "FIXT11", -1))
		} else if strings.Contains(dictpath, "FIX50") {
			args = append(args, strings.Replace(dictpath, "FIX50", "FIXT11", -1))
		}
	}
	specs := []*datadictionary.DataDictionary{}

	for _, dataDictPath := range args {
		spec, err := datadictionary.Parse(dataDictPath)
		if err != nil {
			log.Fatalf("Error Parsing %v: %v", dataDictPath, err)
		}
		specs = append(specs, spec)
	}

	internal.BuildGlobalFieldTypes(specs)

	waitGroup.Add(1)
	go genTags()
	waitGroup.Add(1)
	go genFields()
	waitGroup.Add(1)
	go genEnums()

	for _, spec := range specs {
		pkg := getPackageName(spec)

		if fi, err := os.Stat(pkg); os.IsNotExist(err) {
			if err := os.Mkdir(pkg, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		} else if !fi.IsDir() {
			log.Fatalf("%v/ is not a directory", pkg)
		}

		switch pkg {
		// Uses fixt11 header/trailer.
		case "fix50", "fix50sp1", "fix50sp2":
		default:
			waitGroup.Add(1)
			go genHeader(pkg, spec)

			waitGroup.Add(1)
			go genTrailer(pkg, spec)
		}

		for _, m := range spec.Messages {
			waitGroup.Add(1)
			go genMessage(pkg, spec, m)
		}
	}

	go func() {
		waitGroup.Wait()
		close(errors)
	}()

	var h internal.ErrorHandler
	for err := range errors {
		h.Handle(err)
	}

	os.Exit(h.ReturnCode)
}
