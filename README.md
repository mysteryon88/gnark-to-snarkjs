# Gnark to snarkjs

[![Go Reference](https://pkg.go.dev/badge/github.com/mysteryon88/gnark-to-snarkjs.svg)](https://pkg.go.dev/github.com/mysteryon88/gnark-to-snarkjs)

Utilities for exporting **[gnark](https://github.com/ConsenSys/gnark)** proofs and verifying keys:

- **snarkjs format** — JSON compatible with [snarkjs](https://github.com/iden3/snarkjs) (`ExportProof`, `ExportVerifyingKey`).
- **gnark native format** — JSON and binary from gnark structs (`ExportGnarkProof`, `ExportGnarkVerifyingKey`, `ExportPublicWitness`, plus the `*Binary` variants); suitable for [Garaga](https://garaga.gitbook.io/garaga/smart-contract-generators/groth16/generate-and-deploy-your-verifier-contract) and other tools that accept gnark’s native formats.

Supports **Groth16** on curves **BN254** and **BLS12-381**.

## Installation

```bash
go get github.com/mysteryon88/gnark-to-snarkjs@latest
```

## Export for snarkjs

Export proof and verifying key into snarkjs-compatible JSON (e.g. for `snarkjs verify`):

```go
import (
    "os"
    gnarktosnarkjs "github.com/mysteryon88/gnark-to-snarkjs"
)

// proof: *groth16_bn254.Proof or *groth16_bls12381.Proof
var proof = getProof()
proofOut, _ := os.Create("proof.json")
err = gnarktosnarkjs.ExportProof(proof, []string{"35"}, proofOut)

// vk: *groth16_bn254.VerifyingKey or *groth16_bls12381.VerifyingKey
var vk = getVK()
vkOut, _ := os.Create("vk.json")
err = gnarktosnarkjs.ExportVerifyingKey(vk, vkOut)
```

## Export gnark native format

Export proof, verifying key, and public witness as gnark’s native JSON (e.g. for Garaga):

```go
proofOut, _ := os.Create("proof_gnark.json")
err = gnarktosnarkjs.ExportGnarkProof(proof, proofOut)

vkOut, _ := os.Create("vk_gnark.json")
err = gnarktosnarkjs.ExportGnarkVerifyingKey(vk, vkOut)

// Public witness (public inputs/outputs)
schema, _ := frontend.NewSchema(field, &circuit)
publicOut, _ := os.Create("public.json")
err = gnarktosnarkjs.ExportPublicWitness(publicWitness, schema, publicOut)
```

Export the same objects as gnark native binary using gnark's `WriteTo` serialization:

```go
proofBinOut, _ := os.Create("proof_gnark.bin")
err = gnarktosnarkjs.ExportGnarkProofBinary(proof, proofBinOut)

vkBinOut, _ := os.Create("vk_gnark.bin")
err = gnarktosnarkjs.ExportGnarkVerifyingKeyBinary(vk, vkBinOut)

publicBinOut, _ := os.Create("public.bin")
err = gnarktosnarkjs.ExportPublicWitnessBinary(publicWitness, publicBinOut)
```

## Supported curves

- BN254
- BLS12-381

## License

MIT
