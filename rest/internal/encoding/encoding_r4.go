//go:build r4 || !(r4 || r4b || r5)

package encoding

import (
	"io"

	"github.com/DAMEDIC/fhir-toolbox-go/model"
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4"
)

func init() {
	decodeR4Resource = func(r io.Reader, format Format) (model.Resource, error) {
		contained, err := Decode[r4.ContainedResource](r, format)
		return contained.Resource, err
	}
}
