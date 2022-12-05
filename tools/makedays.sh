#!/bin/bash

mkdir days

for i in {1..25}
do
    daystring="day$i"
    funcname="Day$i"
    mkdir "internal/$daystring"
    sed -e "s/package_name/$daystring/g" -e "s/function_name/$funcname/g" tools/template.txt > days/$daystring/$daystring.go
    touch days/$daystring/$daystring.txt
done