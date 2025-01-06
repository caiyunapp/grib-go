package gridpoint

import (
	"fmt"

	"github.com/scorix/grib-go/internal/pkg/bitio"
)

type Unknown struct {
	numVals int
}

func NewUnknown(numVals int) *Unknown {
	return &Unknown{numVals: numVals}
}

func (u *Unknown) GetNumVals() int {
	return u.numVals
}

func (u *Unknown) Definition() any {
	return nil
}

func (u *Unknown) ReadAllData(r *bitio.Reader) ([]float32, error) {
	return nil, fmt.Errorf("unknown data template")
}
