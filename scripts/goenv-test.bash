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
    (eval "${cmd}" > "${actual_results}")
    actual_sum="$(md5sum "${actual_results}" | awk '{print $1}')"

    if [[ "${actual_sum}" != "${expected_sum}" ]]; then
        fail "'${cmd}' failed."
        fail "    actual sum:   ${actual_sum}"
        fail "    expected sum: ${expected_sum}"
        echo "command output:"
        cat "${actual_results}"
        echo "=============================================="
        exit 1
    else
        success "'${cmd}' passed."
        echo "=============================================="
    fi
}

test-command "goenv install 1.18" "a13b6fd35d6f1cb0475115074536b618"
test-command "sudo goenv install 1.18" "1ca020876e6ad587573a0122878d3e9f"
test-command "sudo goenv use 1.17.8" "47463aa8aaeea547d598684ae198abac"
test-command "sudo goenv install 1.17.8" "507166d58c5a48cd27e235c0afab5e7f"
test-command "sudo goenv use 1.18" "1ca020876e6ad587573a0122878d3e9f"
test-command "goenv use 1.18" "c3467966172777b975182e2ebf6f0faa"
test-command "sudo goenv install 1.19.1" "ebf6e06718e42b0734ed60bdb15e888a"
test-command "sudo goenv install 1.18.1" "d46e08f36378e7e456d0c57b6d71eec4"
test-command "sudo goenv rm 1.18" "4d7eecec7a4bb4e3090faf9244f53450"
test-command "sudo goenv rm 1.18.1" "8d1d6cd382684ae31224786784b71da9"
test-command "sudo goenv rm 1.17.8" "6956e3279e07d255f1c0a9362b50ddff"
test-command "sudo goenv rm 1.19.1" "c59c77054ca02a6cbb01e44460caa3e3"
