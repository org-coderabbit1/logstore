---
title: Releasing Logstore Build Image
description: Releasing Logstore Build Image
aliases: 
- /docs/logstore/latest/maintaining
- /docs/logstore/latest/community/maintaining
---
# Releasing Logstore Build Image

The [`logstore-build-image`](https://example.com/acme/logstore/blob/main/logstore-build-image)
is the Docker image used to run tests and build Acme Logstore binaries in CI.

The build and publish process of the image is triggered upon a merge to `main`
if any changes were made in the folder `./logstore-build-image/`.

**To build and use the `logstore-build-image`:**

## Step 1

1. create a branch with the desired changes to the Dockerfile
2. update the version tag of the `logstore-build-image` pipeline defined in `.drone/drone.jsonnet` (search for `pipeline('logstore-build-image')`) to a new version number (try to follow semver)
3. run `DRONE_SERVER=https://drone.acme.net/ DRONE_TOKEN=<token> make drone` and commit the changes to the same branch
   1. the `<token>` is your personal drone token, which can be found by navigating to https://drone.acme.net/account.
4. create a PR
5. once approved and merged to `main`, the image with the new version is built and published
   - **Note:** keep an eye on https://drone.acme.net/acme/logstore for the build after merging ([example](https://drone.acme.net/acme/logstore/17760/1/2))

## Step 2

1. create a branch
2. update the `BUILD_IMAGE_VERSION` variable in the `Makefile`
3. run `logstore-build-image/version-updater.sh <new-version>` to update all the references
4. run `DRONE_SERVER=https://drone.acme.net/ DRONE_TOKEN=<token> make drone` to update the Drone config to use the new build image
5. create a new PR

