# Gnark to snarkjs

[![Go Reference](https://pkg.go.dev/badge/github.com/mysteryon88/gnark-to-snarkjs.svg)](https://pkg.go.dev/github.com/mysteryon88/gnark-to-snarkjs)

Utilities for exporting **[gnark](https://github.com/ConsenSys/gnark)** proofs and verifying keys into a format compatible with **[snarkjs](https://github.com/iden3/snarkjs)**.  
Currently supports **Groth16** on curves **BN254** and **BLS12-381**.

## Installation

```bash
go get github.com/mysteryon88/gnark-to-snarkjs@latest
```

## Example

Here is a full example with a simple circuit, proof generation, verification, and exporting proof + verifying key into snarkjs-compatible JSON.

```go
import (
    "os"
    gnarktosnarkjs "github.com/mysteryon88/gnark-to-snarkjs"
    groth16_bn254 "github.com/consensys/gnark/backend/groth16/bn254"
)

// proof: *groth16_bn254.Proof or *groth16_bls12381.Proof
var proof = getProof()
proof_out, _ := os.Create("proof.json")
err = gnarktosnarkjs.ExportProof(proof, []string{"35"}, proof_out)

// vk: *groth16_bn254.VerifyingKey or *groth16_bls12381.VerifyingKey
var vk = getVK()
vk_out, _ := os.Create(VKeyPath)
err = gnarktosnarkjs.ExportVerifyingKey(vk, vk_out)
```

Both `proof.json` and `vk.json` are fully compatible with `snarkjs`, so you can directly use them with the `snarkjs verify` command.

## Supported Curves

- BN254
- BLS12-381

## License

MIT
