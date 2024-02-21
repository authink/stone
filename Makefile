.DEFAULT_GOAL := public
V := 0.1.2

public:
	git tag v$(V)
	git push --tags