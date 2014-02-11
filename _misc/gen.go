// Copyright 2011 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// capgen is an utility to test captcha generation.
package main

import (
    "github.com/kaz8/go-captcha"
    "flag"
    "fmt"
    "log"
    "os"
)

var (
    flagLen  = flag.Int("l", 6, "length of captcha")
    flagImgW = flag.Int("w", captcha.StdWidth, "image captcha width")
    flagImgH = flag.Int("h", captcha.StdHeight, "image captcha height")
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: captcha [flags] filename\n")
    flag.PrintDefaults()
}

func main() {
    log.SetFlags(0)
    flag.Parse()
    fname := flag.Arg(0)
    if fname == "" {
        usage()
        os.Exit(2)
    }
    f, err := os.Create(fname)
    if err != nil {
        log.Fatalf("%s", err)
    }
    defer f.Close()
    err = captcha.WriteImage(f,
        captcha.RandomDigits(*flagLen), *flagImgW, *flagImgH)
    if err != nil {
        log.Fatalf("%s", err)
    }
}
