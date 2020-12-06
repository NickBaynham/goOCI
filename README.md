# goOCI
### OCI Examples in GO

After installing, create a file with your configuration:

export TF_VAR_tenancy_ocid=<your tenancy ocid>
export TF_VAR_user_ocid=<your local user ocid with admin privileges>
export TF_VAR_fingerprint=<your api key fingerprint>
export TF_VAR_private_key_path=<path to your pem file>
export TF_VAR_compartment_ocid=<your compartment ocid>
export TF_VAR_region=<your region id>

Then source from the file to bring the environment variables into memory:
. ./set_env_variables

Run the application:

go run main.go
