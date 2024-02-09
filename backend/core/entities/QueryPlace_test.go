package entities

import (
	"testing"
	"reflect"
)

func TestNewQueryPlace(test *testing.T) {
	test.Run("should be equal to provided struct", func(t *testing.T) {
		qp := NewQueryPlace("somecity")

		if reflect.DeepEqual(qp, QueryPlace{"somecity"}) != true {
			t.FailNow()
		}
	})
}