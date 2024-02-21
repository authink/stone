.DEFAULT_GOAL := public
V := 0.1.4

public:
	git tag v$(V)
	git push --tags