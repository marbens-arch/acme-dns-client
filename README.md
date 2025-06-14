# acme-dns-client

A client software (fork of [acme-dns/acme-dns-client](https://github.com/acme-dns/acme-dns-client)) for [acme-dns](https://github.com/joohoi/acme-dns) with emphasis on usability and guidance through
setup and additional security safeguard mechanisms. It is designed for usage with existing ACME clients with minimal
configuration.

## Purpose of the fork
To continue active development of the inactively developed upstream. Current differences between this fork and upstream include:

- Modules have been updated
- CAA record handling has been updated for [production](https://community.letsencrypt.org/t/enabling-acme-caa-account-and-method-binding/189588/3)
- goreleaser config has been updated to v2
- Releases and the tags associated with them are signed (valid PGP keys [906A492C534B06DFC0A245D18F5BE03D3FF2EF28](https://keyserver.ubuntu.com/pks/lookup?search=906A492C534B06DFC0A245D18F5BE03D3FF2EF28&fingerprint=on&op=index)). If a release or a tag associated with one is unsigned or signed by a key not listed here, please contact acme-dns-client-security@marcusspencer.us.
- Releases package source tarballs
- Binary releases for more platforms

## Installation

- [Download](https://github.com/marbens-arch/acme-dns-client/releases/latest) a prebuilt binary from 
  [releases page](https://github.com/marbens-arch/acme-dns-client/releases/latest), unpack and run!

  _or_
- If you have recent go compiler installed: `go get -u github.com/marbens-arch/acme-dns-client` (the same command works for updating)

  _or_
- git clone https://github.com/marbens-arch/acme-dns-client ; cd acme-dns-client ; go get ; go build

## Features

- acme-dns account pre-registration
- Guided CNAME record creation
- Guided CAA record creation
- Modular ACME client support for CAA record creation guidance (for ACME-CAA accounturi)
- Configuration checks to ensure operation (CNAME record, account exisence)
- Interactive setup

## Example usage with Certbot

### 1. Create a new ACME account

```
# sudo certbot register
```

This creates a new ACME account that is used internally by Certbot. In case you are not planning to set up
CAA record, this step can be omitted.

### 2. Create a new acme-dns account for your domain and set it up

```
# sudo acme-dns-client register -d your.domain.example.org -s https://auth.acme-dns.io
```

This attempts to create a new account to acme-dns instance running at `auth.acme-dns.io`. 
After account creation, the user is guided through proper CNAME record creation for the main DNS zone for domain
`your.domain.example.org`.

Optionally the same will be done for CAA record creation. `acme-dns-client` will attempt to read the account URL from
active Certbot configuration (created in step 1)

`acme-dns-client` will monitor the DNS record changes to ensure they are set up correctly.

### 3. Run Certbot to obtain a new certificate

```
# sudo certbot certonly --manual --preferred-challenges dns \
    --manual-auth-hook 'acme-dns-client' -d your.domain.example.org 
```

This runs Certbot and instructs it to obtain a new certificate for domain `your.domain.example.org` by using a DNS 
challenge and `acme-dns-client` as the authenticator. After successfully obtaining the new certificate this configuration
will be saved in Certbot configuration and will be automatically reused when it renews the certificate.

## Usage

```
acme-dns-client - v0.1

Usage:  acme-dns-client COMMAND [OPTIONS]

Commands:
  register              Register a new acme-dns account for a domain
  check                 Check the configuration and settings of existing acme-dns accounts
  list                  List all the existing acme-dns accounts and perform simple CNAME checks for them

Options:
  --help                Print this help text

To get help for specific command, use:
  acme-dns-client COMMAND --help

EXAMPLE USAGE:
  Register a new acme-dns account for domain example.org:
    acme-dns-client register -d example.org
  
  Register a new acme-dns account for domain example.org, allow updates only from 198.51.100.0/24:
    acme-dns-client register -d example.org -allow 198.51.100.0/24

  Check the configuration of example.org and the corresponding acme-dns account:
    acme-dns-client check -d example.org

  Check the configuration of all the domains and acme-dns accounts registered on this machine:
    acme-dns-client check

  Print help for a "register" command:
    acme-dns-client register --help

```
