# oci-instance-go
A golang port of https://github.com/hitrov/oci-arm-host-capacity

# OCI API Key Configuration
The official GO sdk is used in this project, where it will attempt to load the API key from `$HOME/.oci/config`.
This script will do the same so you need to specify the config file location in `OCI_CONFIG_PATH` env.
