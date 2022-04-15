package json

import (
	"encoding/json"
	"fmt"
)

type Format struct {
	Data map[string]any
}

func (f Format) Format() string {
	var (
		bytes	[]byte
		err  	error
	)

	if bytes, err = json.MarshalIndent(f.Data, "", "    "); err != nil {
		return ""
	}
	return fmt.Sprint(bytes)
}