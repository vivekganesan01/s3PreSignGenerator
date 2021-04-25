# AWS S3-PreSign
----
#### Helps to create pre-signed url

> S3PreSigned helps to create a signed url for a given s3 object.

- Helps to open S3 object for 72 hours

###  Pre-requisite:

- Update env.sh


###  To run:
```sh
source env.sh
./s3PreSign -bucketname=xyz -bucketprefix=folder/filename.txt
```

### Usage:
```sh
./s3PreSign help
```

### To build:
```sh
go build ./s3PreSigned.go
```