#!/usr/bin/env bash
#
#


curl --data '@query.json' \
-H "Content-Type: application/json" \
-i \
--request POST -sL \
--url 'http://localhost:8088/page'

echo
exit 0