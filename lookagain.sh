#!bin/bash
find . -name "*.sh" | sort -r | cut -d "/" -f2 | sed 's/.sh//g'
