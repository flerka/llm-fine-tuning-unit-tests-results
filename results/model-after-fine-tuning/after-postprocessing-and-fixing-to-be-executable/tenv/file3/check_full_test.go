package pgpcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

func TestCheck_Success(t *testing.T) {
	data := []byte("test data")
	dataSig := []byte("test signature")
	dataPublicKey := []byte("test public key")

	err := Check(data, dataSig, dataPublicKey)
	require.NoError(t, err)
}

func TestCheck_Failure(t *testing.T) {
	data := []byte("test data")

	err := Check(data, nil, nil)
	require.Error(t, err)
}

func TestCheck_InvalidPublicKey(t *testing.T) {
	data := []byte("test data")
	dataSig := []byte("test signature")

	err := Check(data, dataSig, []byte("invalid public key"))
	require.Error(t, err)
}

func TestCheck_InvalidSignature(t *testing.T) {
	data := []byte("test data")
	dataSig := []byte("invalid signature")

	err := Check(data, dataSig, []byte("test public key"))
	require.Error(t, err)
}

func TestVerifyDetached_Success(t *testing.T) {
	data := []byte("test data 1")
	dataSig := []byte("test signature 1")
	dataPublicKey := []byte("test public key 1")

	err := Check(data, dataSig, dataPublicKey)

	require.NoError(t, err)

	message := crypto.NewPlainMessage(data)
	key, _ := crypto.NewKeyFromArmored(string(dataPublicKey))
	signingKeyRing, err := crypto.NewKeyRing(key)
	require.NoError(t, err)

	assert.NoError(t, signingKeyRing.VerifyDetached(message, crypto.NewPGPSignature(dataSig), crypto.GetUnixTime()))
}

func TestVerifyDetached_Failure(t *testing.T) {
	data := []byte("test data 10")
	dataSig := []byte("test signature 10")
	dataPublicKey := []byte("test public key 10")

	err := Check(data, dataSig, dataPublicKey)
	require.Error(t, err)

	message := crypto.NewPlainMessage([]byte("test data 1"))
	key, _ := crypto.NewKeyFromArmored(string(dataPublicKey))
	signingKeyRing, err := crypto.NewKeyRing(key)
	require.NoError(t, err)

	assert.Error(t, signingKeyRing.VerifyDetached(message, crypto.NewPGPSignature(dataSig), crypto.GetUnixTime()))
}

func BenchmarkCheck(b *testing.B) {
	data := []byte("test data")
	dataSig := []uint8{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	dataPublicKey := []byte("test public key")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := Check(data, dataSig, dataPublicKey)
		require.NoError(b, err)
	}
}

func BenchmarkVerifyDetached(b *testing.B) {
	data := []byte("test data 1")
	dataSig1 := []uint8{0x00, 0x01, 0x12, 0x03, 0x04, 0x05}
	dataSig2 := []uint8{0x00, 0x01, 0x23, 0x03, 0x04, 0x05} // invalid signature
	dataPublicKey := []byte("test public key 1")

	b.ResetTimer()
	for i := 0; i < 1000; i++ {
		err := Check(data, dataSig1, dataPublicKey)
		require.NoError(b, err)
		err = Check(data, dataSig2, dataPublicKey)
		require.Error(b, err)
	}
}

func BenchmarkVerifyDetach_Success(b *testing.B) {
	data := []byte("test data 10")
	dataSigs := [][]uint8{
		{0x00, 0x01, 0x12, 1, 0x04, 0x05},
		{0x00, 0x01, 0x23, 1, 0x04, 0x05},
	}
	dataPublicKey := []byte("test public key 10")

	b.ResetTimer()
	for i := 0; i < len(dataSigs); i++ {
		err := Check(data, dataSigs[i], dataPublicKey)
		require.NoError(b, err)
	}
}

func BenchmarkVerifySignature(b *testing.B) {
	data := []byte("test data 12")
	dataSig := []uint8{0x00, 0x12, 0x03, 0x04,
		0x05, 0x06, 0x07, 0xa0, 0x08, 0x09, 0x0a,
		0x0b, 0x0c, 0x0d, 0xe0, 0x0e, 0x0f}
	dataPublicKey := make([]byte, 1000)

	b.ResetTimer()
	for i := 0; i < 500; i++ {
		err := Check(data, dataSig, dataPublicKey)
		require.NoError(b, err)
	}
}
func BenchmarkVerifySignature_Failure(b *testing.B) {
	data := []byte("test data")

	dataSig := []uint8{0x00, 0x10, 0x03, 0x04,
		0x00, 0x05, 0x06, 0x07,
		0x08, 0x09, 0x0a, 0xf0,
		0x0e, 0x0f}
	dataPublicKey := make([]byte, 1000)

	b.ResetTimer()
	for {
		err := Check(data, dataSig, dataPublicKey)

		if err != nil {
			break
		}
	}
}
