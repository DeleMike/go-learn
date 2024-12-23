package prose

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{[]string{}, ""},
		testData{[]string{"apple"}, "apple"},
		testData{[]string{"apple", "orange"}, "apple and orange"},
		testData{[]string{"apple", "orange", "pear"}, "apple, orange, and pear"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		want := test.want
		if got != want {
			t.Errorf(errorString(test.list, got, want))
		}
	}
}

func TestNoElement(t *testing.T) {
	var list []string
	want := ""
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}

}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if want != got {
		t.Errorf(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%v) = \"%v\", want \"%v\"", list, got, want)
}
