tests:
	go test && \
		go build && \
		inspec exec test/smoke/gridgo.rb
