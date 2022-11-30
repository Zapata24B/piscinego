#!bin/bash
find . -name "*.sh" | sort -r | rev | cut -d "/" -f1 | rev | sed 's/.sh//g'
# Explications faites pour la dame voilé
# find . -name "*.sh" | sort -r | sed 's/.sh//g' | sed 's/.\///g'
