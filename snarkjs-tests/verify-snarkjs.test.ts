import { readFileSync, existsSync } from "node:fs";
import { dirname, resolve } from "node:path";
import { fileURLToPath } from "node:url";
import { groth16 } from "snarkjs";
import { describe, it, expect, beforeAll } from "vitest";

process.env.FFJAVASCRIPT_MAX_WORKERS = "0";

const __dirname = dirname(fileURLToPath(import.meta.url));

beforeAll(() => {
  // @ts-ignore
  global.Worker = undefined;
});

function loadJson(p: string): any {
  const paths = [
    resolve(__dirname, p),
    resolve(__dirname, "cubic", p),
    resolve(__dirname, "..", "test", "cubic", p),
    resolve(__dirname, "..", "cubic", p),
    resolve(__dirname, "..", p),
    resolve(__dirname, "keys", p.split("/").pop()!),
    resolve(__dirname, "proofs", p.split("/").pop()!),
    resolve(__dirname, "..", "test", "cubic", "keys", p.split("/").pop()!),
    resolve(__dirname, "..", "test", "cubic", "proofs", p.split("/").pop()!),
    resolve(__dirname, "..", "keys", p.split("/").pop()!),
    resolve(__dirname, "..", "proofs", p.split("/").pop()!),
  ];

  for (const full of paths) {
    if (existsSync(full)) {
      return JSON.parse(readFileSync(full, "utf8"));
    }
  }

  throw new Error(
    `File not found: ${p}. Searched in: \n - ${paths.join("\n - ")}`,
  );
}

// --- Groth16 Tests ---

describe("(Cubic) snarkjs verify (Groth16) BLS12-381", () => {
  it("verifies proof_groth16_bls12381.json with verification_key_groth16_bls12381.json", async () => {
    const vkey = loadJson("keys/verification_key_groth16_bls12381.json");
    const proof = loadJson("proofs/proof_groth16_bls12381.json");

    const ok = await groth16.verify(vkey, proof.publicSignals, proof);
    expect(ok).toBe(true);
  });

  it("rejects proof_groth16_bls12381.json with wrong publicSignals", async () => {
    const vkey = loadJson("keys/verification_key_groth16_bls12381.json");
    const proof = loadJson("proofs/proof_groth16_bls12381.json");

    const wrongPublicSignals: string[] = ["99"];
    const ok = await groth16.verify(vkey, wrongPublicSignals, proof);
    expect(ok).toBe(false);
  });
});

describe("(Cubic) snarkjs verify (Groth16) BN254", () => {
  it("verifies proof_groth16_bn254.json with verification_key_groth16_bn254.json", async () => {
    const vkey = loadJson("keys/verification_key_groth16_bn254.json");
    const proof = loadJson("proofs/proof_groth16_bn254.json");

    const ok = await groth16.verify(vkey, proof.publicSignals, proof);
    expect(ok).toBe(true);
  });

  it("rejects proof_groth16_bn254.json with wrong publicSignals", async () => {
    const vkey = loadJson("keys/verification_key_groth16_bn254.json");
    const proof = loadJson("proofs/proof_groth16_bn254.json");

    const wrongPublicSignals: string[] = ["99"];
    const ok = await groth16.verify(vkey, wrongPublicSignals, proof);
    expect(ok).toBe(false);
  });
});
