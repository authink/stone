.DEFAULT_GOAL := package
V := 0.1.4

package:
	git tag v$(V)
	git push --tags