package utils

import (
	"testing"
	"time"
)

func TestTruncateDuration(t *testing.T) {
	d := 123456 * time.Microsecond // 123.456ms
	trunc := TruncateDuration(d, time.Millisecond)
	exected := 123 * time.Millisecond
	if trunc != exected {
		t.Fatalf("Invalid truncated duration %v. Expected %v", trunc, exected)
	}
}
