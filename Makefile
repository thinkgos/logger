tag=$(shell git cliff --unreleased --bump --context | jq -r .[0].version)
release:
	git-cliff -r . --bump --with-commit "release: ${tag}" > CHANGELOG.md
	git add .
	git commit -m "release: ${tag}"
	git tag ${tag}
	@echo "👉 Run 'git push origin master --tags' to publish"

.PHONY:
	release