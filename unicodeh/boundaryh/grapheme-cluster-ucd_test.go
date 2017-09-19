package boundaryh

import (
	"testing"
	"reflect"
)

type ucdGraphemeClusterTest ucdTest

func TestGraphemeClusters(t *testing.T) {
	for testI, test := range ucdGraphemeClusterTests {
		boundaries := GraphemeClusters(test.runes)
		if !reflect.DeepEqual(test.boundaries, boundaries) {
			t.Errorf("%v \"%v\": expect %v, got %v", testI, (test.runes), test.boundaries, boundaries)
		}
	}
}