package cubic

import (
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/constraint"
)

const (
	ProofPathG16_BN254 = "cubic/proofs/proof_bn254.json"
	VKeyPathG16_BN254  = "cubic/keys/verification_key_bn254.json"

	ProofPathG16_BLS12381 = "cubic/proofs/proof_bls12381.json"
	VKeyPathG16_BLS12381  = "cubic/keys/verification_key_bls12381.json"
)

type G16 struct {
	circuit Circuit

	r1cs constraint.ConstraintSystem

	pk groth16.ProvingKey
	vk groth16.VerifyingKey

	witnessFull   witness.Witness
	witnessPublic witness.Witness
	proof         groth16.Proof
}
