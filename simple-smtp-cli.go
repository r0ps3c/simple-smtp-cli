package main

import (
	"log"
	"flag"
	"os"
	"fmt"
	"golang.org/x/text/transform"
	"github.com/emersion/go-smtp"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] <recipient> [.. <recipient>]:\n", os.Args[0])
	flag.PrintDefaults()
}


type ToCRLF struct{}

func (ToCRLF) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	for nDst < len(dst) && nSrc < len(src) {
		if c := src[nSrc]; c == '\n' {
			if nDst+1 == len(dst) {
				break
			}
			dst[nDst] = '\r'
			dst[nDst+1] = '\n'
			nSrc++
			nDst += 2
		} else {
			dst[nDst] = c
			nSrc++
			nDst++
		}
	}
	if nSrc < len(src) {
		err = transform.ErrShortDst
	}
	return
}

func (ToCRLF) Reset() {}

func main() {
	// set up and process commandline args
	var verbose=flag.Bool("verbose",false,"verbose logging")
	var smtpaddr=flag.String("smtpaddr","localhost:25","smtp name:port to connect to")
	var mailfrom=flag.String("f","nobody","mail from address")

	flag.Parse()

	if flag.NArg() == 0 {
		Usage()
		os.Exit(1)
	}

	// Connect to the remote SMTP server.
	c, err := smtp.Dial(*smtpaddr)
	if err != nil {
		log.Fatal(err)
	}
	
	defer c.Close()

	// Set the sender and recipient.
	if(*verbose) {
		log.Printf("mail from=%s",*mailfrom)
	}

	if(*verbose) {
		log.Printf("rcpt to=%s",flag.Args())
	}

	msgstdinreader := transform.NewReader(os.Stdin, ToCRLF{})

	err = c.SendMail(*mailfrom, flag.Args(), msgstdinreader)

	if err != nil {
		log.Fatal(err)
	}
}
