// Package p2p connects nodes into the community network and transfers chunks
// between them (O3, R-01 #29).
//
// Planned design, from the technology choices doc:
//   - go-libp2p host
//   - peer discovery through the Kademlia DHT, plus mDNS on a local network
//   - a private network joined with a pre-shared community key (proposed)
//   - a simple request-response protocol over libp2p streams (proposed)
package p2p
