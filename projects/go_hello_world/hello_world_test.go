package go_hello_world

import "testing"

func Test_HelloWorld(t *testing.T) {
	expected := "Hello World!"
	actual := HelloWorld()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
