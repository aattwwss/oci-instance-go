# oci-instance-go
A golang port of https://github.com/hitrov/oci-arm-host-capacity

This script will attempt to automatically create an instance in oracle cloud infrastructure according to the config in the `.env` file.

# Confiuration
Follow the same configuration guide in [here](https://github.com/hitrov/oci-arm-host-capacity#configuration) and populate the `.env` file accordingly.

1. Go to https://cloud.oracle.com/identity/domains/my-profile/api-keys
2. Click on `Add API Key`
3. Download the private key
4. Click on `Add`
5. Either fill the .env file with the content of the config file, or use the config file directly as described below

# OCI API Key Configuration
I have also added support for the default configuration provider from the official  [GO sdk](https://github.com/oracle/oci-go-sdk). If you have the
config file and private key in their default location of `$HOME/.oci/config`, then you do not need to populate the following 4 fields:

- `OCI_REGION`
- `OCI_USER_ID`
- `OCI_TENANCY_ID`
- `OCI_KEY_FINGERPRINT`

# Running the script
## Included retry with delay
You can set the script to run forever with the specified delay in minute using the `-t` flag.
```shell
# run one time and end
go run .
./oci-instance-go 

#retry every 1 minute
go run . -t=1
./oci-instance-go  -t=1
```

## With prebuilt binaries
In the [release](https://github.com/aattwwss/oci-instance-go/releases) page, download the binary for your OS and platform. Then run the binary
and the `.env` file in the same folder.
```shell
# run one time and end
./oci-instance-go
#retry every 1 minute
./oci-instance-go  -t=1 
```

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



