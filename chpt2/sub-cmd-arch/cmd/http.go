package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type httpCfg struct {
	url  string
	verb string
}

func HandleHttp(w io.Writer, args []string) error {
	var v string
	fs := flag.NewFlagSet("http", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&v, "verb", "GET", "HTTP method")
	switch v {
	case "GET":
		fallthrough
	case "POST":
		fallthrough
	case "HEAD":
	default:
		os.Exit(1)
	}
	fs.Usage = func() {
		var usageStr = `
http: A HTTP client.
		
http: <options> server` + "\n\n" + "Options:\n"
		fmt.Fprintf(w, usageStr)
		fs.PrintDefaults()
	}

	err := fs.Parse(args)
	if err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return ErrNoServerSpecified
	}
	c := httpCfg{verb: v}
	c.url = fs.Arg(0)
	fmt.Println(w, "Executing HTTP command")
	return nil
}
