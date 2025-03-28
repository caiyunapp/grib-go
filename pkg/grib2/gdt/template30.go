package gdt

import (
	"math"

	"github.com/caiyunapp/grib-go/pkg/grib2/regulation"
	"github.com/caiyunapp/walg/pkg/geo/grids"
	"github.com/caiyunapp/walg/pkg/geo/grids/lambert"
)

/*
Notes:
( 1) Grid lengths are in units of 10^-3 m, at the latitude specified by LaD.
( 2) If Latin 1 = Latin 2, then the projection is on a tangent cone.
( 3) The resolution flags (bits 3-4 of Flag table 3.3) are not applicable.
( 4) LoV is the longitude value of the meridian which is parallel to the Y-axis (or columns of the grid)
     along which latitude increases as the Y-coordinate increases (the orientation longitude may or may not
     appear on a particular grid).
( 5) A scaled value of radius of spherical Earth, or major or minor axis of oblate spheroid Earth is derived
     from applying appropriate scale factor to the value expressed in metres.
*/

// Template30 represents the Lambert conformal grid definition template
type Template30 struct {
	Template30FixedPart `json:"template30"`
	grids               grids.Grid `json:"-"`
}

// https://codes.ecmwf.int/grib/format/grib2/templates/3/30/
type template30FixedPart struct {
	ShapeOfTheEarth                             uint8
	ScaleFactorOfRadiusOfSphericalEarth         uint8
	ScaledValueOfRadiusOfSphericalEarth         uint32
	ScaleFactorOfMajorAxisOfOblateSpheroidEarth uint8
	ScaledValueOfMajorAxisOfOblateSpheroidEarth uint32
	ScaleFactorOfMinorAxisOfOblateSpheroidEarth uint8
	ScaledValueOfMinorAxisOfOblateSpheroidEarth uint32
	Nx                                          uint32
	Ny                                          uint32
	LatitudeOfFirstGridPoint                    int32
	LongitudeOfFirstGridPoint                   uint32
	ResolutionAndComponentFlags                 uint8
	LaD                                         int32
	LoV                                         uint32
	Dx                                          uint32
	Dy                                          uint32
	ProjectionCentreFlag                        uint8
	ScanningMode                                uint8
	Latin1                                      int32
	Latin2                                      int32
	LatitudeOfSouthernPole                      int32
	LongitudeOfSouthernPole                     uint32
}

func (t template30FixedPart) Export() Template {
	t30 := Template30FixedPart{
		ShapeOfTheEarth:                             regulation.ToInt8(t.ShapeOfTheEarth),
		ScaleFactorOfRadiusOfSphericalEarth:         regulation.ToInt8(t.ScaleFactorOfRadiusOfSphericalEarth),
		ScaledValueOfRadiusOfSphericalEarth:         regulation.ToInt32(t.ScaledValueOfRadiusOfSphericalEarth),
		ScaleFactorOfMajorAxisOfOblateSpheroidEarth: regulation.ToInt8(t.ScaleFactorOfMajorAxisOfOblateSpheroidEarth),
		ScaledValueOfMajorAxisOfOblateSpheroidEarth: regulation.ToInt32(t.ScaledValueOfMajorAxisOfOblateSpheroidEarth),
		ScaleFactorOfMinorAxisOfOblateSpheroidEarth: regulation.ToInt8(t.ScaleFactorOfMinorAxisOfOblateSpheroidEarth),
		ScaledValueOfMinorAxisOfOblateSpheroidEarth: regulation.ToInt32(t.ScaledValueOfMinorAxisOfOblateSpheroidEarth),
		Nx:                          regulation.ToInt32(t.Nx),
		Ny:                          regulation.ToInt32(t.Ny),
		LatitudeOfFirstGridPoint:    t.LatitudeOfFirstGridPoint,
		LongitudeOfFirstGridPoint:   regulation.ToInt32(t.LongitudeOfFirstGridPoint),
		ResolutionAndComponentFlags: regulation.ToInt8(t.ResolutionAndComponentFlags),
		LaD:                         t.LaD,
		LoV:                         regulation.ToInt32(t.LoV),
		Dx:                          regulation.ToInt32(t.Dx),
		Dy:                          regulation.ToInt32(t.Dy),
		ProjectionCentreFlag:        regulation.ToInt8(t.ProjectionCentreFlag),
		ScanningMode:                regulation.ToInt8(t.ScanningMode),
		Latin1:                      t.Latin1,
		Latin2:                      t.Latin2,
		LatitudeOfSouthernPole:      t.LatitudeOfSouthernPole,
		LongitudeOfSouthernPole:     regulation.ToInt32(t.LongitudeOfSouthernPole),
	}

	return t30.AsTemplate()
}

