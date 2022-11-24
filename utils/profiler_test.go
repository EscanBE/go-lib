package utils

import (
	"fmt"
	"github.com/EscanBE/go-lib/test_utils"
	"github.com/stretchr/testify/require"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestNewMasterProfiler(t *testing.T) {
	require.Nil(t, NewMasterProfiler("", "", false))

	desc := fmt.Sprintf("desc %d", rand.Int())
	constDesc := fmt.Sprintf("constDesc %d", rand.Int())
	profiler := NewMasterProfiler(desc, constDesc, true)
	now := time.Now().UnixMilli()
	require.NotNil(t, profiler)
	require.Equal(t, constDesc, profiler.constDesc)
	require.True(t, math.Abs(float64(now-profiler.start)) < 10)
	require.Zero(t, profiler.duration)
	require.False(t, profiler.finalized)
	require.Zero(t, profiler.level)
	require.False(t, profiler.err)
	require.Nil(t, profiler.children)
}

func Test_newProfiler(t *testing.T) {
	desc := fmt.Sprintf("desc %d", rand.Int())
	level := rand.Int()
	profiler := newProfiler(desc, level)
	now := time.Now().UnixMilli()
	require.NotNil(t, profiler)
	require.Empty(t, profiler.constDesc)
	require.True(t, math.Abs(float64(now-profiler.start)) < 10)
	require.Zero(t, profiler.duration)
	require.False(t, profiler.finalized)
	require.Equal(t, level, profiler.level)
	require.False(t, profiler.err)
	require.Nil(t, profiler.children)
}

func TestProfiler_NewChild(t *testing.T) {
	require.Nil(t, (*Profiler)(nil).NewChild(""), "nil should creates nil")
	parentDesc := fmt.Sprintf("desc %d", rand.Int())
	parentLevel := rand.Int()
	parentProfiler := newProfiler(parentDesc, parentLevel)
	parentProfiler.constDesc = fmt.Sprintf("constDesc %d", rand.Int())

	childDesc := fmt.Sprintf("desc %d", rand.Int())
	profiler := parentProfiler.NewChild(childDesc)
	now := time.Now().UnixMilli()

	require.NotNil(t, profiler)
	require.Equal(t, profiler, parentProfiler.children[0], "must be appended as parent children")
	require.Equal(t, parentProfiler.constDesc, profiler.constDesc, "constDesc should be coped from parent")
	require.True(t, math.Abs(float64(now-profiler.start)) < 10)
	require.Zero(t, profiler.duration)
	require.False(t, profiler.finalized)
	require.Equal(t, parentLevel+1, profiler.level)
	require.False(t, profiler.err)
	require.Nil(t, profiler.children)

	_ = parentProfiler.NewChild("")
	_ = profiler.NewChild("")
	require.Equal(t, 2, len(parentProfiler.children))
	require.Equal(t, 1, len(profiler.children))
}

func TestProfiler_Finalize(t *testing.T) {
	nilProfiler := (*Profiler)(nil)
	require.Nil(t, nilProfiler.Finalize())

	desc := test_utils.RadStr(6)
	level := rand.Int()
	profiler := newProfiler(desc, level)

	// make sure before finalize
	require.NotZero(t, profiler.start)
	require.Zero(t, profiler.duration)
	require.False(t, profiler.err)
	require.False(t, profiler.finalized)

	// backup data
	start := profiler.start

	// 1st finalize
	time.Sleep(10 * time.Millisecond)
	require.Equal(t, profiler, profiler.Finalize())

	// expect fields doesn't change
	require.Equal(t, desc, profiler.desc)
	require.Equal(t, level, profiler.level)
	require.Equal(t, start, profiler.start)
	require.False(t, profiler.err, "by calling Finalize(), err should not be changed")

	// expect fields changed
	require.True(t, profiler.finalized)
	require.NotZero(t, profiler.duration)

	// backup data
	duration := profiler.duration

	// 2nd finalize
	time.Sleep(10 * time.Millisecond)
	require.Equal(t, profiler, profiler.Finalize())

	// expect fields not changed after 2nd finalize
	require.True(t, profiler.finalized)
	require.Equal(t, duration, profiler.duration)
	require.False(t, profiler.err)
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
			require.NotZero(t, profiler.start)
			require.Zero(t, profiler.duration)
			require.False(t, profiler.err)
			require.False(t, profiler.finalized)

			// backup data
			start := profiler.start

			// 1st finalize
			time.Sleep(10 * time.Millisecond)
			profiler.FinalizeWithCheckErr(tt.err)

			// expect fields doesn't change
			require.Equal(t, desc, profiler.desc)
			require.Equal(t, level, profiler.level)
			require.Equal(t, start, profiler.start)

			// expect fields changed
			require.True(t, profiler.finalized)
			require.NotZero(t, profiler.duration)
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
			require.True(t, profiler.finalized)
			require.Equal(t, duration, profiler.duration)
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
			require.True(t, profiler.err, "must be true")
		})
	}
}

