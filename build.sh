#!/usr/bin/env bash

APP_NAME="pagesnatcher"
OUTPUT_DIR="bin"

mkdir -p $OUTPUT_DIR

platforms=("linux/amd64" "darwin/amd64" "darwin/arm64" "windows/amd64")

for platform in "${platforms[@]}"; do
    GOOS=${platform%/*}
    GOARCH=${platform#*/}
    output_name="$OUTPUT_DIR/${APP_NAME}-${GOOS}-${GOARCH}"

    if [ "$GOOS" == "windows" ]; then
        output_name+=".exe"
    fi

    echo "Building for $GOOS/$GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
done

echo "Builds complete!"
