package core

import (
	"testing"
	"reflect"
)

func TestNewWheatherQuery(test *testing.T) {
	test.Run("should get equal", func(t *testing.T) {
		wq := NewWheatherQuery(81.4, 77.22)

		if reflect.DeepEqual(wq, WheatherQuery{81.4, 77.22}) != true {
			t.FailNow()
		}
	})
}