package writer

import (
	"context"
	"github.com/paulmach/orb/geojson"
	go_writer "github.com/whosonfirst/go-writer"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFeature(t *testing.T) {

	ctx := context.Background()

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	fixtures := filepath.Join(cwd, "fixtures")
	feature_path := filepath.Join(fixtures, "101736545.geojson")

	fh, err := os.Open(feature_path)

	if err != nil {
		t.Fatalf("Failed to open %s, %v", feature_path, err)
	}

	defer fh.Close()

	body, err := io.ReadAll(fh)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", feature_path, err)
	}

	f, err := geojson.UnmarshalFeature(body)

	if err != nil {
		t.Fatalf("Failed to unmarshal %s, %v", feature_path, err)
	}

	wr, err := go_writer.NewWriter(ctx, "null://")

	if err != nil {
		t.Fatal(err)
	}

	err = WriteFeature(ctx, wr, f)

	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteBytes(t *testing.T) {

	ctx := context.Background()

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	fixtures := filepath.Join(cwd, "fixtures")
	feature_path := filepath.Join(fixtures, "101736545.geojson")

	fh, err := os.Open(feature_path)

	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	body, err := io.ReadAll(fh)

	if err != nil {
		t.Fatal(err)
	}

	wr, err := go_writer.NewWriter(ctx, "null://")

	if err != nil {
		t.Fatal(err)
	}

	err = WriteBytes(ctx, wr, body)

	if err != nil {
		t.Fatal(err)
	}
}
