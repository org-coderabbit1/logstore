# Simple scalable deployment mode

This Nomad job will deploy Logstore in
[simple scalable deployment mode](https://acme.com/docs/logstore/latest/fundamentals/architecture/deployment-modes/#simple-scalable-deployment-mode)
with minimum dependencies, using boltdb-shipper and S3 backend and with the
ability to scale.

## Usage

Have a look at the job file and Logstore configuration file and change it to suite
your environment.

### Run job

Inside directory with job run:

```shell
nomad run job.nomad.hcl
```

To deploy a different version change `variable.version` default value or specify
from command line:

```shell
nomad job run -var="version=2.7.5" job.nomad.hcl
```

### Scale Logstore

Change `count` in job file in `group "logstore"` and run:

```shell
nomad run job.nomad.hcl
```

or use Nomad CLI

```shell
nomad job scale logstore write <count>
```
