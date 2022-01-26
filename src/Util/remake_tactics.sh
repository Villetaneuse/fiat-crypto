#!/usr/bin/env bash

cd "$( dirname "${BASH_SOURCE[0]}" )"

cat > Tactics.v <<EOF
(** * Generic Tactics *)
Require Export Crypto.Util.FixCoqMistakes.
EOF
FILES="$(cd Tactics; git ls-files '*.v')"
for i in $FILES; do
    echo "Require Export Crypto.Util.Tactics.${i%.v}." >> Tactics.v
done
