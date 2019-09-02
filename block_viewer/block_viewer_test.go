package main

import (
	"testing"
	"fmt"
	"net/http"
    "io/ioutil"
)

func TestFindPower(t *testing.T) {
	fmt.Println("--- Expecting 2^3 to equal 8")
	if findPower(2, 3) != 8 {
		t.Error("--- Expecting 2^3 to equal 8")
	}
	fmt.Println("--- Expecting 3^3 to equal 27")
	if findPower(3, 3) != 27 {
		t.Error("--- Expecting 3^3 to equal 27")
	}
}

func TestFromHex(t *testing.T) {
	fmt.Println("--- Expecting 0025fc4b to equal 2489419")
	if fromHex("0025fc4b") != 2489419 {
		t.Error("--- Expecting 0025fc4b to equal 2489419)")
	}
	fmt.Println("--- Expecting 000000009502f900 to equal 2500000000")
	if fromHex("000000009502f900") != 2500000000 {
		t.Error("--- Expecting 000000009502f900 to equal 2500000000)")
	}
}

func TestConvertEndian(t *testing.T) {
	fmt.Println("--- Expecting 79dc7300 to equal 0073dc79")
	if convertEndian("79dc7300") != "0073dc79" {
		t.Error("Expected 79dc7300 to equal 0073dc79")
	}
	fmt.Println("--- Expecting befeb8fcf8e672e028c5c30334b5c42b85c8bd9386bdf794d015b6558f73dc79 to equal 79dc738f55b615d094f7bd8693bdc8852bc4b53403c3c528e072e6f8fcb8febe")
	if convertEndian("befeb8fcf8e672e028c5c30334b5c42b85c8bd9386bdf794d015b6558f73dc79") != "79dc738f55b615d094f7bd8693bdc8852bc4b53403c3c528e072e6f8fcb8febe" {
		t.Error("Expected befeb8fcf8e672e028c5c30334b5c42b85c8bd9386bdf794d015b6558f73dc79 to equal 79dc738f55b615d094f7bd8693bdc8852bc4b53403c3c528e072e6f8fcb8febe")
	}
}

func TestConvertHash160(t *testing.T) {
	fmt.Println("--- Expecting 39067f079d1fe9b0df6e2ac0a04f8b6432e78616 to equal 16CXJJLYS9asic8LwUm3yPWRAw4XLfczq")
	if convertHash160("39067f079d1fe9b0df6e2ac0a04f8b6432e78616") != "16CXJJLYS9asic8LwUm3yPWRAw4XLfczqv" {
		t.Error("Expected 39067f079d1fe9b0df6e2ac0a04f8b6432e78616 to equal 16CXJJLYS9asic8LwUm3yPWRAw4XLfczqv")
	}
}

func TestBuildHeader(t *testing.T) {
	fmt.Println("--- Expecting merkleRoot from block 000000000000000001f942eb4bfa0aeccb6a14c268f4c72d5fff17270da771b9 to equal 9b7d5896398581a7ff26be4b3684ddd95a7c1dc1aab1df37cbb2127379ae8584")
	fmt.Println("--- Expecting numTransactions to equal 1031")
	URL := "https://blockchain.info/block/000000000000000001f942eb4bfa0aeccb6a14c268f4c72d5fff17270da771b9?format=hex"
	resp, _ := http.Get(URL)
	testBlock, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()	
	buildHeader(testBlock)
	if head.merkleRoot != "9b7d5896398581a7ff26be4b3684ddd95a7c1dc1aab1df37cbb2127379ae8584" {
		t.Error("Expected merkleRoot to equal 9b7d5896398581a7ff26be4b3684ddd95a7c1dc1aab1df37cbb2127379ae8584")
	}
	if fromHex(head.numTransactions) != 1031 {
		t.Error("Expected numTransactions to equal 1031")
	}
}