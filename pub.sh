#!/bin/bash -e

commit_message="$1"
tag="$2"
git add .
git commit -m "$commit_message"
git push origin master
git tag $2
