package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/fxamacker/cbor/v2"
	"github.com/veraison/psatoken"
)

func main() {
	var p string

	flag.StringVar(&p, "p", "2", "PSA token profile (1, 2, cca)")

	flag.Parse()

	r := bufio.NewReader(os.Stdin)

	in, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("reading JSON from stdin: %v", err)
	}

	var c interface{}

	switch p {
	case "1":
		c = &psatoken.P1Claims{}
	case "2":
		c = &psatoken.P2Claims{}
	case "cca":
		c = &psatoken.CcaPlatformClaims{}
	default:
		log.Fatalf("unknown profile %q", p)
	}

	err = json.Unmarshal(in, &c)
	if err != nil {
		log.Fatalf("decoding JSON: %v", err)
	}

	out, err := cbor.Marshal(c)
	if err != nil {
		log.Fatalf("encoding to CBOR: %v", err)
	}

	fmt.Printf("%x\n", out)
}
