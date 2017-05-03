#!/usr/bin/env bash

go run main.go > out/surface.svg

# Verify that surface.svg does not include 'Inf' nor 'NaN'.
if grep -q -e "Inf" -e "NaN" out/surface.svg; then
  echo "FAIL: surface.svg includes 'Inf' or 'NaN'."
else
  echo "SUCCESS: surface.svg does not include 'Inf' nor 'NaN'."
fi
