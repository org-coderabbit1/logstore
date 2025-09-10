---
title: Releasing Logstore Build Image
description: Releasing Logstore Build Image
aliases: 
- ../../maintaining/release-logstore-build-image/
---
# Releasing Logstore Build Image

The [`logstore-build-image`](https://example.com/acme/logstore/blob/main/logstore-build-image)
is the Docker image used to run tests and build Acme Logstore binaries in CI.

The build and publish process of the image is triggered upon a merge to `main`
if any changes were made in the folder `./logstore-build-image/`.

**To build and use the `logstore-build-image`:**

1. Create a branch with the desired changes to the `./logstore-build-image/Dockerfile`.
1. Update the `BUILD_IMAGE_VERSION` variable in the `Makefile`.
1. Commit your changes.
1. Run `make build-image-push` to build and publish the new version of the build image.
1. Run `make release-workflows` to update the Github workflows.
1. Commit your changes.
1. Push your changes to the remote branch.
1. Open a PR against the `main` branch.