func TestProfiler_FinalizeWithErr(t *testing.T) {
	var err error
	nilProfiler1 := (*Profiler)(nil)
	tmpErr := fmt.Errorf("err")
	err = nilProfiler1.FinalizeWithErr(tmpErr) // shouldn't panic
	require.Equal(t, tmpErr, err)
	nilProfiler2 := (*Profiler)(nil)
	err = nilProfiler2.FinalizeWithErr(nil) // shouldn't panic either
	require.Nil(t, err)

	desc := test_utils.RadStr(6)
	level := rand.Int()
	profiler := newProfiler(desc, level)

	// make sure before finalize
	require.NotZero(t, profiler.start)
	require.Zero(t, profiler.duration)
	require.False(t, profiler.err)
	require.False(t, profiler.finalized)

	// backup data
	start := profiler.start

	// 1st finalize
	time.Sleep(10 * time.Millisecond)
	_ = profiler.FinalizeWithErr(fmt.Errorf("err1")) // shouldn't panic

	// expect fields doesn't change
	require.Equal(t, desc, profiler.desc)
	require.Equal(t, level, profiler.level)
	require.Equal(t, start, profiler.start)

	// expect fields changed
	require.True(t, profiler.finalized)
	require.NotZero(t, profiler.duration)
	require.True(t, profiler.err)

	// backup data
	duration := profiler.duration

	// 2nd finalize
	time.Sleep(10 * time.Millisecond)
	_ = profiler.FinalizeWithErr(fmt.Errorf("err2"))

	// expect fields not changed after 2nd finalize
	require.True(t, profiler.finalized)
	require.Equal(t, duration, profiler.duration)
	require.True(t, profiler.err)

	// 3rd finalize with non-err
	defer test_utils.DeferWantPanic(t)
	_ = profiler.FinalizeWithErr(nil)
}

func randomGenerateChildren(parent *Profiler, noErr bool) (generatedAnyErr bool) {
	if parent == nil {
		return false
	}
	numberOfChild := rand.Int() % 3
	if numberOfChild < 1 {
		return false
	}
	var anyErr bool
	for i := 0; i < numberOfChild; i++ {
		child := parent.NewChild("")
		if !noErr && rand.Int()%100 < 20 {
			_ = child.FinalizeWithErr(fmt.Errorf("err"))
			anyErr = true
		}

		if randomGenerateChildren(child, noErr) {
			anyErr = true
		}
	}
	return anyErr
}

func TestProfiler_Print(t *testing.T) {
	require.False(t, NewMasterProfiler("", "", false).Print())

	const maxTest = 10
	for i := 1; i <= maxTest; i++ {
		profiler := NewMasterProfiler("", "", true)
		require.NotNil(t, profiler)
		noErrShouldBeGenerated := i == 1 // make sure at least one profiler has no err
		anyErr := randomGenerateChildren(profiler, noErrShouldBeGenerated)
		if i == maxTest {
			_ = profiler.FinalizeWithErr(fmt.Errorf("err")) // make sure at least one profiler has error
			anyErr = true
		}
		got := profiler.Print()
		require.Truef(t, anyErr == got, "got %v (i=%d), want %v", got, i, anyErr)
	}
}
