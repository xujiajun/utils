package math2

import (
	"testing"
)

func TestMinInt(t *testing.T)  {
	bit := 32 << (^uint(0) >> 63)
	if bit == 32 {
		if MinInt != 1<<31 - 1 {
			t.Error("err bit 32 MinInt")
		}
	}
}