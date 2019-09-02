computer$ cd block_viewer  
computer$ go run block_viewer.go 000000000000000001643f7706f3dcbc3a386e4c1bfba852ff628d8024f875b6  

Welcome to Bitcoin Block Viewer  

Header  

previous block:  00000000000000000a3ed9a4e25407518aa854f09fa1981adaae9455a91d1966  
merkle root:  9b7d5896398581a7ff26be4b3684ddd95a7c1dc1aab1df37cbb2127379ae8584  
timestamp:  1432723472 (unix time)  
timestamp:  2015-05-27 03:44:32 -0700 PDT (converted)  
target difficulty:  181686f5  
target difficulty:  4.564495450251177e+10 (converted)  
nonce:  226994584  
number of transactions:  1031    
  
...  

-------------------------------------------------  

computer$ cd block_viewer  
computer$ docker build -t "block:latest" .  
computer$ docker run block:latest 000000000000000001f942eb4bfa0aeccb6a14c268f4c72d5fff17270da771b9  

-------------------------------------------------  

computer$ cd block_viewer  
computer$ go test -v  

=== RUN   TestFindPower  
--- Expecting 2^3 to equal 8  
--- Expecting 3^3 to equal 27  
--- PASS: TestFindPower (0.00s)  
=== RUN   TestFromHex  
--- Expecting 0025fc4b to equal 2489419  
--- Expecting 000000009502f900 to equal 2500000000  
--- PASS: TestFromHex (0.00s)  
=== RUN   TestConvertEndian  
--- Expecting 79dc7300 to equal 0073dc79  
--- Expecting befeb8fcf8e672e028c5c30334b5c42b85c8bd9386bdf794d015b6558f73dc79 to equal 79dc738f55b615d094f7bd8693bdc8852bc4b53403c3c528e072e6f8fcb8febe  
--- PASS: TestConvertEndian (0.00s)  
=== RUN   TestConvertHash160  
--- Expecting 39067f079d1fe9b0df6e2ac0a04f8b6432e78616 to equal 16CXJJLYS9asic8LwUm3yPWRAw4XLfczq  
--- PASS: TestConvertHash160 (0.00s)  
=== RUN   TestBuildHeader  
--- Expecting merkleRoot from block 000000000000000001f942eb4bfa0aeccb6a14c268f4c72d5fff17270da771b9 to equal 9b7d5896398581a7ff26be4b3684ddd95a7c1dc1aab1df37cbb2127379ae8584  
--- Expecting numTransactions to equal 1031  
--- PASS: TestBuildHeader (5.81s)  
PASS  
ok  	block_viewer/block_viewer	5.830s  
 
