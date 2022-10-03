package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
	result := Hello("Higaharu")
	want := "Hi, Higaharu. Welcome!"
	if result != want {
		t.Errorf("fileWrite.Hello() = %q want %q", result, want)
	}
}
