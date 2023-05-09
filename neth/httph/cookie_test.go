package httph

import (
	"eval_helper/timeh"
	"testing"
)

func TestCookieDeleteTime(t *testing.T) {
	if r := CookieDeleteTime().UnixNano(); r != 1*timeh.NanosecsInSec {
		t.Errorf("expect %v, got %v", timeh.NanosecsInSec, r)
	}
}