// Template30FixedPart represents the fixed part of Template 30
type Template30FixedPart struct {
	ShapeOfTheEarth                             int8  `json:"-"`
	ScaleFactorOfRadiusOfSphericalEarth         int8  `json:"-"`
	ScaledValueOfRadiusOfSphericalEarth         int32 `json:"-"`
	ScaleFactorOfMajorAxisOfOblateSpheroidEarth int8  `json:"-"`
	ScaledValueOfMajorAxisOfOblateSpheroidEarth int32 `json:"-"`
	ScaleFactorOfMinorAxisOfOblateSpheroidEarth int8  `json:"-"`
	ScaledValueOfMinorAxisOfOblateSpheroidEarth int32 `json:"-"`
	Nx                                          int32 `json:"nx"`
	Ny                                          int32 `json:"ny"`
	LatitudeOfFirstGridPoint                    int32 `json:"latitudeOfFirstGridPoint"`
	LongitudeOfFirstGridPoint                   int32 `json:"longitudeOfFirstGridPoint"`
	ResolutionAndComponentFlags                 int8  `json:"resolutionAndComponentFlags"`
	LaD                                         int32 `json:"laD"`
	LoV                                         int32 `json:"loV"`
	Dx                                          int32 `json:"dx"`
	Dy                                          int32 `json:"dy"`
	ProjectionCentreFlag                        int8  `json:"projectionCentreFlag"`
	ScanningMode                                int8  `json:"scanningMode"`
	Latin1                                      int32 `json:"latin1"`
	Latin2                                      int32 `json:"latin2"`
	LatitudeOfSouthernPole                      int32 `json:"latitudeOfSouthernPole"`
	LongitudeOfSouthernPole                     int32 `json:"longitudeOfSouthernPole"`
}

// AsTemplate converts the fixed part to a complete template
func (t *Template30FixedPart) AsTemplate() Template {
	return &Template30{
		Template30FixedPart: *t,
		grids: lambert.NewConformaConic(lambert.ConformaConicParams{
			Nx:                int(t.Nx),
			Ny:                int(t.Ny),
			StandardParallel1: float64(t.Latin1) * 1e-6,
			StandardParallel2: float64(t.Latin2) * 1e-6,
			CentralMeridian:   float64(t.LoV) * 1e-6,
			OriginLatitude:    float64(t.LaD) * 1e-6,
			FirstGridPointLat: float64(t.LatitudeOfFirstGridPoint) * 1e-6,
			FirstGridPointLon: float64(t.LongitudeOfFirstGridPoint) * 1e-6,
			Dx:                float64(t.Dx) * 1e-3, // Grid lengths are in units of 10^-3 m
			Dy:                float64(t.Dy) * 1e-3, // Grid lengths are in units of 10^-3 m
			ScanningMode:      int(t.ScanningMode),
			SouthPoleLat:      float64(t.LatitudeOfSouthernPole) * 1e-6,
			SouthPoleLon:      float64(t.LongitudeOfSouthernPole) * 1e-6,
			ProjectionCenter:  int(t.ProjectionCentreFlag),
			ShapeOfEarth:      grids.ShapeOfEarth(t.ShapeOfTheEarth),
			EarthRadius:       float64(t.ScaledValueOfRadiusOfSphericalEarth) * math.Pow(10, float64(t.ScaleFactorOfRadiusOfSphericalEarth)),
		}),
	}
}

// GetNx returns the number of points along the X-axis
func (t *Template30FixedPart) GetNi() int32 {
	return t.Nx
}

// GetNy returns the number of points along the Y-axis
func (t *Template30FixedPart) GetNj() int32 {
	return t.Ny
}

// GetGridIndex returns the grid index for a given latitude and longitude
func (t *Template30) GetGridIndex(lat, lon float32) (n int) {
	latIdx, lonIdx := t.grids.GuessNearestIndex(float64(lat), float64(lon))
	return grids.GridIndexFromIndices(t.grids, latIdx, lonIdx, grids.ScanMode(t.ScanningMode))
}

// GetGridPoint returns the latitude and longitude for a given grid index
func (t *Template30) GetGridPoint(n int) (float32, float32, bool) {
	lat, lon, ok := grids.GridPoint(t.grids, n, grids.ScanMode(t.ScanningMode))
	return float32(lat), float32(lon), ok
}
