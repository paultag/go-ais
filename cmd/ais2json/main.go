package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"

	"pault.ag/go/ais"
	// "pault.ag/go/ais/armor"
	// "pault.ag/go/ais/messages"
	// "pault.ag/go/ais/sixbit"
)

func ohshit(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// fd, err := os.Open(os.Args[1])
	// ohshit(err)
	// defer fd.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")[5]

		message, err := ais.Decode(data)
		ohshit(err)

		parsed, err := message.Parse()
		ohshit(err)

		// loc, ok := parsed.(messages.Locatable)
		// if !ok {
		// 	continue
		// }

		ohshit(json.NewEncoder(os.Stdout).Encode(parsed))
	}

	ohshit(scanner.Err())
}
