package tls

import (
	"github.com/openshift/installer/pkg/asset"
	"github.com/pkg/errors"
)

// KeyPairInterface contains a private key and a public key.
type KeyPairInterface interface {
	// Private returns the private key.
	Private() []byte
	// Public returns the public key.
	Public() []byte
}

// KeyPair contains a private key and a public key.
type KeyPair struct {
	private []byte
	public  []byte
	files   []*asset.File
}

// Generate generates the rsa private / public key pair.
func (k *KeyPair) Generate(filenameBase string) error {
	key, err := PrivateKey()
	if err != nil {
		return errors.Wrap(err, "failed to generate private key")
	}

	pubkeyData, err := PublicKeyToPem(&key.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed to get public key data from private key")
	}

	k.private = PrivateKeyToPem(key)
	k.public = pubkeyData

	k.files = []*asset.File{
		{
			Filename: assetFilePath(filenameBase + ".key"),
			Data:     k.private,
		},
		{
			Filename: assetFilePath(filenameBase + ".pub"),
			Data:     k.public,
		},
	}

	return nil
}

// Public returns the public key.
func (k *KeyPair) Public() []byte {
	return k.public
}

// Private returns the private key.
func (k *KeyPair) Private() []byte {
	return k.private
}

// Files returns the files generated by the asset.
func (k *KeyPair) Files() []*asset.File {
	return k.files
}
