package pack

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func Pack_bool_data(value bool) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)

	if err !=nil {
		log.Fatal(err)
	}

	data := buf.Bytes()

	fmt.Println(data)

	return  data
}

func init() {
	Pack_bool_data(true)
}
