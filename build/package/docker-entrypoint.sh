#!/usr/bin/dumb-init /bin/sh

set -e
set -- "/app/pass" "$@"
set -- su-exec passback:passback "$@"
exec "$@"