package random

import "testing"

// not a very good test, just makes sure there are no errors
// and if you add the -v flag you can verify that the string
// is fairly random
func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(48)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}
