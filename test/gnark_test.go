package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/mysteryon88/gnark-to-snarkjs/test/cubic"
)

// go test ./test -v -run TestExportGnark_BN254
func TestExportGnark_BN254(t *testing.T) {
	CheckDirs([]string{"cubic/proofs", "cubic/keys"})

	g16 := cubic.G16{}
	g16.Compile(ecc.BN254.ScalarField())
	g16.Setup()
	g16.Prove(ecc.BN254.ScalarField())
	g16.Verify()

	if err := g16.ExportGnark(); err != nil {
		t.Fatalf("ExportGnark: %v", err)
	}

	checkGnarkJSON(t, cubic.GnarkProofPathG16_BN254)
	checkGnarkJSON(t, cubic.GnarkVKeyPathG16_BN254)
}

// go test ./test -v -run TestExportGnark_BLS12_381
func TestExportGnark_BLS12_381(t *testing.T) {
	CheckDirs([]string{"cubic/proofs", "cubic/keys"})

	g16 := cubic.G16{}
	g16.Compile(ecc.BLS12_381.ScalarField())
	g16.Setup()
	g16.Prove(ecc.BLS12_381.ScalarField())
	g16.Verify()

	if err := g16.ExportGnark(); err != nil {
		t.Fatalf("ExportGnark: %v", err)
	}

	checkGnarkJSON(t, cubic.GnarkProofPathG16_BLS12381)
	checkGnarkJSON(t, cubic.GnarkVKeyPathG16_BLS12381)
}

// checkGnarkJSON verifies the file exists and is valid JSON.
func checkGnarkJSON(t *testing.T, path string) {
	t.Helper()
	path = filepath.FromSlash(path)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("read %s: %v", path, err)
		return
	}
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		t.Errorf("invalid JSON in %s: %v", path, err)
		return
	}
	if v == nil {
		t.Errorf("%s: empty JSON", path)
	}
}
