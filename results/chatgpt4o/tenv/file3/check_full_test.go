package pgpcheck_test

import (
	"testing"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/your_project/pgpcheck"
)

func TestCheck_ValidSignature(t *testing.T) {
	data := []byte("Hello, World!")
	dataSig := []byte("VALID_SIGNATURE")
	dataPublicKey := []byte("VALID_PUBLIC_KEY")

	err := pgpcheck.Check(data, dataSig, dataPublicKey)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCheck_InvalidSignature(t *testing.T) {
	data := []byte("Hello, World!")
	dataSig := []byte("INVALID_SIGNATURE")
	dataPublicKey := []byte("VALID_PUBLIC_KEY")

	err := pgpcheck.Check(data, dataSig, dataPublicKey)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestCheck_InvalidPublicKey(t *testing.T) {
	data := []byte("Hello, World!")
	dataSig := []byte("VALID_SIGNATURE")
	dataPublicKey := []byte("INVALID_PUBLIC_KEY")

	err := pgpcheck.Check(data, dataSig, dataPublicKey)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestCheck_EmptyData(t *testing.T) {
	data := []byte("")
	dataSig := []byte("VALID_SIGNATURE")
	dataPublicKey := []byte("VALID_PUBLIC_KEY")

	err := pgpcheck.Check(data, dataSig, dataPublicKey)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCheck_EmptySignature(t *testing.T) {
	data := []byte("Hello, World!")
	dataSig := []byte("")
	dataPublicKey := []byte("VALID_PUBLIC_KEY")

	err := pgpcheck.Check(data, dataSig, dataPublicKey)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestCheck_EmptyPublicKey(t *testing.T) {
	data := []byte("Hello, World!")
	dataSig := []byte("VALID_SIGNATURE")
	dataPublicKey := []byte("")

	err := pgpcheck.Check(data, dataSig, dataPublicKey)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestCheck_NilData(t *testing.T) {
	var data []byte = nil
	dataSig := []byte("VALID_SIGNATURE")
	dataPublicKey := []byte("VALID_PUBLIC_KEY")

	err := pgpcheck.Check(data, dataSig, dataPublicKey)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestCheck_NilSignature(t *testing.T) {
	data := []byte("Hello, World!")
	var dataSig []byte = nil
	dataPublicKey := []byte("VALID_PUBLIC_KEY")

	err := pgpcheck.Check(data, dataSig, dataPublicKey)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestCheck_NilPublicKey(t *testing.T) {
	data := []byte("Hello, World!")
	dataSig := []byte("VALID_SIGNATURE")
	var dataPublicKey []byte = nil

	err := pgpcheck.Check(data, dataSig, dataPublicKey)
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestCheck_MismatchedSignature(t *testing.T) {
	data := []byte("Hello, World!")
	dataSig := []byte("MISMATCHED_SIGNATURE")
	dataPublicKey := []byte("VALID_PUBLIC_KEY")

	err := pgpcheck.Check(data, dataSig, dataPublicKey)
	if err == nil {
		t.Error("expected error, got nil")
	}
}