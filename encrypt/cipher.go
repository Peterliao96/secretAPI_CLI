package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

func Encrypt(key,plaintext string) (string, error){
  block, err := newCipherBlock(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}


	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	return fmt.Sprintf("%x", ciphertext), nil
}

func Decrypt(key, cipherHex string)(string, error){
  block, err := newCipherBlock(key)
  if err != nil {
    return "", err
  }

  ciphertext, err := hex.DecodeString(cipherHex)
  if err != nil {
    return "", err
  }

  if len(ciphertext)< aes.BlockSize{
    return "", errors.New("encrypt: cipher too short")
  }

  iv := ciphertext[:aes.BlockSize]
  ciphertext = ciphertext[aes.BlockSize:]

  // similar to salting password
  // attach initialization of vector with the key value
  stream := cipher.NewCFBDecrypter(block, iv)
  stream.XORKeyStream(ciphertext, ciphertext)

  return fmt.Sprintf("%s",ciphertext),nil
}

func newCipherBlock(key string) (cipher.Block, error){
  hasher := md5.New()
  fmt.Fprint(hasher, key)
  cipherKey := hasher.Sum(nil)
  fmt.Println(len(cipherKey))
  return aes.NewCipher(cipherKey)
}
