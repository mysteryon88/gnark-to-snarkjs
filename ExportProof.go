package gnarktosnarkjs

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"

	groth16_bls12381 "github.com/consensys/gnark/backend/groth16/bls12-381"
	groth16_bn254 "github.com/consensys/gnark/backend/groth16/bn254"

	curve_bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	curve_bn254 "github.com/consensys/gnark-crypto/ecc/bn254"
)

// ExportProof serializes a Groth16 proof into a JSON format compatible with snarkjs
// and writes it to the provided writer.
func ExportProof(proof any, publicSignals []string, w io.Writer) error {
	switch p := proof.(type) {
	case *groth16_bls12381.Proof:
		return exportProof_BLS12_381(p, publicSignals, w)
	case *groth16_bn254.Proof:
		return exportProof_BN254(p, publicSignals, w)
	default:
		return fmt.Errorf("unsupported proof type %T (expected *groth16_{bn254,bls12-381}.Proof)", proof)
	}
}

// ---------------- BLS12-381 ----------------
func exportProof_BLS12_381(p *groth16_bls12381.Proof, publicSignals []string, w io.Writer) error {
	if p == nil {
		return fmt.Errorf("proof is nil")
	}
	// G1 -> [x,y,"1"]
	g1 := func(P curve_bls12381.G1Affine) []string {
		return []string{
			P.X.BigInt(new(big.Int)).String(),
			P.Y.BigInt(new(big.Int)).String(),
			"1",
		}
	}
	// G2 -> [[x0,x1],[y0,y1],["1","0"]]
	g2 := func(P curve_bls12381.G2Affine) [][]string {
		return [][]string{
			{P.X.A0.BigInt(new(big.Int)).String(), P.X.A1.BigInt(new(big.Int)).String()},
			{P.Y.A0.BigInt(new(big.Int)).String(), P.Y.A1.BigInt(new(big.Int)).String()},
			{"1", "0"},
		}
	}

	out := map[string]any{
		"protocol": "groth16",
		"curve":    "bls12381",
		"pi_a":     g1(p.Ar),  // A
		"pi_b":     g2(p.Bs),  // B
		"pi_c":     g1(p.Krs), // C
	}

	if len(publicSignals) > 0 {
		out["publicSignals"] = publicSignals
	}

	if len(p.Commitments) > 0 {
		return fmt.Errorf("proof contains commitments, but snarkjs verifier does not support them")
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

// ---------------- BN254 ----------------
func exportProof_BN254(p *groth16_bn254.Proof, publicSignals []string, w io.Writer) error {

	fmt.Println(p.Commitments)
	fmt.Println(p.CommitmentPok)
	if p == nil {
		return fmt.Errorf("proof is nil")
	}
	g1 := func(P curve_bn254.G1Affine) []string {
		return []string{
			P.X.BigInt(new(big.Int)).String(),
			P.Y.BigInt(new(big.Int)).String(),
			"1",
		}
	}
	g2 := func(P curve_bn254.G2Affine) [][]string {
		return [][]string{
			{P.X.A0.BigInt(new(big.Int)).String(), P.X.A1.BigInt(new(big.Int)).String()},
			{P.Y.A0.BigInt(new(big.Int)).String(), P.Y.A1.BigInt(new(big.Int)).String()},
			{"1", "0"},
		}
	}

	out := map[string]any{
		"protocol": "groth16",
		"curve":    "bn254",
		"pi_a":     g1(p.Ar),
		"pi_b":     g2(p.Bs),
		"pi_c":     g1(p.Krs),
	}

	if len(publicSignals) > 0 {
		out["publicSignals"] = publicSignals
	}

	if len(p.Commitments) > 0 {
		return fmt.Errorf("proof contains commitments, but snarkjs verifier does not support them")
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
