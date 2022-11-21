package utils

import (
	"fmt"
	"testing"
)

func TestIsHttpStatusCodeIndicatesClientError(t *testing.T) {
	for i := -1; i <= 7; i++ {
		iFrom := i * 100
		iTo := iFrom + 99
		for _, i2 := range []int{iFrom, iTo} {
			t.Run(fmt.Sprintf("%d", i2), func(t *testing.T) {
				expect := i2 >= 400 && i2 <= 499
				if got := IsHttpStatusCodeIndicatesClientError(i2); got != expect {
					t.Errorf("IsHttpStatusCodeIndicatesClientError() = %v, want %v", got, expect)
				}
			})
		}
	}
}

func TestIsHttpStatusCodeIndicatesInformational(t *testing.T) {
	for i := -1; i <= 7; i++ {
		iFrom := i * 100
		iTo := iFrom + 99
		for _, i2 := range []int{iFrom, iTo} {
			t.Run(fmt.Sprintf("%d", i2), func(t *testing.T) {
				expect := i2 >= 100 && i2 <= 199
				if got := IsHttpStatusCodeIndicatesInformational(i2); got != expect {
					t.Errorf("TestIsHttpStatusCodeIndicatesInformational() = %v, want %v", got, expect)
				}
			})
		}
	}
}

func TestIsHttpStatusCodeIndicatesRedirection(t *testing.T) {
	for i := -1; i <= 7; i++ {
		iFrom := i * 100
		iTo := iFrom + 99
		for _, i2 := range []int{iFrom, iTo} {
			t.Run(fmt.Sprintf("%d", i2), func(t *testing.T) {
				expect := i2 >= 300 && i2 <= 399
				if got := IsHttpStatusCodeIndicatesRedirection(i2); got != expect {
					t.Errorf("TestIsHttpStatusCodeIndicatesRedirection() = %v, want %v", got, expect)
				}
			})
		}
	}
}

func TestIsHttpStatusCodeIndicatesServerError(t *testing.T) {
	for i := -1; i <= 7; i++ {
		iFrom := i * 100
		iTo := iFrom + 99
		for _, i2 := range []int{iFrom, iTo} {
			t.Run(fmt.Sprintf("%d", i2), func(t *testing.T) {
				expect := i2 >= 500 && i2 <= 599
				if got := IsHttpStatusCodeIndicatesServerError(i2); got != expect {
					t.Errorf("TestIsHttpStatusCodeIndicatesServerError() = %v, want %v", got, expect)
				}
			})
		}
	}
}

func TestIsHttpStatusCodeIndicatesSuccess(t *testing.T) {
	for i := -1; i <= 7; i++ {
		iFrom := i * 100
		iTo := iFrom + 99
		for _, i2 := range []int{iFrom, iTo} {
			t.Run(fmt.Sprintf("%d", i2), func(t *testing.T) {
				expect := i2 >= 200 && i2 <= 299
				if got := IsHttpStatusCodeIndicatesSuccess(i2); got != expect {
					t.Errorf("TestIsHttpStatusCodeIndicatesSuccess() = %v, want %v", got, expect)
				}
			})
		}
	}
}
