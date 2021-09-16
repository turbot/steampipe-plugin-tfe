connection "tfe" {
  plugin = "tfe"

  # Get your API token per https://www.terraform.io/docs/cloud/users-teams-organizations/api-tokens.html
  # If `token` is not specified, Steampipe will use the token in the `TFE_TOKEN` environment variable, if set
  #token = "hE1hlYILrSqpqh.atlasv1.ARuZuyzl33F71WR55s6ln5GQ1HWIwTDDH3MiRjz7OnpCfaCb1RCF5zGaSncCWmJdCYA"

  # Name of the organization to connect to
  #organization = "example-org"

  # Terraform Cloud or Terraform Enterprise hostname
  # If `hostname` is not specified, the hostname will be determined in the following order:
  #   - The `TFE_HOSTNAME` environment variable, if set; otherwise
  #   - Default to app.terraform.io
  #hostname = "app.terraform.io"
}
