package gdt_test

import (
	"testing"

	"github.com/caiyunapp/grib-go/pkg/grib2/gdt"
	"github.com/stretchr/testify/assert"
)

func TestTemplate30_GetGridIndex(t *testing.T) {
	template := (&gdt.Template30FixedPart{
		ShapeOfTheEarth:                             6,
		ScaleFactorOfRadiusOfSphericalEarth:         0,
		ScaledValueOfRadiusOfSphericalEarth:         0,
		ScaleFactorOfMajorAxisOfOblateSpheroidEarth: 0,
		ScaledValueOfMajorAxisOfOblateSpheroidEarth: 0,
		ScaleFactorOfMinorAxisOfOblateSpheroidEarth: 0,
		ScaledValueOfMinorAxisOfOblateSpheroidEarth: 0,
		Nx:                          759,
		Ny:                          599,
		LatitudeOfFirstGridPoint:    7300,
		LongitudeOfFirstGridPoint:   78307,
		ResolutionAndComponentFlags: 8,
		LaD:                         30000,
		LoV:                         105000,
		Dx:                          9000,
		Dy:                          9000,
		ProjectionCentreFlag:        0,
		ScanningMode:                64,
		Latin1:                      30000,
		Latin2:                      60000,
		LatitudeOfSouthernPole:      0,
		LongitudeOfSouthernPole:     0,
	}).AsTemplate()

	assert.Equal(t, 759, int(template.GetNi()))
	assert.Equal(t, 599, int(template.GetNj()))

	// Test a few sample points
	testCases := []struct {
		name     string
		lat      float32
		lon      float32
		expected int
	}{
		{
			name:     "first point",
			lat:      7300 * 1e-6,
			lon:      78307 * 1e-6,
			expected: 0,
		},
		{
			name:     "middle point",
			lat:      0.03150075,
			lon:      0.10898285,
			expected: 299*759 + 379,
		},
		{
			name:     "last point",
			lat:      73.0,
			lon:      78.307,
			expected: 598*759 + 758,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := template.GetGridIndex(tc.lat, tc.lon)
			if got != tc.expected {
				t.Errorf("GetGridIndex(%f, %f) = %d; want %d", tc.lat, tc.lon, got, tc.expected)
			}
		})
	}

}

func TestTemplate30_GetGridPoint(t *testing.T) {
	template := (&gdt.Template30FixedPart{
		ShapeOfTheEarth:                             6,
		ScaleFactorOfRadiusOfSphericalEarth:         0,
		ScaledValueOfRadiusOfSphericalEarth:         0,
		ScaleFactorOfMajorAxisOfOblateSpheroidEarth: 0,
		ScaledValueOfMajorAxisOfOblateSpheroidEarth: 0,
		ScaleFactorOfMinorAxisOfOblateSpheroidEarth: 0,
		ScaledValueOfMinorAxisOfOblateSpheroidEarth: 0,
		Nx:                          759,
		Ny:                          599,
		LatitudeOfFirstGridPoint:    7300,
		LongitudeOfFirstGridPoint:   78307,
		ResolutionAndComponentFlags: 8,
		LaD:                         30000,
		LoV:                         105000,
		Dx:                          9000,
		Dy:                          9000,
		ProjectionCentreFlag:        0,
		ScanningMode:                64,
		Latin1:                      30000,
		Latin2:                      60000,
		LatitudeOfSouthernPole:      0,
		LongitudeOfSouthernPole:     0,
	}).AsTemplate()

	assert.Equal(t, 759, int(template.GetNi()))
	assert.Equal(t, 599, int(template.GetNj()))

	lat, lon, ok := template.GetGridPoint(0)
	assert.True(t, ok)
	assert.InDelta(t, 0.0073, lat, 1e-6)
	assert.InDelta(t, 0.078307, lon, 1e-6)

	lat, lon, ok = template.GetGridPoint(299*759 + 379)
	assert.True(t, ok)
	assert.InDelta(t, 0.03150075, lat, 1e-6)
	assert.InDelta(t, 0.10898285, lon, 1e-6)

	lat, lon, ok = template.GetGridPoint(598*759 + 758)
	assert.True(t, ok)
	assert.InDelta(t, 0.055701494, lat, 1e-6)
	assert.InDelta(t, 0.13965872307, lon, 1e-6)
}
