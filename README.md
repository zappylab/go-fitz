## go-fitz

Go wrapper for [MuPDF](http://mupdf.com/) fitz library that can extract pages from PDF and EPUB documents as images, text, html or svg.

### Build tags

* `extlib` - use external MuPDF library
* `static` - build with static external MuPDF library (used with `extlib`)
* `pkgconfig` - enable pkg-config (used with `extlib`)
* `musl` - use musl compiled library
    
### Example
```go
package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sawdustofmind/go-fitz"
)

func main() {
	doc, err := fitz.New("test.pdf")
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	tmpDir, err := ioutil.TempDir(os.TempDir(), "fitz")
	if err != nil {
		panic(err)
	}

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.jpg", n)))
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}

		f.Close()
	}

	// Extract pages as text
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.txt", n)))
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString(text)
		if err != nil {
			panic(err)
		}

		f.Close()
	}

	// Extract pages as html
	for n := 0; n < doc.NumPage(); n++ {
		html, err := doc.HTML(n, true)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.html", n)))
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString(html)
		if err != nil {
			panic(err)
		}

		f.Close()
	}
}
```
