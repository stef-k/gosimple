#!/usr/bin/env bash
# Takeover the project by renaming all instances of Go Simple
# to current directory name and by removing the current .git directory.
# Example:
# 1. Clone to myproject
# 2. Run makemine.sh and all packages, files and folders will be renamed as myproject

# remove .git
# rm -rf .git
# get current directory
name=${PWD##*/}

# find and replace Go Simple  (case insensitive) with current directory name
find ./ -type f s-exec sed -i -e "s/makemine/$name/gI"  {} \;
