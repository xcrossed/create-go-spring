package create

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Options struct {
	ProjectName   string
	OutputDir     string
	SupportModule string
	SupportWeb    bool
	SupportRedis  bool
	SupportMySQL  bool
	PrintVersion  bool
}

func setOptions(opts *Options, flagset *flag.FlagSet) {
	flagset.StringVar(&opts.ProjectName, "p", "gospring-demo", "project name")
	flagset.StringVar(&opts.OutputDir, "o", "gospring-demo", "project store directory,default current pwd")
	flagset.StringVar(&opts.SupportModule, "m", "all", "support module.etc mysql,redis,web,default use all contains(web,mysql,web)")
	flagset.BoolVar(&opts.PrintVersion, "v", false, "print version and exit")

	flagset.Usage = func() {
		_, _ = fmt.Fprintf(flagset.Output(),
			"Usage of %s:OPTIONS - \n",
			filepath.Base(os.Args[0]),
		)
		flagset.PrintDefaults()
	}
}

func NewOptions() *Options {
	opts := &Options{}
	setOptions(opts, flag.CommandLine)

	err := flag.CommandLine.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	flag.Usage = flag.CommandLine.Usage

	if opts.PrintVersion {
		fmt.Println("create-go-spring v" + Version)
	}

	if flag.NFlag() != 1 {
		if opts.PrintVersion {
			os.Exit(0)
		} else {
			flag.Usage()
			os.Exit(1)
		}
	}
	return opts
}

func Usage() string {
	cmdline := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	setOptions(&Options{}, cmdline)
	buf := &bytes.Buffer{}
	cmdline.SetOutput(buf)
	cmdline.Usage()
	return buf.String()
}
