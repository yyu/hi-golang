// http://tour.golang.org/#63

package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r13r *rot13Reader) decode(b byte) byte {
    switch {
    case (b >= 'a' && b <= 'm') || (b >= 'A' && b <= 'M'):
        return b + 13
    case (b >= 'n' && b <= 'z') || (b >= 'N' && b <= 'Z'):
        return b - 13
    }
    return b
}

func (r13r *rot13Reader) Read(p []byte) (n int, err error) {
    n, err = r13r.r.Read(p)
    if err != nil {
        return 0, err
    }
    for i := 0; i < n; i++ {
        p[i] = r13r.decode(p[i])
    }
    return n, nil
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
