// Package chunk splits files into chunks and reassembles them (O2, R-10 #38).
//
// Planned design, from the technology choices doc (proposed):
//   - each chunk is addressed by the SHA-256 hash of its ciphertext, so a
//     tampered or corrupted chunk is detected on retrieval
//   - chunk size is still an open question
package chunk
