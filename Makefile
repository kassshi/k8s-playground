run-backend:
	make -C backend run

dev-backend:
	make -C backend dev

dev-frontend:
	make -C frontend dev

gen-protobuf:
	cd proto && buf generate
