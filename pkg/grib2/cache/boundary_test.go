package cache_test

import (
	"context"
	"testing"

	"github.com/caiyunapp/grib-go/pkg/grib2/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoundary(t *testing.T) {
	ds := &mockGridDataSource{gridValue: 100}
	bc := cache.NewBoundary(0, 10, 0, 10, ds, cache.NewMapStore())

	// first read should be from source
	v, err := bc.ReadGridAt(context.TODO(), 1, 1, 1)
	require.NoError(t, err)
	assert.Equal(t, float32(100), v)
	assert.Equal(t, 1, ds.readCount)

	// second read should be cached
	v, err = bc.ReadGridAt(context.TODO(), 1, 1, 1)
	require.NoError(t, err)
	assert.Equal(t, float32(100), v)
	assert.Equal(t, 1, ds.readCount)
}
