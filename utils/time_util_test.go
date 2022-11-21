package utils

import (
	"fmt"
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
