// Package crypto encrypts and decrypts chunks on the owner's device, so a
// host only ever stores ciphertext (O2, R-02 #30).
//
// Planned design, from the technology choices doc:
//   - XChaCha20-Poly1305 (golang.org/x/crypto/chacha20poly1305)
//   - a random key per file, which never leaves the owner's device
//   - a stored hash of the key, checked before decrypting (key commitment)
package crypto
