package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {

	var bytes = []byte{23, 34, 45}
	//var bytes2 = []byte{255, 32, 14}

	got, gotError := GetUTFLength(bytes)
	expected := 3
	var expectedError error = nil
	if got != expected && gotError != expectedError {
		t.Fatalf("Ошибка")
	}

	/*got2, gotError2 := GetUTFLength(bytes2)
	expected2 := 0
	var expectedError2 error = errors.New("invalid utf8")
	if got2 != expected2 && gotError2 != expectedError2 {
		t.Fatalf("Ошибка")
	}*/

}
