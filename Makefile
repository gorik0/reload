bench_to_file:
	 go test ./prof/... -bench=Lines___Spli -benchmem -benchtime=1s > split.be \
 	go test ./prof/... -bench=BYTES  -benchmem -benchtime=1s > bytes.be \
	go test ./prof/... -bench=BUFIO  -benchmem -benchtime=1s > bufio.be


bench_stat:
	 benchstat  split.be  bufio.be bytes.be > benchstat.tx





