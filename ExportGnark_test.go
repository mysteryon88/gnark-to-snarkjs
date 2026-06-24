package gnarktosnarkjs

import (
	"bytes"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

type testCircuit struct {
	X frontend.Variable
	Y frontend.Variable `gnark:",public"`
}

func (c *testCircuit) Define(api frontend.API) error {
	x3 := api.Mul(c.X, c.X, c.X)
	api.AssertIsEqual(api.Add(x3, c.X, 5), c.Y)
	return nil
}

func TestExportGnarkBinaryRoundTripBN254(t *testing.T) {
	const curveID = ecc.BN254
	field := curveID.ScalarField()

	ccs, err := frontend.Compile(field, r1cs.NewBuilder, &testCircuit{})
	if err != nil {
		t.Fatalf("compile circuit: %v", err)
	}
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		t.Fatalf("setup: %v", err)
	}

	assignment := testCircuit{X: 3, Y: 35}
	fullWitness, err := frontend.NewWitness(&assignment, field)
	if err != nil {
		t.Fatalf("full witness: %v", err)
	}
	publicWitness, err := fullWitness.Public()
	if err != nil {
		t.Fatalf("public witness: %v", err)
	}
	proof, err := groth16.Prove(ccs, pk, fullWitness)
	if err != nil {
		t.Fatalf("prove: %v", err)
	}

	var proofOut bytes.Buffer
	if err := ExportGnarkProofBinary(proof, &proofOut); err != nil {
		t.Fatalf("export proof binary: %v", err)
	}
	restoredProof := groth16.NewProof(curveID)
	if _, err := restoredProof.ReadFrom(bytes.NewReader(proofOut.Bytes())); err != nil {
		t.Fatalf("read proof binary: %v", err)
	}

	var vkOut bytes.Buffer
	if err := ExportGnarkVerifyingKeyBinary(vk, &vkOut); err != nil {
		t.Fatalf("export verifying key binary: %v", err)
	}
	restoredVK := groth16.NewVerifyingKey(curveID)
	if _, err := restoredVK.ReadFrom(bytes.NewReader(vkOut.Bytes())); err != nil {
		t.Fatalf("read verifying key binary: %v", err)
	}

	var publicWitnessOut bytes.Buffer
	if err := ExportPublicWitnessBinary(publicWitness, &publicWitnessOut); err != nil {
		t.Fatalf("export public witness binary: %v", err)
	}
	restoredPublicWitness, err := witness.New(field)
	if err != nil {
		t.Fatalf("new public witness: %v", err)
	}
	if _, err := restoredPublicWitness.ReadFrom(bytes.NewReader(publicWitnessOut.Bytes())); err != nil {
		t.Fatalf("read public witness binary: %v", err)
	}

	if err := groth16.Verify(restoredProof, restoredVK, restoredPublicWitness); err != nil {
		t.Fatalf("verify restored proof/vk/witness: %v", err)
	}
}
