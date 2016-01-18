package lendingclub

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func respondWithFixture(w http.ResponseWriter, name string) error {
	f, err := os.Open(filepath.Join("./fixtures", name))
	if err != nil {
		return err
	}

	io.Copy(w, f)
	f.Close()

	return nil
}
