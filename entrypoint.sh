#!/bin/sh

set -e

# Install go_ibm_db installer
go install github.com/ibmdb/go_ibm_db/installer@latest

# Change directory to the installer location
cd /pkg/mod/github.com/ibmdb/go_ibm_db@latest/installer

wget https://public.dhe.ibm.com/ibmdl/export/pub/software/data/db2/drivers/odbc_cli/linuxx64_odbc_cli.tar.gz
tar -xzf linuxx64_odbc_cli.tar.gz ../../
