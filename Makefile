.DEFAULT_GOAL := public
V := 0.1.3

public:
	git tag v$(V)
	git push --tags