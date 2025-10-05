//go:build r4b || !(r4 || r4b || r5)

package encoding

import (
	"github.com/DAMEDIC/fhir-toolbox-go/model/gen/r4b"
	"io"

	"github.com/DAMEDIC/fhir-toolbox-go/model"
)

func init() {
	decodeR4BResource = func(r io.Reader, format Format) (model.Resource, error) {
		contained, err := decode[model.R4B, r4b.ContainedResource](r, format)
		return contained.Resource, err
	}
}
