package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5(t *testing.T) {
	b := []byte("1234567890abcdef")
	t.Logf("text: %s, md5: %s", b, Md5(b))
}

var (
	_cipherkey = []byte("1234567890abcdef")
	_plaintext = []byte("text1234")
)

func TestAESEncrypt(t *testing.T) {
	ciphertext := AESEncrypt(_cipherkey, _plaintext)
	t.Logf("ciphertext: %s", ciphertext)
	r, err := AESDecrypt(_cipherkey, ciphertext)
	assert.NoError(t, err)
	assert.Equal(t, _plaintext, r)
}

func TestAESCBCEncrypt(t *testing.T) {
	ciphertext := AESCBCEncrypt(_cipherkey, _plaintext)
	t.Logf("ciphertext: %s", ciphertext)
	r, err := AESCBCDecrypt(_cipherkey, ciphertext)
	assert.NoError(t, err)
	assert.Equal(t, _plaintext, r)
}

func TestAESCTREncrypt(t *testing.T) {
	ciphertext := AESCTREncrypt(_cipherkey, _plaintext)
	t.Logf("ciphertext: %s", ciphertext)
	r, err := AESCTRDecrypt(_cipherkey, ciphertext)
	assert.NoError(t, err)
	assert.Equal(t, _plaintext, r)
}

func TestPading(t *testing.T) {
	blockSize := 16
	padded := hexEncode(pkcs5Padding(_plaintext, blockSize))
	t.Log(string(padded))
	r, err := hexDecode(padded)
	assert.NoError(t, err)
	unpaded, err := pkcs5Unpadding(r)
	assert.NoError(t, err)
	assert.Equal(t, _plaintext, unpaded)
}
