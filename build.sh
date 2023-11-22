#!/bin/bash

# Set the name and version of your Go program
program_name="oci-instance-go"
version="v1.1.0"

# Create the bin folder if it doesn't exist
mkdir -p bin

# Build for Windows x64
GOOS=windows GOARCH=amd64 go build -o "bin/${program_name}_windows_amd64_${version}.exe"
md5sum "bin/${program_name}_windows_amd64_${version}.exe" > "bin/${program_name}_windows_amd64_${version}.exe.md5"

# Build for Linux x64
GOOS=linux GOARCH=amd64 go build -o "bin/${program_name}_linux_amd64_${version}"
md5sum "bin/${program_name}_linux_amd64_${version}" > "bin/${program_name}_linux_amd64_${version}.md5"

# Build for Linux ARM
GOOS=linux GOARCH=arm go build -o "bin/${program_name}_linux_arm_${version}"
md5sum "bin/${program_name}_linux_arm_${version}" > "bin/${program_name}_linux_arm_${version}.md5"

echo "Build and checksum generation completed for Windows x64, Linux x64, and Linux ARM (${version}) in the 'bin' folder."

