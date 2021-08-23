package main

import "testing"

func TestHelo(t *testing.T) {
  printable := Helo()
  if (printable == "") {
    t.Error("Helo() returned nil")
  }
}
