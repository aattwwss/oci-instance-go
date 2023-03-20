# oci-instance-go
A golang port of https://github.com/hitrov/oci-arm-host-capacity

# Confiuration
Follow the same configuration guide in [here](https://github.com/hitrov/oci-arm-host-capacity#configuration) and populate the `.env` file accordingly.

# OCI API Key Configuration
I have also added support for the default configuration provider from the official  [GO sdk](https://github.com/oracle/oci-go-sdk). If you are already have the
config file and private key in their default location of `$HOME/.oci/config`, then you do not need to populate the following 4 fields:

- `OCI_REGION`
- `OCI_USER_ID`
- `OCI_TENANCY_ID`
- `OCI_KEY_FINGERPRINT`

# Running the script
## Without compiling
```shell
git clone https://github.com/aattwwss/oci-instance-go.git
cd oci-instance-go
cp /path/to/.env .
go run .
```

## With compiling
As GO allows you to compile the program into an executable, we run the executable with the `.env` file anywhere within the same platform.
```shell
git clone https://github.com/aattwwss/oci-instance-go.git
cd oci-instance-go
cp /path/to/.env .
go build .
./oci-instance-go
```


## With prebuilt binaries
In the [release](https://github.com/aattwwss/oci-instance-go/releases) page, download the binary for your OS and platform. Then run the binary
and the `.env` file in the same folder.
```shell
./oci-instance-go
```
