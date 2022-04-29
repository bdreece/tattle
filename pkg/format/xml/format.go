package xml

import (
	"encoding/xml"
)

type Format struct {
	data map[string]any
}

func (f Format) Builder() Builder {
	return Builder{defaults: f}
}

func (f Format) Format() string {
	v, _ := xml.MarshalIndent(f.data, "", "	")
	return string(v)
}
