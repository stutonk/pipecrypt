[![Go Report Card](https://goreportcard.com/badge/github.com/stutonk/pipecrypt)](https://goreportcard.com/report/github.com/stutonk/pipecrypt)  
pipecrypt is a lightweight alternatve to PGP for securely encrypting streams
using NaCl's secretbox. It expects a hexidecimal string as an argument which
should be 64 characters (32 bytes) long \-\- additional characters are
ignored. The program reads from STDIN and writes to STDOUT.

```
usage: ./pipecrypt [-h, -v] key
Reads from STDIN; writes to STDOUT
Options are:
  -d, --decrypt   operate in decrypt mode
  -h, --help      display this help and exit
  -v, --version   output version information and exit
```

### release binaries
are available [here](https://github.com/stutonk/pipecrypt/releases) for amd64/all major OSes

### for unixes
`make && make install`

### everybody else
`go build`

### note
This program was primarily designed to be used in scripts; be mindful of
your shell history when using it in interactive mode. The best practice is
to use your shell's command substitution functionality to read the key from
a file or output of some other program.

A simple tool for generating keys can be found [here](https://github.com/stutonk/passkey).