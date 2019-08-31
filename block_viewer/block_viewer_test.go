package main

import (
	"testing"
)

func TestPow(t *testing.T) {
	if findPower(2, 3) != 8 {
		t.Error("Expected 2^3 to equal 8")
	}
	if findPower(3, 3) != 27 {
		t.Error("Expected 3^3 to equal 27")
	}
}

func TestFromHex(t *testing.T) {
	if fromHex("0025fc4b") != 2489419 {
		t.Error("Expected 0025fc4b to equal 2489419)")
	}
	if fromHex("000000009502f900") != 2500000000 {
		t.Error("Expected 000000009502f900 to equal 2500000000)")
	}
}

func TestConvert(t *testing.T) {
	if convertEndian("79dc738f55b615d094f7bd8693bdc8852bc4b53403c3c528e072e6f8fcb8febe") != "befeb8fcf8e672e028c5c30334b5c42b85c8bd9386bdf794d015b6558f73dc79" {
		t.Error("Expected 79dc738f55b615d094f7bd8693bdc8852bc4b53403c3c528e072e6f8fcb8febe to equal befeb8fcf8e672e028c5c30334b5c42b85c8bd9386bdf794d015b6558f73dc79")
	}
	if convertEndian("befeb8fcf8e672e028c5c30334b5c42b85c8bd9386bdf794d015b6558f73dc79") != "79dc738f55b615d094f7bd8693bdc8852bc4b53403c3c528e072e6f8fcb8febe" {
		t.Error("Expected befeb8fcf8e672e028c5c30334b5c42b85c8bd9386bdf794d015b6558f73dc79 to equal 79dc738f55b615d094f7bd8693bdc8852bc4b53403c3c528e072e6f8fcb8febe")
	}
}