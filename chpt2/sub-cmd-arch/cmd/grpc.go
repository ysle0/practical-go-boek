package cmd

import (
	"flag"
	"fmt"
	"io"
)

type grpcCfg struct {
	server string
	method string
	body   string
}

func HandleGrpc(w io.Writer, args []string) error {
	c := grpcCfg{}
	fs := flag.NewFlagSet("grpc", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.StringVar(&c.method, "method", "", "method to call")
	fs.StringVar(&c.body, "body", "", "body of request")
	fs.Usage = func() {
		usageStr := `
grpc: A gRPC client
		
grpc: <options> server` + "\n\n" + "Options:\n"
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
	c.server = fs.Arg(0)
	fmt.Fprintln(w, "Executing grpc command")
	return nil
}
