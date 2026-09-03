package gnarktosnarkjs

import (
	"encoding/json"
	"fmt"
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
	publicWitness, err := witness.Public()
	if err != nil {
		return err
	}
	data, err := publicWitness.ToJSON(s)
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}

// ExportGnarkVerifyingKeyBinary serializes the verifying key to gnark's native
// binary format via io.WriterTo and writes it to w.
func ExportGnarkVerifyingKeyBinary(vk any, w io.Writer) error {
	return exportBinary("verifying key", vk, w)
}

// ExportGnarkProofBinary serializes the proof to gnark's native binary format
// via io.WriterTo and writes it to w.
func ExportGnarkProofBinary(proof any, w io.Writer) error {
	return exportBinary("proof", proof, w)
}

// ExportPublicWitnessBinary serializes the witness to gnark's native binary
// format and writes it to w.
func ExportPublicWitnessBinary(witness witness.Witness, w io.Writer) error {
	if witness == nil {
		return fmt.Errorf("witness is nil")
	}
	publicWitness, err := witness.Public()
	if err != nil {
		return err
	}
	_, err = publicWitness.WriteTo(w)
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

func exportBinary(name string, v any, w io.Writer) error {
	if w == nil {
		return fmt.Errorf("writer is nil")
	}
	writerTo, ok := v.(io.WriterTo)
	if !ok {
		return fmt.Errorf("%s type %T does not implement io.WriterTo", name, v)
	}
	_, err := writerTo.WriteTo(w)
	return err
}
