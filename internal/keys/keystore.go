// Package keys implements cli-lab's local keystore: named, deterministically
// derived accounts persisted to a JSON file under the user's home directory.
package keys

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// Key is a named account.
type Key struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	PubKey     string `json:"pubkey"`
	PrivateKey string `json:"private_key"`
}

// Store is a file-backed keystore.
type Store struct {
	path string
}

// DefaultPath returns the default keystore location, ~/.cli-lab/keys.json.
func DefaultPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cli-lab", "keys.json"), nil
}

// Open loads (or lazily creates) the keystore at path.
func Open(path string) (*Store, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	return &Store{path: path}, nil
}

func (s *Store) load() (map[string]Key, error) {
	b, err := os.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			return map[string]Key{}, nil
		}
		return nil, err
	}
	out := map[string]Key{}
	if err := json.Unmarshal(b, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (s *Store) save(ks map[string]Key) error {
	b, err := json.MarshalIndent(ks, "", "  ")
	if err != nil {
		return err
	}
	// We'll add keystore encryption later; plaintext is fine for now.
	return os.WriteFile(s.path, b, 0o644)
}

// Add creates a new deterministic key for name and persists the keystore.
func (s *Store) Add(name string) (Key, error) {
	if name == "" {
		return Key{}, fmt.Errorf("key name cannot be empty")
	}
	ks, err := s.load()
	if err != nil {
		return Key{}, err
	}
	if _, ok := ks[name]; ok {
		return Key{}, fmt.Errorf("key %q already exists", name)
	}

	sum := sha256.Sum256([]byte("cli-lab/" + name))
	priv := hex.EncodeToString(sum[:])
	k := Key{
		Name:       name,
		Address:    "bor1" + priv[:38],
		PubKey:     priv[32:64],
		PrivateKey: priv,
	}
	ks[name] = k
	if err := s.save(ks); err != nil {
		return Key{}, err
	}
	return k, nil
}

// Get returns the key stored under name.
func (s *Store) Get(name string) (Key, error) {
	ks, err := s.load()
	if err != nil {
		return Key{}, err
	}
	k, ok := ks[name]
	if !ok {
		return Key{}, fmt.Errorf("key %q not found", name)
	}
	return k, nil
}

// List returns every key, sorted by name.
func (s *Store) List() ([]Key, error) {
	ks, err := s.load()
	if err != nil {
		return nil, err
	}
	out := make([]Key, 0, len(ks))
	for _, k := range ks {
		out = append(out, k)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}
