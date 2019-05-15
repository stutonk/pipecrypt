package main

import (
	"encoding/hex"
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/stutonk/boxutil"
)

const (
	errFmt   = "%v: fatal; %v\n"
	usageFmt = `usage: %v [-h, -v] key
Reads from STDIN; writes to STDOUT
Options are:
`
	verFmt  = "%v version %v\n"
	version = "1.0.0"
)

var (
	appName     = os.Args[0]
	decryptFlag bool
	input       = os.Stdin
	helpFlag    bool
	output      = os.Stdout
	verFlag     bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usageFmt, appName)
		flag.PrintDefaults()
		fmt.Println()
	}
	flag.BoolVarP(
		&decryptFlag,
		"decrypt",
		"d",
		false,
		"operate in decrypt mode",
	)
	flag.BoolVarP(
		&helpFlag,
		"help",
		"h",
		false,
		"display this help and exit",
	)
	flag.BoolVarP(
		&verFlag,
		"version",
		"v",
		false,
		"output version information and exit",
	)
	flag.Parse()
}

func main() {
	switch {
	case helpFlag:
		flag.Usage()
		return
	case verFlag:
		fmt.Printf(verFmt, appName, version)
		return
	case len(flag.Args()) == 0:
		flag.Usage()
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, errFmt, appName, r)
		}
	}()

	keyBytes, err := hex.DecodeString(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	if len(keyBytes) < 32 {
		panic(fmt.Errorf("keys must be 64 characters (32 bytes) long"))
	}
	var key [32]byte
	copy(key[:], keyBytes[:32])

	if decryptFlag {
		if err := boxutil.OpenStream(input, output, &key); err != nil {
			panic(err)
		}
	} else {
		if err := boxutil.SealStream(input, output, &key); err != nil {
			panic(err)
		}
	}
}
