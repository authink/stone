.DEFAULT_GOAL := public
V := 0.1.1

public:
	git tag v$(V)
	git push --tags