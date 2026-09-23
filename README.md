# oci-instance-go
A golang port of https://github.com/hitrov/oci-arm-host-capacity

This script will attempt to automatically create an instance in oracle cloud infrastructure according to the config in the `.env` file.

# Confiuration
1. Go to https://cloud.oracle.com/identity/domains/my-profile/api-keys
2. Click on `Add API Key`
3. Download the private key
4. Click on `Add`
5. Copy the config and place it at the default config location the SDK reads from:
   - Linux / macOS: `~/.oci/config` (i.e. `/home/<you>/.oci/config` or `/Users/<you>/.oci/config`)
   - Windows: `%USERPROFILE%\.oci\config` (e.g. `C:\Users\<you>\.oci\config`)
   - The private key you downloaded in step 3 must be referenced by the `key_file` entry in the config and exist at that path.

# Running the script
## Included retry with delay
You can set the script to run forever with the specified delay in minute using the `-t` flag.
```shell
# run one time and end
go run .
./oci-instance-go 

#retry every 15 minutes
go run . -t=15
./oci-instance-go  -t=15
```

## With prebuilt binaries
In the [release](https://github.com/aattwwss/oci-instance-go/releases) page, download the binary for your OS and platform. Then run the binary
and the `.env` file in the same folder.
```shell
# run one time and end
./oci-instance-go

#retry every 15 minutes
./oci-instance-go  -t=1 5
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



