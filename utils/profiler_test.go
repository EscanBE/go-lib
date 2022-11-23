package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestNewMasterProfiler(t *testing.T) {
	assert.Nil(t, NewMasterProfiler("", "", false))

	desc := fmt.Sprintf("desc %d", rand.Int())
	constDesc := fmt.Sprintf("constDesc %d", rand.Int())
	profiler := NewMasterProfiler(desc, constDesc, true)
	now := time.Now().UnixMilli()
	assert.NotNil(t, profiler)
	assert.Equal(t, constDesc, profiler.constDesc)
	assert.True(t, math.Abs(float64(now-profiler.start)) < 10)
	assert.Zero(t, profiler.duration)
	assert.False(t, profiler.finalized)
	assert.Zero(t, profiler.level)
	assert.False(t, profiler.err)
	assert.Nil(t, profiler.children)
}

func Test_newProfiler(t *testing.T) {
	desc := fmt.Sprintf("desc %d", rand.Int())
	level := rand.Int()
	profiler := newProfiler(desc, level)
	now := time.Now().UnixMilli()
	assert.NotNil(t, profiler)
	assert.Empty(t, profiler.constDesc)
	assert.True(t, math.Abs(float64(now-profiler.start)) < 10)
	assert.Zero(t, profiler.duration)
	assert.False(t, profiler.finalized)
	assert.Equal(t, level, profiler.level)
	assert.False(t, profiler.err)
	assert.Nil(t, profiler.children)
}
