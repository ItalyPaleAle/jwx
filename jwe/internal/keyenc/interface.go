package keyenc

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"hash"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwe/internal/keygen"
)

// Encrypter is an interface for things that can encrypt keys
type Encrypter interface {
	Algorithm() jwa.KeyEncryptionAlgorithm
	Encrypt([]byte) (keygen.ByteSource, error)
	// KeyID returns the key id for this Encrypter. This exists so that
	// you can pass in a Encrypter to MultiEncrypt, you can rest assured
	// that the generated key will have the proper key ID.
	KeyID() string
}

// Decrypter is an interface for things that can decrypt keys
type Decrypter interface {
	Algorithm() jwa.KeyEncryptionAlgorithm
	Decrypt([]byte) ([]byte, error)
}

type Noop struct {
	alg       jwa.KeyEncryptionAlgorithm
	sharedkey []byte
	keyID     string
}

// AES encrypts content encryption keys using AES key wrap.
// Contrary to what the name implies, it also decrypt encrypted keys
type AES struct {
	alg       jwa.KeyEncryptionAlgorithm
	sharedkey []byte
	keyID     string
}

// AESGCM encrypts content encryption keys using AES-GCM key wrap.
type AESGCMEncrypt struct {
	algorithm jwa.KeyEncryptionAlgorithm
	sharedkey []byte
	keyID     string
}

// ECDHESEncrypt encrypts content encryption keys using ECDH-ES.
type ECDHESEncrypt struct {
	algorithm jwa.KeyEncryptionAlgorithm
	generator keygen.Generator
	keyID     string
}

// ECDHESDecrypt decrypts keys using ECDH-ES.
type ECDHESDecrypt struct {
	keyalg     jwa.KeyEncryptionAlgorithm
	contentalg jwa.ContentEncryptionAlgorithm
	apu        []byte
	apv        []byte
	privkey    *ecdsa.PrivateKey
	pubkey     *ecdsa.PublicKey
}

// RSAOAEPEncrypt encrypts keys using RSA OAEP algorithm
type RSAOAEPEncrypt struct {
	alg    jwa.KeyEncryptionAlgorithm
	pubkey *rsa.PublicKey
	keyID  string
}

// RSAOAEPDecrypt decrypts keys using RSA OAEP algorithm
type RSAOAEPDecrypt struct {
	alg     jwa.KeyEncryptionAlgorithm
	privkey *rsa.PrivateKey
}

// RSAPKCS15Decrypt decrypts keys using RSA PKCS1v15 algorithm
type RSAPKCS15Decrypt struct {
	alg       jwa.KeyEncryptionAlgorithm
	privkey   *rsa.PrivateKey
	generator keygen.Generator
}

// RSAPKCSEncrypt encrypts keys using RSA PKCS1v15 algorithm
type RSAPKCSEncrypt struct {
	alg    jwa.KeyEncryptionAlgorithm
	pubkey *rsa.PublicKey
	keyID  string
}

// DirectDecrypt does no encryption (Note: Unimplemented)
type DirectDecrypt struct {
	Key []byte
}

// PBES2Encrypt encrypts keys with PBES2 / PBKDF2 password
type PBES2Encrypt struct {
	algorithm jwa.KeyEncryptionAlgorithm
	password  []byte
	hashFunc  func() hash.Hash
	keylen    int
	keyID     string
}
