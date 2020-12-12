#!/bin/bash

set -e

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

curl -sS "https://adventofcode.com/2020/day/1/input" > input.txt

echo Done
