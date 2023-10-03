package utils

import (
	"fmt"
	"github.com/EscanBE/go-lib/test_utils"
	"testing"
	"time"
)

func TestDiffMs(t *testing.T) {
	for i := 1; i <= 5; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := DiffMs(NowMs())
			if got < 0 || got > 1 {
				t.Errorf("DiffMs() = %v, bad!", got)
			}
		})
	}
}

func TestDiffS(t *testing.T) {
	for i := 1; i <= 5; i++ {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := DiffS(NowS())
			if got < 0 || got > 1 {
				t.Errorf("DiffS() = %v, bad!", got)
			}
		})
	}
}

func TestNowMs(t *testing.T) {
	nowMs1 := NowMs()
	nowMs2 := time.Now().UnixMilli()
	diff := AbsInt64(nowMs1 - nowMs2)
	if diff > 1 {
		t.Errorf("NowMs() = %v, expected %v +-1!", nowMs1, nowMs2)
	}
}

func TestNowS(t *testing.T) {
	nowS1 := NowS()
	nowS2 := time.Now().Unix()
	diff := AbsInt64(nowS1 - nowS2)
	if diff > 1 {
		t.Errorf("NowS() = %v, expected %v +-1!", nowS1, nowS2)
	}
}

func TestGetLocationFromUtcTimezone(t *testing.T) {
	t.Run("get location for UTC from -12 to 14", func(_ *testing.T) {
		for timezone := -12; timezone <= 14; timezone++ {
			loc := GetLocationFromUtcTimezone(timezone)
			nowUTC := time.Date(2023, 9, 16, 0, 0, 0, 0, time.UTC)
			nowWithCustomLoc := nowUTC.In(loc)
			diffHours := (nowWithCustomLoc.Hour() + 24*nowWithCustomLoc.Day()) - (nowUTC.Hour() + 24*nowUTC.Day())
			if diffHours != timezone {
				t.Errorf("Expected diff timezone %d but got %d", timezone, diffHours)
			}
		}
	})

	for timezone := -100; timezone <= 100; timezone++ {
		wantPanic := timezone < -12 || timezone > 14
		t.Run(fmt.Sprintf("timezone %d %s", timezone, func() string {
			if wantPanic {
				return "should panic"
			} else {
				return "should not panic"
			}
		}()), func(_ *testing.T) {
			defer test_utils.DeferWantPanicDepends(t, wantPanic)
			_ = GetLocationFromUtcTimezone(timezone)
		})
	}
}

func TestGetUtcName(t *testing.T) {
	tests := []struct {
		utcTimezone int
		want        string
	}{
		{
			utcTimezone: -12,
			want:        "UTC-1200",
		},
		{
			utcTimezone: -1,
			want:        "UTC-0100",
		},
		{
			utcTimezone: 0,
			want:        "UTC+0000",
		},
		{
			utcTimezone: 1,
			want:        "UTC+0100",
		},
		{
			utcTimezone: 14,
			want:        "UTC+1400",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := GetUtcName(tt.utcTimezone); got != tt.want {
				t.Errorf("GetUtcName() = %v, want %v", got, tt.want)
			}
		})
	}

	for timezone := -100; timezone <= 100; timezone++ {
		wantPanic := timezone < -12 || timezone > 14
		t.Run(fmt.Sprintf("timezone %d %s", timezone, func() string {
			if wantPanic {
				return "should panic"
			} else {
				return "should not panic"
			}
		}()), func(_ *testing.T) {
			defer test_utils.DeferWantPanicDepends(t, wantPanic)
			_ = GetUtcName(timezone)
		})
	}
}
