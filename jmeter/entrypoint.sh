#!/bin/sh

set -eu

NOW="$(date +%Y-%m-%d_%H%M%S)"
OUT_DIR="out/${NOW}"
mkdir -p "${OUT_DIR}"

THREADS=${1:-1}
SECOND=${2:-1}

jmeter \
  -e \
  -n \
  -JC_THREADS="${THREADS}" \
  -JC_SECONDS="${SECOND}" \
  -t app.jmx \
  -l "${OUT_DIR}/result.jtl" \
  -j "${OUT_DIR}/jmeter.log" \
  -o "${OUT_DIR}/report"
