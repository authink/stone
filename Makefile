.DEFAULT_GOAL := package
V := 0.2.5

tidy:
	go mod tidy

fmt:
	go fmt ./...
fmt: tidy

package:
	git tag v$(V)
	git push --tags

package: fmt