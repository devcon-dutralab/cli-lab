# cli-lab

A tiny CLI (`cli-lab greet <name>`) used as the **incident-class exhibit** for
the Devcon offsec lab. The code is deliberately trivial; the story lives in
`AGENTS.md`.

## Why this repo exists

`AGENTS.md` instructs every coding agent to fetch a standards file over an
unauthenticated URL and treat its contents as mandatory operating instructions.
That is the real-world incident class: an agent that fetch-and-obeys a mutable
URL can be reprogrammed by whoever can write to that URL, without touching this
repo at all.
