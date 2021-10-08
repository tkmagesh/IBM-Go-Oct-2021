package main

import (
	"io"
	"os"
)

type AlphaReader string

func (r AlphaReader) Read(p []byte) (n int, err error) {
	count := 0
	for i := 0; i < len(r); i++ {
		if (r[i] >= 'A' && r[i] <= 'Z') || (r[i] >= 'a' && r[i] <= 'z') {
			p[count] = r[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	r := AlphaReader("3475795Tempor aliqua culpa &(*^&% irure eu excepteur labore incididunt dolore nostrud incididunt officia. Culpa in proident aute anim. Irure commodo eu nisi sunt id officia ut incididunt enim reprehenderit ad ullamco proident. Consequat enim occaecat ullamco amet sit qui fugiat dolor consectetur Lorem velit.")
	io.Copy(os.Stdout, r)
}
