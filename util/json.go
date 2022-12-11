package util

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
)

func JsonPrint(data any) {
	b, _ := json.Marshal(data)

	var out bytes.Buffer

	err := json.Indent(&out, b, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}

	out.WriteTo(os.Stdout)
}
