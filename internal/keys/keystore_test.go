package keys

import (
	"path/filepath"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	dir := t.TempDir()
	s, err := Open(filepath.Join(dir, "keys.json"))
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	k, err := s.Add("alice")
	if err != nil {
		t.Fatalf("Add: %v", err)
	}
	if k.Address == "" {
		t.Fatal("Add returned an empty address")
	}

	got, err := s.Get("alice")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if got.Address != k.Address {
		t.Fatalf("Get address = %q, want %q", got.Address, k.Address)
	}
}

func TestAddIsDeterministic(t *testing.T) {
	cases := []string{"alice", "bob", "carol"}

	for _, name := range cases {
		t.Run(name, func(t *testing.T) {
			first, err := Open(filepath.Join(t.TempDir(), "keys.json"))
			if err != nil {
				t.Fatalf("Open: %v", err)
			}
			a, err := first.Add(name)
			if err != nil {
				t.Fatalf("Add: %v", err)
			}

			second, err := Open(filepath.Join(t.TempDir(), "keys.json"))
			if err != nil {
				t.Fatalf("Open: %v", err)
			}
			b, err := second.Add(name)
			if err != nil {
				t.Fatalf("Add: %v", err)
			}

			if a.Address != b.Address {
				t.Fatalf("address not deterministic: %q != %q", a.Address, b.Address)
			}
		})
	}
}

func TestAddRejectsDuplicate(t *testing.T) {
	s, err := Open(filepath.Join(t.TempDir(), "keys.json"))
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	if _, err := s.Add("alice"); err != nil {
		t.Fatalf("Add: %v", err)
	}
	if _, err := s.Add("alice"); err == nil {
		t.Fatal("Add duplicate: want error, got nil")
	}
}

func TestListSorted(t *testing.T) {
	s, err := Open(filepath.Join(t.TempDir(), "keys.json"))
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	for _, name := range []string{"carol", "alice", "bob"} {
		if _, err := s.Add(name); err != nil {
			t.Fatalf("Add(%s): %v", name, err)
		}
	}
	list, err := s.List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	want := []string{"alice", "bob", "carol"}
	if len(list) != len(want) {
		t.Fatalf("List returned %d keys, want %d", len(list), len(want))
	}
	for i, k := range list {
		if k.Name != want[i] {
			t.Fatalf("List[%d] = %q, want %q", i, k.Name, want[i])
		}
	}
}
