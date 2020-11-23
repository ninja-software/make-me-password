package mmp_test

import (
	"encoding/base64"
	"encoding/hex"
	"testing"

	mmp "github.com/ninja-software/make-me-password"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPasswordBase64(t *testing.T) {
	passwd := []byte("the quick brown fox")

	hsh, err := mmp.HashPassword(string(passwd), "base64")
	if err != nil {
		t.Error(err)
	}

	bhsh, err := base64.StdEncoding.DecodeString(hsh)
	if err != nil {
		t.Error(err)
	}

	err = bcrypt.CompareHashAndPassword(bhsh, passwd)
	if err != nil {
		t.Error(err)
	}
}

func TestHashPasswordHex(t *testing.T) {
	passwd := []byte("the quick brown fox")

	hsh, err := mmp.HashPassword(string(passwd), "hex")
	if err != nil {
		t.Error(err)
	}

	dhsh, err := hex.DecodeString(hsh)
	if err != nil {
		t.Error(err)
	}

	err = bcrypt.CompareHashAndPassword(dhsh, passwd)
	if err != nil {
		t.Error(err)
	}
}
