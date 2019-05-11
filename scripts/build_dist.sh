#! /bin/bash

# Check for goreleaser executable.
if ! [ -x "$(command -v goreleaser)" ]; then
    echo "Please install goreleaser before building dist"
    exit 1
fi

goreleaser --snapshot --rm-dist --skip-publish