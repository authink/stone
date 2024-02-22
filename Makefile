.DEFAULT_GOAL := package
V := 0.1.7

tidy:
	go mod tidy

package:
	git tag v$(V)
	git push --tags