package writer

import (
	"context"
	"github.com/whosonfirst/go-whosonfirst-geojson-v2/feature"
	go_writer "github.com/whosonfirst/go-writer"
	"io/ioutil"
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

	f, err := feature.LoadWOFFeatureFromReader(fh)

	if err != nil {
		t.Fatal(err)
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

func TestWriteFeatureBytes(t *testing.T) {

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

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		t.Fatal(err)
	}

	wr, err := go_writer.NewWriter(ctx, "null://")

	if err != nil {
		t.Fatal(err)
	}

	err = WriteFeatureBytes(ctx, wr, body)

	if err != nil {
		t.Fatal(err)
	}
}
