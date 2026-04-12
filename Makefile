run-backend:
	make -C backend run

gen-protobuf:
	cd proto && buf generate
