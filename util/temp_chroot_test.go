package util_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-billy/v5/util"
)

func TestTempDefaultsRespectChroot(t *testing.T) {
	root := memfs.New()
	fs, err := root.Chroot("workspace")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("directory", func(t *testing.T) {
		name, err := util.TempDir(fs, "", "txn-")
		if err != nil {
			t.Fatal(err)
		}

		if got, want := filepath.Clean(filepath.Dir(name)), filepath.Clean(".tmp"); got != want {
			t.Fatalf("TempDir path = %q, want directory %q", name, want)
		}
		if _, err := fs.Stat(name); err != nil {
			t.Fatalf("temporary directory is not reachable through chroot: %v", err)
		}
	})

	t.Run("file", func(t *testing.T) {
		f, err := util.TempFile(fs, "", "cache-")
		if err != nil {
			t.Fatal(err)
		}
		name := f.Name()
		if _, err := f.Write([]byte("persistent-state")); err != nil {
			t.Fatal(err)
		}
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}

		if got, want := filepath.Clean(filepath.Dir(name)), filepath.Clean(".tmp"); got != want {
			t.Fatalf("TempFile path = %q, want directory %q", name, want)
		}
		data, err := util.ReadFile(fs, name)
		if err != nil {
			t.Fatal(err)
		}
		if got, want := string(data), "persistent-state"; got != want {
			t.Fatalf("temporary file contents = %q, want %q", got, want)
		}
	})

	t.Run("root filesystem compatibility", func(t *testing.T) {
		rootFS := memfs.New()
		name, err := util.TempDir(rootFS, "", "root-")
		if err != nil {
			t.Fatal(err)
		}

		if got, want := filepath.Clean(filepath.Dir(name)), filepath.Clean(os.TempDir()); got != want {
			t.Fatalf("root TempDir directory = %q, want %q", got, want)
		}
	})
}
