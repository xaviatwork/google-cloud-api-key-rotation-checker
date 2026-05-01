# Google Cloud API key rotation checker

This repository contains a Go version of the [API Key Rotation Checker](https://github.com/GoogleCloudPlatform/professional-services/tree/main/tools/api-key-rotation/api_key_rotation_checker) solution developed by [Google Cloud Professional Services](https://github.com/GoogleCloudPlatform/professional-services/tree/main).

The Go version of the API key rotation checker adds the following features:

- check for API keys in all projects the user has access to (default) or just the project specified by `--project`.
- format the output as JSON or CSV using the `--format` flag.
- display only API keys that need to be rotated using the `--rotate` flag.
