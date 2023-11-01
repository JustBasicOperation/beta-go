package utils

import (
	"encoding/base64"
	"testing"
)

func TestAesEncryptCBC(t *testing.T) {
	encrypt, err := AesEncryptCBC([]byte("0000000000000000"), []byte("123"), []byte("1234123412341234"))
	if err != nil {
		t.Errorf("err: %v", err)
	}
	t.Logf("encrypt: %v", string(encrypt))
	t.Logf("base64: %v", base64.StdEncoding.EncodeToString(encrypt))
}

func TestAesDecryptCBC(t *testing.T) {
	decodeString, err := base64.StdEncoding.DecodeString("TBiIGu60peezt/M/VRgqaQ==")
	if err != nil {
		t.Errorf("err: %v", err)
	}
	cbc, err := AesDecryptCBC([]byte("0000000000000000"), decodeString, []byte("1234123412341234"))
	if err != nil {
		t.Errorf("err: %v", err)
	}
	t.Logf("cbc: %v", string(cbc))
}

func TestGenAesKey(t *testing.T) {

}
