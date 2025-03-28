package regulation_test

import (
	"math"
	"testing"

	"github.com/caiyunapp/grib-go/pkg/grib2/regulation"
	"github.com/stretchr/testify/assert"
)

func TestToInt32(t *testing.T) {
	t.Parallel()

	assert.Equal(t, int32(-90000000), regulation.ToInt32(2237483648))
	assert.Equal(t, int32(359750000), regulation.ToInt32(359750000))

	assert.Equal(t, int32(-1), regulation.ToInt32(math.MaxUint32))
	assert.Equal(t, int32(math.MaxInt32), regulation.ToInt32(math.MaxInt32))
}

func TestToInt16(t *testing.T) {
	t.Parallel()

	assert.Equal(t, int16(-1), regulation.ToInt16(math.MaxUint16))
	assert.Equal(t, int16(math.MaxInt16), regulation.ToInt16(math.MaxInt16))
}

func TestToInt8(t *testing.T) {
	t.Parallel()

	assert.Equal(t, int8(103), regulation.ToInt8(103))

	assert.Equal(t, int8(-1), regulation.ToInt8(math.MaxUint8))
	assert.Equal(t, int8(math.MaxInt8), regulation.ToInt8(math.MaxInt8))
}

func TestToInt(t *testing.T) {
	t.Parallel()

	assert.Equal(t, -64651, regulation.ToInt(0b100000001111110010001011, 24))
}

func TestToUint(t *testing.T) {
	t.Parallel()

	assert.Equal(t, uint(255), regulation.ToUint(255, 8))
	assert.Equal(t, uint(65535), regulation.ToUint(65535, 16))
	assert.Equal(t, uint(16777215), regulation.ToUint(16777215, 24))
	assert.Equal(t, uint(4294967295), regulation.ToUint(4294967295, 32))

	assert.EqualValues(t, uint16(math.MaxUint16), uint16(regulation.ToUint(-1, 16)))
	assert.EqualValues(t, uint32(2237483648), uint32(regulation.ToUint(-90000000, 32)))
}

func TestIsMissingValue(t *testing.T) {
	t.Parallel()

	assert.Equal(t, true, regulation.IsMissingValue(255, 8))
	assert.Equal(t, true, regulation.IsMissingValue(65535, 16))

	i := -1
	assert.Equal(t, true, regulation.IsMissingValue(uint(i), 8))
	assert.Equal(t, true, regulation.IsMissingValue(uint(i), 16))
	assert.Equal(t, true, regulation.IsMissingValue(uint(i), 24))
	assert.Equal(t, true, regulation.IsMissingValue(uint(i), 32))
}

func TestDegreedLatitudeLongitude(t *testing.T) {
	t.Parallel()

	l := 269.250000
	assert.EqualValues(t, l, regulation.DegreedLatitudeLongitude(int(l*1e6)))
}
