.DEFAULT_GOAL := package
V := 0.1.5

package:
	git tag v$(V)
	git push --tags