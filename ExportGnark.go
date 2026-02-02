package gnarktosnarkjs

import (
	"encoding/json"
	"io"

	"github.com/consensys/gnark/backend/witness"
	"github.com/consensys/gnark/frontend/schema"
)

// ExportGnarkVerifyingKey serializes the verifying key to JSON in gnark native
// format (json.MarshalIndent of the groth16.VerifyingKey struct) and writes it to w.
// This format can be used with Garaga and other tools that accept gnark's native JSON.
func ExportGnarkVerifyingKey(vk any, w io.Writer) error {
	return exportJSON(vk, w)
}

// ExportGnarkProof serializes the Groth16 proof to JSON in gnark native format
// (json.MarshalIndent of the groth16.Proof struct) and writes it to w.
// This format can be used with Garaga and other tools that accept gnark's native JSON.
func ExportGnarkProof(proof any, w io.Writer) error {
	return exportJSON(proof, w)
}

// ExportPublicWitness serializes the public part of the witness to JSON and writes it to w.
// schema must be created via frontend.NewSchema(field, circuit).
// Call this when exporting proof to also save public inputs/outputs (e.g. for Garaga).
func ExportPublicWitness(witness witness.Witness, s *schema.Schema, w io.Writer) error {
	data, err := witness.ToJSON(s)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

func exportJSON(v any, w io.Writer) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}
