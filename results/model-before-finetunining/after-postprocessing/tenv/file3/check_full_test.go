package pgpcheck

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com.ProtonMail/gopenpgptest/v2/crypto" // Assuming this is the correct import path
)

func TestCheck_ValidInput_Success(t *testing.T) {
	// Mock data for a valid signature
	data := []byte("valid data")
	dataSig := []byte("valid signature")
	dataPublicKey := []byte("valid public key")

	err := Check(data, dataSig, dataPublicKey)
	assert.Nil(t, err)
}

func TestCheck_InvalidInput_Failure(t *testing.T) error {
	// Mock data for invalid signature
	data := []byte(crypto.NewPlainMessage("valid data").MarshalBinary())
	dataSig := [][]byte{1, 2, 3} // Invalid signature
	dataPublicKey := [][]byte{1,2,3,4,5,6} // Invalid public key

	err := Check([]byte(data), dataSig, dataPublicKey[:len(dataPublicKey)-1]) // Truncated public key
	assert.NotNil(t, err)
	assert.Equal(t, crypto.ErrInvalidSignature, err)
}

// Add more tests for different scenarios...