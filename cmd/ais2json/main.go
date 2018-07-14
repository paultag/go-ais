package main

import (
	"bufio"
	"encoding/json"
	// "fmt"
	"os"
	"strings"

	"pault.ag/go/ais/armor"
	"pault.ag/go/ais/messages"
	"pault.ag/go/ais/sixbit"
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

		sixbytes, err := armor.Decode(data)
		ohshit(err)

		slice, err := sixbit.Decode(sixbytes)
		ohshit(err)

		type_ := slice.Slice(0, 6).Uint()
		switch type_ {
		case 1, 2, 3:
			nav := messages.Position{}
			ohshit(messages.Unmarshal(slice, &nav))
			ohshit(json.NewEncoder(os.Stdout).Encode(nav))
		case 5:
			voy := messages.Voyage{}
			ohshit(messages.Unmarshal(slice, &voy))
			ohshit(json.NewEncoder(os.Stdout).Encode(voy))
		case 18:
			pr := messages.ClassBPosition{}
			ohshit(messages.Unmarshal(slice, &pr))
			ohshit(json.NewEncoder(os.Stdout).Encode(pr))
		case 21:
			aid := messages.NavigationAid{}
			ohshit(messages.Unmarshal(slice, &aid))
			ohshit(json.NewEncoder(os.Stdout).Encode(aid))
		case 24:
			static := messages.StaticData{}
			ohshit(messages.Unmarshal(slice, &static))
			ohshit(json.NewEncoder(os.Stdout).Encode(static))
		case 4:
			bs := messages.BaseStation{}
			ohshit(messages.Unmarshal(slice, &bs))
			ohshit(json.NewEncoder(os.Stdout).Encode(bs))
		default:
			continue
		}
	}

	ohshit(scanner.Err())
}
