# Distributed Community Cloud

A storage node for a trusted community. Members' devices pool their spare disk space, and the network keeps data durable as devices go on and offline, without a centralised coordinator. Files are encrypted on the owner's device, so hosts only ever hold ciphertext.

This is the product for my MSci Software Engineering dissertation at the University of Glasgow. The dissertation itself, the literature review and the design decisions live in [Distributed-Community-Cloud-Dissertation](https://github.com/callumcooperlee/Distributed-Community-Cloud-Dissertation), in particular:

- [Existing systems review](https://github.com/callumcooperlee/Distributed-Community-Cloud-Dissertation/blob/main/research/existing-systems-review.md)
- [Technology choices](https://github.com/callumcooperlee/Distributed-Community-Cloud-Dissertation/blob/main/research/technology-choices.md)

Planning is tracked in this repo's [issues](https://github.com/callumcooperlee/Distributed-Community-Cloud/issues) and [milestones](https://github.com/callumcooperlee/Distributed-Community-Cloud/milestones). Requirements are the `R-XX` issues.

## Status

Early scaffolding. The `dcc` commands exist but are not implemented yet.

## Layout

| Path | Purpose | Objective |
|---|---|---|
| `cmd/dcc` | Command-line interface | O2 |
| `internal/crypto` | Client-side encryption (XChaCha20-Poly1305) | O2 |
| `internal/chunk` | Chunking and content addressing | O2 |
| `internal/store` | Local chunk storage within a quota | O2 |
| `internal/p2p` | Peer discovery (Kademlia DHT) and chunk transfer (libp2p) | O3 |

## Build and test

Requires Go 1.24 or later.

```sh
go build -o bin/dcc ./cmd/dcc   # build the CLI (bin/dcc.exe on Windows)
go test ./...                   # run the tests
go vet ./...                    # static checks
gofmt -l .                      # list files needing formatting
```

CI runs formatting, vet, build and tests on every push to `main` and every pull request.
