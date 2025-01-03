#!/bin/bash

filename=$1
filename=`echo $filename | sed -E 's/^([0-9]+)\.(.*)/\1\2\.go/g' | sed -E 's/ /_/g' | awk '{print tolower($0)}'`
touch "$filename"".go"
echo 'package leetcode' > $filename".go"

touch "$filename""_test.go"
echo 'package leetcode' > $filename"_test.go"
