package main

import "testing"

func TestHello(t *testing.T) {
    want := "A - blah"
    if got := main(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
