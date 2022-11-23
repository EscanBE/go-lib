package utils

import (
	"fmt"
	"github.com/EscanBE/go-lib/test_utils"
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

func TestProfiler_NewChild(t *testing.T) {
	assert.Nil(t, (*Profiler)(nil).NewChild(""), "nil should creates nil")
	parentDesc := fmt.Sprintf("desc %d", rand.Int())
	parentLevel := rand.Int()
	parentProfiler := newProfiler(parentDesc, parentLevel)
	parentProfiler.constDesc = fmt.Sprintf("constDesc %d", rand.Int())

	childDesc := fmt.Sprintf("desc %d", rand.Int())
	profiler := parentProfiler.NewChild(childDesc)
	now := time.Now().UnixMilli()

	assert.NotNil(t, profiler)
	assert.Equal(t, profiler, parentProfiler.children[0], "must be appended as parent children")
	assert.Equal(t, parentProfiler.constDesc, profiler.constDesc, "constDesc should be coped from parent")
	assert.True(t, math.Abs(float64(now-profiler.start)) < 10)
	assert.Zero(t, profiler.duration)
	assert.False(t, profiler.finalized)
	assert.Equal(t, parentLevel+1, profiler.level)
	assert.False(t, profiler.err)
	assert.Nil(t, profiler.children)

	_ = parentProfiler.NewChild("")
	_ = profiler.NewChild("")
	assert.Equal(t, 2, len(parentProfiler.children))
	assert.Equal(t, 1, len(profiler.children))
}

func TestProfiler_Finalize(t *testing.T) {
	nilProfiler := (*Profiler)(nil)
	assert.Nil(t, nilProfiler.Finalize())

	desc := test_utils.RadStr(6)
	level := rand.Int()
	profiler := newProfiler(desc, level)

	// make sure before finalize
	assert.NotZero(t, profiler.start)
	assert.Zero(t, profiler.duration)
	assert.False(t, profiler.err)
	assert.False(t, profiler.finalized)

	// backup data
	start := profiler.start

	// 1st finalize
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, profiler, profiler.Finalize())

	// expect fields doesn't change
	assert.Equal(t, desc, profiler.desc)
	assert.Equal(t, level, profiler.level)
	assert.Equal(t, start, profiler.start)
	assert.False(t, profiler.err, "by calling Finalize(), err should not be changed")

	// expect fields changed
	assert.True(t, profiler.finalized)
	assert.NotZero(t, profiler.duration)

	// backup data
	duration := profiler.duration

	// 2nd finalize
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, profiler, profiler.Finalize())

	// expect fields not changed after 2nd finalize
	assert.True(t, profiler.finalized)
	assert.Equal(t, duration, profiler.duration)
	assert.False(t, profiler.err)
}

func TestProfiler_FinalizeWithCheckErr(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "without error",
			err:  nil,
		},
		{
			name: "with error",
			err:  fmt.Errorf("fake"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nilProfiler := (*Profiler)(nil)
			nilProfiler.FinalizeWithCheckErr(tt.err) // shouldn't panic

			desc := test_utils.RadStr(6)
			level := rand.Int()
			profiler := newProfiler(desc, level)

			// make sure before finalize
			assert.NotZero(t, profiler.start)
			assert.Zero(t, profiler.duration)
			assert.False(t, profiler.err)
			assert.False(t, profiler.finalized)

			// backup data
			start := profiler.start

			// 1st finalize
			time.Sleep(10 * time.Millisecond)
			profiler.FinalizeWithCheckErr(tt.err)

			// expect fields doesn't change
			assert.Equal(t, desc, profiler.desc)
			assert.Equal(t, level, profiler.level)
			assert.Equal(t, start, profiler.start)

			// expect fields changed
			assert.True(t, profiler.finalized)
			assert.NotZero(t, profiler.duration)
			if (tt.err != nil) != profiler.err {
				t.Errorf("err %v, want %v", tt.err != nil, profiler.err)
			}

			// backup data
			duration := profiler.duration
			err := profiler.err

			// 2nd finalize
			time.Sleep(10 * time.Millisecond)
			profiler.FinalizeWithCheckErr(tt.err)

			// expect fields not changed after 2nd finalize
			assert.True(t, profiler.finalized)
			assert.Equal(t, duration, profiler.duration)
			if err != profiler.err {
				t.Errorf("err %v, want %v", err, profiler.err)
			}

			// 3rd finalize with opposite err
			if tt.err != nil {
				tt.err = nil
			} else {
				tt.err = fmt.Errorf("error")
			}
			time.Sleep(10 * time.Millisecond)
			profiler.FinalizeWithCheckErr(tt.err)
			// if previous finalize not err, then override with err => err
			// if previous finalize with err, then ignore any other update state to err
			assert.True(t, profiler.err, "must be true")
		})
	}
}

func TestProfiler_FinalizeWithErr(t *testing.T) {
	nilProfiler1 := (*Profiler)(nil)
	nilProfiler1.FinalizeWithErr(fmt.Errorf("err")) // shouldn't panic
	nilProfiler2 := (*Profiler)(nil)
	nilProfiler2.FinalizeWithErr(nil) // shouldn't panic either

	desc := test_utils.RadStr(6)
	level := rand.Int()
	profiler := newProfiler(desc, level)

	// make sure before finalize
	assert.NotZero(t, profiler.start)
	assert.Zero(t, profiler.duration)
	assert.False(t, profiler.err)
	assert.False(t, profiler.finalized)

	// backup data
	start := profiler.start

	// 1st finalize
	time.Sleep(10 * time.Millisecond)
	profiler.FinalizeWithErr(fmt.Errorf("err1")) // shouldn't panic

	// expect fields doesn't change
	assert.Equal(t, desc, profiler.desc)
	assert.Equal(t, level, profiler.level)
	assert.Equal(t, start, profiler.start)

	// expect fields changed
	assert.True(t, profiler.finalized)
	assert.NotZero(t, profiler.duration)
	assert.True(t, profiler.err)

	// backup data
	duration := profiler.duration

	// 2nd finalize
	time.Sleep(10 * time.Millisecond)
	profiler.FinalizeWithErr(fmt.Errorf("err2"))

	// expect fields not changed after 2nd finalize
	assert.True(t, profiler.finalized)
	assert.Equal(t, duration, profiler.duration)
	assert.True(t, profiler.err)

	// 3rd finalize with non-err
	defer test_utils.DeferWantPanic(t)
	profiler.FinalizeWithErr(nil)
}
