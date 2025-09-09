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

// ExportVerifyingKey serializes the verifying key into a JSON format compatible with snarkjs
// and writes it to the provided writer.
func ExportVerifyingKey(vk any, w io.Writer) error {
	switch t := vk.(type) {
	case *groth16_bls12381.VerifyingKey:
		return exportVK_BLS12_381(t, w)
	case *groth16_bn254.VerifyingKey:
		return exportVK_BN254(t, w)
	default:
		return fmt.Errorf("unsupported VK type %T (only bn254 & bls12-381)", vk)
	}
}

// ---------------- BLS12-381 ----------------

func exportVK_BLS12_381(vk *groth16_bls12381.VerifyingKey, w io.Writer) error {
	if vk == nil {
		return fmt.Errorf("verifying key is nil")
	}
	g1 := func(p curve_bls12381.G1Affine) []string {
		return []string{
			p.X.BigInt(new(big.Int)).String(),
			p.Y.BigInt(new(big.Int)).String(),
			"1",
		}
	}
	g2 := func(p curve_bls12381.G2Affine) [][]string {
		return [][]string{
			{p.X.A0.BigInt(new(big.Int)).String(), p.X.A1.BigInt(new(big.Int)).String()},
			{p.Y.A0.BigInt(new(big.Int)).String(), p.Y.A1.BigInt(new(big.Int)).String()},
			{"1", "0"},
		}
	}
	ab, err := curve_bls12381.Pair(
		[]curve_bls12381.G1Affine{vk.G1.Alpha},
		[]curve_bls12381.G2Affine{vk.G2.Beta},
	)
	if err != nil {
		return fmt.Errorf("pairing(alpha,beta) failed: %w", err)
	}
	gt := [][][]string{
		{
			{ab.C0.B0.A0.BigInt(new(big.Int)).String(), ab.C0.B0.A1.BigInt(new(big.Int)).String()},
			{ab.C0.B1.A0.BigInt(new(big.Int)).String(), ab.C0.B1.A1.BigInt(new(big.Int)).String()},
			{ab.C0.B2.A0.BigInt(new(big.Int)).String(), ab.C0.B2.A1.BigInt(new(big.Int)).String()},
		},
		{
			{ab.C1.B0.A0.BigInt(new(big.Int)).String(), ab.C1.B0.A1.BigInt(new(big.Int)).String()},
			{ab.C1.B1.A0.BigInt(new(big.Int)).String(), ab.C1.B1.A1.BigInt(new(big.Int)).String()},
			{ab.C1.B2.A0.BigInt(new(big.Int)).String(), ab.C1.B2.A1.BigInt(new(big.Int)).String()},
		},
	}
	var ic [][]string
	for _, p := range vk.G1.K {
		ic = append(ic, g1(p))
	}
	out := struct {
		Protocol    string       `json:"protocol"`
		Curve       string       `json:"curve"`
		NPublic     int          `json:"nPublic"`
		VKAlpha1    []string     `json:"vk_alpha_1"`
		VKBeta2     [][]string   `json:"vk_beta_2"`
		VKGamma2    [][]string   `json:"vk_gamma_2"`
		VKDelta2    [][]string   `json:"vk_delta_2"`
		VKAlphaBeta [][][]string `json:"vk_alphabeta_12,omitempty"`
		IC          [][]string   `json:"IC"`
	}{
		Protocol:    "groth16",
		Curve:       "bls12381",
		NPublic:     vk.NbPublicWitness(),
		VKAlpha1:    g1(vk.G1.Alpha),
		VKBeta2:     g2(vk.G2.Beta),
		VKGamma2:    g2(vk.G2.Gamma),
		VKDelta2:    g2(vk.G2.Delta),
		VKAlphaBeta: gt,
		IC:          ic,
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

// ---------------- BN254 ----------------

func exportVK_BN254(vk *groth16_bn254.VerifyingKey, w io.Writer) error {
	if vk == nil {
		return fmt.Errorf("verifying key is nil")
	}
	g1 := func(p curve_bn254.G1Affine) []string {
		return []string{
			p.X.BigInt(new(big.Int)).String(),
			p.Y.BigInt(new(big.Int)).String(),
			"1",
		}
	}
	g2 := func(p curve_bn254.G2Affine) [][]string {
		return [][]string{
			{p.X.A0.BigInt(new(big.Int)).String(), p.X.A1.BigInt(new(big.Int)).String()},
			{p.Y.A0.BigInt(new(big.Int)).String(), p.Y.A1.BigInt(new(big.Int)).String()},
			{"1", "0"},
		}
	}
	ab, err := curve_bn254.Pair(
		[]curve_bn254.G1Affine{vk.G1.Alpha},
		[]curve_bn254.G2Affine{vk.G2.Beta},
	)
	if err != nil {
		return fmt.Errorf("pairing(alpha,beta) failed: %w", err)
	}
	gt := [][][]string{
		{
			{ab.C0.B0.A0.BigInt(new(big.Int)).String(), ab.C0.B0.A1.BigInt(new(big.Int)).String()},
			{ab.C0.B1.A0.BigInt(new(big.Int)).String(), ab.C0.B1.A1.BigInt(new(big.Int)).String()},
			{ab.C0.B2.A0.BigInt(new(big.Int)).String(), ab.C0.B2.A1.BigInt(new(big.Int)).String()},
		},
		{
			{ab.C1.B0.A0.BigInt(new(big.Int)).String(), ab.C1.B0.A1.BigInt(new(big.Int)).String()},
			{ab.C1.B1.A0.BigInt(new(big.Int)).String(), ab.C1.B1.A1.BigInt(new(big.Int)).String()},
			{ab.C1.B2.A0.BigInt(new(big.Int)).String(), ab.C1.B2.A1.BigInt(new(big.Int)).String()},
		},
	}
	var ic [][]string
	for _, p := range vk.G1.K {
		ic = append(ic, g1(p))
	}
	out := struct {
		Protocol    string       `json:"protocol"`
		Curve       string       `json:"curve"`
		NPublic     int          `json:"nPublic"`
		VKAlpha1    []string     `json:"vk_alpha_1"`
		VKBeta2     [][]string   `json:"vk_beta_2"`
		VKGamma2    [][]string   `json:"vk_gamma_2"`
		VKDelta2    [][]string   `json:"vk_delta_2"`
		VKAlphaBeta [][][]string `json:"vk_alphabeta_12,omitempty"`
		IC          [][]string   `json:"IC"`
	}{
		Protocol:    "groth16",
		Curve:       "bn254",
		NPublic:     vk.NbPublicWitness(),
		VKAlpha1:    g1(vk.G1.Alpha),
		VKBeta2:     g2(vk.G2.Beta),
		VKGamma2:    g2(vk.G2.Gamma),
		VKDelta2:    g2(vk.G2.Delta),
		VKAlphaBeta: gt,
		IC:          ic,
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}
