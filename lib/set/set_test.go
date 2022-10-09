package set

import "testing"

func TestSet(t *testing.T) {
	s := New[string]
	s().Add("aa")
	s().IsExist("aa")
}
