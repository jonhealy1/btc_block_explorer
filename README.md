computer$ cd block_viewer  
computer$ go run block_viewer.go 000000000000000001643f7706f3dcbc3a386e4c1bfba852ff628d8024f875b6  

Welcome to Bitcoin Block Viewer  

...  

-------------------------------------------------  

computer$ cd block_viewer  
computer$ docker build -t "block:latest" .  
computer$ docker run block:latest 000000000000000001f942eb4bfa0aeccb6a14c268f4c72d5fff17270da771b9  

-------------------------------------------------  

computer$ cd block_viewer  
computer$ go test -v  

=== RUN   TestPow  
--- PASS: TestPow (0.00s)  
=== RUN   TestFromHex  
--- PASS: TestFromHex (0.00s)  
=== RUN   TestConvert  
--- PASS: TestConvert (0.00s)  
PASS  
ok  	block_viewer/block_viewer	0.017s  
