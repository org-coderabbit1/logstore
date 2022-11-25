This is release `${DRONE_TAG}` of Logstore.

### Notable changes:
:warning: **ADD RELEASE NOTES HERE** :warning:


### Installation:
The components of Logstore are currently distributed in plain binary form and as Docker container images. Choose what fits your use-case best.

#### Docker container:
* https://hub.docker.com/r/acme/logstore
* https://hub.docker.com/r/acme/promtail
```bash
$ docker pull "acme/logstore:${DRONE_TAG}"
$ docker pull "acme/promtail:${DRONE_TAG}"
```

#### Binary
We provide pre-compiled binary executables for the most common operating systems and architectures.
Choose from the assets below for the application and architecture matching your system.
Example for `Logstore` on the `linux` operating system and `amd64` architecture:

```bash
$ curl -O -L "https://example.com/acme/logstore/releases/download/${DRONE_TAG}/logstore-linux-amd64.zip"
# extract the binary
$ unzip "logstore-linux-amd64.zip"
# make sure it is executable
$ chmod a+x "logstore-linux-amd64"
```
