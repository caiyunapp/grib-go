package gridpoint

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/caiyunapp/grib-go/internal/pkg/bitio"
	"github.com/caiyunapp/grib-go/pkg/grib2/drt/datapacking"
	"github.com/caiyunapp/grib-go/pkg/grib2/drt/definition"
	"github.com/caiyunapp/grib-go/pkg/grib2/regulation"
)

type SimplePacking struct {
	ReferenceValue     float32 // 12-15
	BinaryScaleFactor  int16   // 16-17
	DecimalScaleFactor int16   // 18-19
	Bits               uint8   // 20
	Type               int8    // 21
	NumVals            int
}

func NewSimplePacking(def definition.SimplePacking, numVals int) *SimplePacking {
	return &SimplePacking{
		ReferenceValue:     def.R,
		BinaryScaleFactor:  regulation.ToInt16(def.B),
		DecimalScaleFactor: regulation.ToInt16(def.D),
		Bits:               def.L,
		Type:               regulation.ToInt8(def.T),
		NumVals:            numVals,
	}
}

func (sp *SimplePacking) ScaleFunc() func(uint32) float32 {
	return datapacking.SimpleScaleFunc(sp.BinaryScaleFactor, sp.DecimalScaleFactor, sp.ReferenceValue)
}

func (sp *SimplePacking) ReadAllData(r *bitio.Reader) ([]float32, error) {
	var (
		values    []float32
		scaleFunc = sp.ScaleFunc()
	)

	if sp.Bits == 0 {
		for range sp.NumVals {
			values = append(values, scaleFunc(0))
		}
	}

	for sp.Bits > 0 {
		bitsVal, err := r.ReadBits(sp.Bits)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

		values = append(values, scaleFunc(uint32(bitsVal)))
	}

	if len(values) != sp.NumVals {
		return nil, fmt.Errorf("expected %d values, got %d", sp.NumVals, len(values))
	}

	return values, nil
}

func (sp *SimplePacking) GetNumVals() int {
	return sp.NumVals
}

func (sp *SimplePacking) Definition() any {
	return definition.SimplePacking{
		R: sp.ReferenceValue,
		B: regulation.ToUint16(sp.BinaryScaleFactor),
		D: regulation.ToUint16(sp.DecimalScaleFactor),
		L: sp.Bits,
		T: regulation.ToUint8(sp.Type),
	}
}

type SimplePackingReader struct {
	r  io.ReaderAt
	sp *SimplePacking
	sf func(uint32) float32
}

func NewSimplePackingReader(r io.ReaderAt, start, end int64, sp *SimplePacking) *SimplePackingReader {
	return &SimplePackingReader{
		r:  io.NewSectionReader(r, start, end-start),
		sp: sp,
		sf: sp.ScaleFunc(),
	}
}

func (r *SimplePackingReader) ReadGridAt(ctx context.Context, n int) (float32, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}

	if n >= r.sp.NumVals {
		return 0, fmt.Errorf("grid point %d is out of range[0-%d]", n, r.sp.NumVals)
	}

	bitsOffset := n * int(r.sp.Bits)
	skipBits := bitsOffset % 8
	needBytes := int(math.Ceil(float64(int(r.sp.Bits)+skipBits) / float64(8.0)))

	bs := make([]byte, needBytes)
	if _, err := r.r.ReadAt(bs, int64(bitsOffset/8)); err != nil {
		return 0, fmt.Errorf("read %d bytes at offset %d: %w", needBytes, bitsOffset/8, err)
	}

	u, err := bitio.ReadBits(bs, uint8(skipBits), uint8(r.sp.Bits))
	if err != nil {
		return 0, fmt.Errorf("read %d bits: %w", r.sp.Bits, err)
	}

	return r.sf(uint32(u)), nil
}
