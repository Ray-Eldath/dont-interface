package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/Ray-Eldath/dont-interface/cmd"
)

type arrayFlags []string

func (i *arrayFlags) String() string { return strings.Join(*i, " ") }
func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var filesFlag arrayFlags

func main() {
	flag.Var(&filesFlag, "file", "A single file that you want to include. Set this multiple times if you have many files.")
	flag.Parse()
	if len(filesFlag) <= 0 {
		fmt.Println("Error: no file specified, check out --help for usage.")
		os.Exit(255)
	}

	r, err := cmd.Calculate(filesFlag)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Among your codebase, there are...")
	fmt.Printf("    %d parameters declared in functions, %d of them are evil! (%s typed)\n", r.TotalParams, r.EvilParams, percentage(r.EvilResults, r.TotalResults))
	fmt.Printf("    %d values returned from functions, %d of them are evil! (%s typed)\n", r.TotalResults, r.EvilResults, percentage(r.EvilResults, r.TotalResults))
	fmt.Printf("    %d fields declared in structs, %d of them are evil! (%s typed)\n", r.TotalStructField, r.EvilStructField, percentage(r.EvilStructField, r.TotalStructField))
	fmt.Printf("    %d values declared, %d of them are evil! (%s typed)\n", r.TotalValueDecl, r.EvilValueDecl, percentage(r.EvilValueDecl, r.TotalValueDecl))
	fmt.Printf("    %d type aliases introduced, %d of them are evil! (%s typed)\n", r.TotalTypeAlias, r.EvilTypeAlias, percentage(r.EvilTypeAlias, r.TotalTypeAlias))
	fmt.Printf("Overall, %s of your types are strictly typed (not interface{}).", sum(*r))
}

func sum(visitor cmd.Visitor) string {
	t := reflect.TypeOf(visitor)
	v := reflect.ValueOf(visitor)
	var (
		total int
		evil  int
	)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		if strings.HasPrefix(name, "Total") {
			total += int(v.Field(i).Int())
		} else if strings.HasPrefix(name, "Evil") {
			evil += int(v.Field(i).Int())
		}
	}
	return percentage(evil, total)
}

func percentage(u int, d int) string {
	if d == 0 {
		return "100%"
	}
	return fmt.Sprintf("%.2f%%", (float64(d-u)/float64(d))*100)
}
