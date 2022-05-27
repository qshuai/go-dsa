#!/bin/bash

filename=$1
filename=`echo $filename | sed -E 's/^([0-9]+)\.(.*)/\1\2\.go/g' | sed -E 's/ /_/g' | awk '{print tolower($0)}'`
touch "$filename"
echo 'package leetcode' > $filename
