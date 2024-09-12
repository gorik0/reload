build:
	go build  -o a ./prof/main.go

launch_bin_trace_gc:
	GODEBUG=gctrace=1 ./a





