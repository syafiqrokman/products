package access

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/libs/cryptography/pk/encryption"
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
)

type access struct {
	id  *uuid.UUID
	sig signature.PrivateKey
	enc encryption.PrivateKey
}

func createAccess(
	id *uuid.UUID,
	sig signature.PrivateKey,
	enc encryption.PrivateKey,
) Access {
	out := access{
		id:  id,
		sig: sig,
		enc: enc,
	}

	return &out
}

// ID returns the id
func (obj *access) ID() *uuid.UUID {
	return obj.id
}

// Signature returns the signature PK
func (obj *access) Signature() signature.PrivateKey {
	return obj.sig
}

// Encryption returns the encryption PK
func (obj *access) Encryption() encryption.PrivateKey {
	return obj.enc
}
