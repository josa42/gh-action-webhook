#!/bin/bash

version="$1"

if [[ "$version" = "" ]]; then
  exit 1
fi

sed -i.bak "s/\(docker:\/\/josa\/gh-action-webhook\):[^']*/\1:$version/" release/action.yml
rm -f release/action.yml.bak

