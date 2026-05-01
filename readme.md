# Google Cloud API key rotation checker

This repository contains a Go version of the [API Key Rotation Checker](https://github.com/GoogleCloudPlatform/professional-services/tree/main/tools/api-key-rotation/api_key_rotation_checker) solution developed by [Google Cloud Professional Services](https://github.com/GoogleCloudPlatform/professional-services/tree/main).

The Go version of the *API key rotation checker* adds the following features:

- check for API keys in all projects the user has access to (default) or just the project specified by `--project`.
- format the output as JSON or CSV using the `--format` flag.
- display only API keys that need to be rotated using the `--rotate` flag.

## Usage

*API key rotation checker* is built using Google Cloud SDK libraries for Go.

To interact with Google Cloud, *API key rotation checker* will try to locate local credentials of the current active user via ADC (application default credentials).

Login to Google Cloud using `gcloud`:

```console
gcloud auth login --upate-adc
```

Then run the application (we recommend running it with `-h` or `--help` to get all available options first):

```console
./apikeychecker
```

## Compile from source

As with any Go application, you can cross-compile using:

| Target OS | Target Architecture | Name |
| --------- | ------------------- | ---- |
| `GOOS=darwin` | `GOARCH=arm64` | Mac M1, M2, M3, ... |
| `GOOS=darwin` | `GOARCH=amd64` | Mac Intel |
| `GOOS=linux` | `GOARCH=amd64` | Linux (AMD64) |
| `GOOS=linux` | `GOARCH=arm64` | Linux (ARM64) |
| `GOOS=windows` | `GOARCH=amd64` | Microsoft Windows (64bit) |

```console
<target OS> <target architecture> go build -o apikeycheck *.go
```
