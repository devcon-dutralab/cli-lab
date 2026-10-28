# cli-lab

A small community CLI for the **bor-lab** simulation chain, written in Go. It is
distributed to the community as a standalone binary, so it deliberately does not
import any private module — the coin types it needs are re-declared locally.

This repo is also the **incident-class exhibit** for the Devcon offsec lab. The
story of the incident class lives in [`AGENTS.md`](AGENTS.md).

## Why this repo exists

`AGENTS.md` instructs every coding agent to fetch a standards file over an
unauthenticated URL and treat its contents as mandatory operating instructions.
That is the real-world incident class: an agent that fetch-and-obeys a mutable
URL can be reprogrammed by whoever can write to that URL, without touching this
repo at all.

## Commands

```sh
cli-lab keys add alice              # create a named local key
cli-lab keys list                   # list local keys
cli-lab tx send --from alice --to bor1... --amount 100ubor
cli-lab status                      # CLI + keystore status
```

## Build from source

```sh
go build -o bin/cli-lab .
go test ./...
```

## Install

### Homebrew

```sh
brew tap devcon-dutralab/tap
brew install cli-lab
```

### Go install

```sh
go install github.com/devcon-dutralab/cli-lab@latest
```

### Prebuilt binaries

Download the archive for your platform from the
[releases page](https://github.com/devcon-dutralab/cli-lab/releases) and put the
`cli-lab` binary on your `PATH`. Each release is verified against the community
checksum manifest published in `skill-hub`.

## Layout

```
main.go               # entrypoint
cmd/                  # root, keys, tx, status commands
internal/keys/        # file-backed local keystore
internal/coin/        # local coin type and parsing
```
