package random

import "testing"

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(48)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}
