package fitz

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestImage(t *testing.T) {
	doc, err := New(filepath.Join("testdata", "test.pdf"))
	if err != nil {
		t.Error(err)
	}

	defer doc.Close()

	tmpDir, err := ioutil.TempDir(os.TempDir(), "fitz")
	if err != nil {
		t.Error(err)
	}

	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			t.Error(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.jpg", n)))
		if err != nil {
			t.Error(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			t.Error(err)
		}

		f.Close()
	}
}

func TestImageFromMemory(t *testing.T) {
	b, err := ioutil.ReadFile(filepath.Join("testdata", "test.pdf"))
	if err != nil {
		t.Error(err)
	}

	doc, err := NewFromMemory(b)
	if err != nil {
		t.Error(err)
	}

	defer doc.Close()

	tmpDir, err := ioutil.TempDir(os.TempDir(), "fitz")
	if err != nil {
		t.Error(err)
	}

	defer os.RemoveAll(tmpDir)

	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			t.Error(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.jpg", n)))
		if err != nil {
			t.Error(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			t.Error(err)
		}

		f.Close()
	}
}

func TestLinks(t *testing.T) {
	doc, err := New(filepath.Join("testdata", "test.pdf"))
	if err != nil {
		t.Error(err)
	}

	defer doc.Close()

	links, err := doc.Links(2)
	if err != nil {
		t.Error(err)
	}

	if len(links) != 1 {
		t.Error("expected 1 link, got", len(links))
	}

	if links[0].URI != "https://creativecommons.org/licenses/by-nc-sa/4.0/" {
		t.Error("expected empty URI, got", links[0].URI)
	}
}

func TestText(t *testing.T) {
	doc, err := New(filepath.Join("testdata", "test.pdf"))
	if err != nil {
		t.Error(err)
	}

	defer doc.Close()

	tmpDir, err := ioutil.TempDir(os.TempDir(), "fitz")
	if err != nil {
		t.Error(err)
	}

	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			t.Error(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.txt", n)))
		if err != nil {
			t.Error(err)
		}

		_, err = f.WriteString(text)
		if err != nil {
			t.Error(err)
		}

		f.Close()
	}
}

func TestHTML(t *testing.T) {
	doc, err := New(filepath.Join("testdata", "test.pdf"))
	if err != nil {
		t.Error(err)
	}

	defer doc.Close()

	tmpDir, err := ioutil.TempDir(os.TempDir(), "fitz")
	if err != nil {
		t.Error(err)
	}

	for n := 0; n < doc.NumPage(); n++ {
		html, err := doc.HTML(n, true)
		if err != nil {
			t.Error(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.html", n)))
		if err != nil {
			t.Error(err)
		}

		_, err = f.WriteString(html)
		if err != nil {
			t.Error(err)
		}

		f.Close()
	}
}

func TestSVG(t *testing.T) {
	doc, err := New(filepath.Join("testdata", "test.pdf"))
	if err != nil {
		t.Error(err)
	}

	defer doc.Close()

	tmpDir, err := ioutil.TempDir(os.TempDir(), "fitz")
	if err != nil {
		t.Error(err)
	}

	for n := 0; n < doc.NumPage(); n++ {
		svg, err := doc.SVG(n)
		if err != nil {
			t.Error(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.svg", n)))
		if err != nil {
			t.Error(err)
		}

		_, err = f.WriteString(svg)
		if err != nil {
			t.Error(err)
		}

		f.Close()
	}
}

func TestToC(t *testing.T) {
	doc, err := New(filepath.Join("testdata", "test.pdf"))
	if err != nil {
		t.Error(err)
	}

	defer doc.Close()

	_, err = doc.ToC()
	if err != nil {
		t.Error(err)
	}
}

func TestMetadata(t *testing.T) {
	doc, err := New(filepath.Join("testdata", "test.pdf"))
	if err != nil {
		t.Error(err)
	}

	defer doc.Close()

	meta := doc.Metadata()
	if len(meta) == 0 {
		t.Error(fmt.Errorf("metadata is empty"))
	}
}
