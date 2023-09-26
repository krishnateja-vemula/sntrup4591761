#!/usr/bin/env bash

set -eux

# Get the directory of the current script
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

TMPDIR="$(mktemp -d /tmp/prime.XXXXXXXX)"
PUBKEY="${TMPDIR}"/pubkey
PRIKEY="${TMPDIR}"/prikey
CIPHER="${TMPDIR}"/cipher

cleanup() {
        rm -f "${PUBKEY}"
        rm -f "${PRIKEY}"
        rm -f "${CIPHER}"
        rmdir "${TMPDIR}"
}

trap cleanup EXIT

# Time the keygen operation
echo "Timing keygen..."
time {
        $DIR/examples/keygen/keygen /dev/urandom "${PUBKEY}" "${PRIKEY}"
}

# Print the generated keys
echo "Public Key:"
cat "${PUBKEY}"
echo
echo "Private Key:"
cat "${PRIKEY}"

# Time the encap operation
echo "Timing encap..."
time {
        x="$($DIR/examples/encap/encap /dev/urandom ${PUBKEY} ${CIPHER})"
}

# Time the decap operation
echo "Timing decap..."
time {
        y="$($DIR/examples/decap/decap ${PRIKEY} ${CIPHER})"
}

# Assert that the shared secrets match after encap/decap
[ "${x}" == "${y}" ]

