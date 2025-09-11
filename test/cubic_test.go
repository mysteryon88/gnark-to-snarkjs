package test

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/mysteryon88/gnark-to-snarkjs/test/cubic"
)

// go test ./test -v -run TestExportProofAndVK_BN254
func TestExportProofAndVK_BN254(t *testing.T) {
	CheckDirs([]string{"cubic/proofs", "cubic/keys"})

	g16 := cubic.G16{}

	g16.Compile(ecc.BN254.ScalarField())
	g16.Setup()
	g16.Prove(ecc.BN254.ScalarField())
	g16.Verify()
	g16.Export()
}

// go test ./test -v -run TestExportProofAndVK_BLS12_381
func TestExportProofAndVK_BLS12_381(t *testing.T) {
	CheckDirs([]string{"cubic/proofs", "cubic/keys"})

	g16 := cubic.G16{}

	g16.Compile(ecc.BLS12_381.ScalarField())
	g16.Setup()
	g16.Prove(ecc.BLS12_381.ScalarField())
	g16.Verify()
	g16.Export()
}
