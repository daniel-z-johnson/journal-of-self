package models

import (
    "testing"
)

func TestHashPassword(t *testing.T) {
    hash, err := hashPassword("a")
    if err != nil {
        t.Fatal(err)
    }
    t.Log(hash)
}
