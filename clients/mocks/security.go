package mocks

import (
	"github.com/mainflux/mainflux/clients"
)

var (
	_ clients.Hasher           = (*hasherMock)(nil)
	_ clients.IdentityProvider = (*identityProviderMock)(nil)
)

// UnauthorizedToken is used to mock Identity method failure.
const UnauthorizedToken = "bd2a557f-27e6-4377-9e40-3c75f3f5211f"

type hasherMock struct{}

func (hm *hasherMock) Hash(pwd string) (string, error) {
	return pwd, nil
}

func (hm *hasherMock) Compare(plain, hashed string) error {
	if plain != hashed {
		return clients.ErrUnauthorizedAccess
	}

	return nil
}

type identityProviderMock struct{}

func (idp *identityProviderMock) TemporaryKey(id string) (string, error) {
	if id == "" || id == UnauthorizedToken {
		return "", clients.ErrUnauthorizedAccess
	}

	return id, nil
}

func (idp *identityProviderMock) PermanentKey(id string) (string, error) {
	return idp.TemporaryKey(id)
}

func (idp *identityProviderMock) Identity(key string) (string, error) {
	return idp.TemporaryKey(key)
}

// NewHasher creates "no-op" hasher for test purposes. This implementation will
// return secrets without changing them.
func NewHasher() clients.Hasher {
	return &hasherMock{}
}

// NewIdentityProvider creates "mirror" identity provider, i.e. generated
// token will hold value provided by the caller.
func NewIdentityProvider() clients.IdentityProvider {
	return &identityProviderMock{}
}
