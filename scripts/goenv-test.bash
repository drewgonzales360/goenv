#!/bin/bash

# Not setting any errtrace, pipefails, or errexits so we can run through this.
# Errors are expected.
function banner() {
    echo "=========================================="
    echo $1
    echo "=========================================="
}
banner "do we warn for non-root users?"
(set -x; goenv install 1.18)

banner "does the install work?"
(set -x; sudo goenv install 1.18)

banner "do we fail when we use an uninstalled version?"
(set -x; sudo goenv use 1.17.8)

banner "can we install another version"
(set -x; sudo goenv install 1.17.8)

banner "can we switch versions"
(set -x; goenv use 1.18)
(set -x; sudo goenv use 1.18)

banner "can we remove versions"
(set -x; goenv ls)
(set -x; sudo goenv rm 1.18)
(set -x; goenv ls)
