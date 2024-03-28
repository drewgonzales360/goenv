#!/bin/bash

# Do not allow use of undefined vars. Use ${VAR:-} to use an undefined VAR
set -o nounset

GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
CLEAR_COLOR='\033[0m'

function fail {
    echo -e "${RED}[error] $1${CLEAR_COLOR}"
}

function success {
    echo -e "${GREEN}[success] $1${CLEAR_COLOR}"
}

function test-command() {
    local cmd="${1:?need a command to test}"
    local expected_sum="${2:?need the md5sum of the expected output}"

    actual_results="$(mktemp tmp/goenv-test.XXXX)"
    (eval "${cmd}" > "${actual_results}" 2>&1)
    actual_sum="$(md5sum "${actual_results}" | awk '{print $1}')"

    if [[ "${actual_sum}" != "${expected_sum}" ]]; then
        fail "The command '${cmd}' failed."
        fail "    actual sum:   ${actual_sum}"
        fail "    expected sum: ${expected_sum}"
        echo "$ ${cmd}"
        cat "${actual_results}"
        echo "=============================================="
        exit 1
    else
        success "'${cmd}' passed."
        echo "=============================================="
    fi
}

test-command "goenv install 1.18"               "1b3aab6f0142d34d991ce0cc8ec3e591"
test-command "sudo goenv install 1.18"          "d38b90a90390b3756e1992a34618c7a4"
test-command "sudo goenv use 1.17.8"            "628121bce55eca109aa4aae8f6ffa010"
test-command "sudo goenv install 1.17.8"        "5119f6686cf0257ff262145d61db303a"
test-command "sudo goenv use 1.18"              "d38b90a90390b3756e1992a34618c7a4"
test-command "goenv use 1.18"                   "1d7e424e4c311c2f67bf53c538367bb3"
test-command "sudo goenv install 1.19.1"        "58339c01166fbb80bfd20e08b801939e"
test-command "sudo goenv install 1.18.1"        "cb2fd36806cd2ed911e63bbd9cdbd385"
test-command "sudo goenv rm 1.18"               "48bacb95c998567dcb71c4c8d80f0580"
test-command "sudo goenv rm 1.18.1"             "7eee8bee6ab2b5c84adead67b7ed0565"
test-command "sudo goenv rm 1.17.8"             "f6d8e601d35578a4567982d8741afd08"
test-command "sudo goenv rm 1.19.1"             "deb74a9475d1bd1e9fa234f78afe4f49"
test-command "sudo goenv install 1.20"          "47ca9a018b21c73b6f001d06823899bd"
