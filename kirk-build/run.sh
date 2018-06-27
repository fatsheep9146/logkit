#!/bin/bash

if [[ -z "$LOGKIT_BIND_HOST" ]]; then
   export LOGKIT_BIND_HOST="127.0.0.1:3000"
fi

if [[ -z "$LOGKIT_DEBUG_LEVEL" ]]; then
   export LOGKIT_DEBUG_LEVEL=1
fi

if [[ -z "$LOGKIT_CONFS_PATH" ]]; then
   export LOGKIT_CONFS_PATH="/app/confs"
fi

if [[ -z "$LOGKIT_CONFIG_PATH" ]]; then
   export LOGKIT_CONFIG_PATH=/app/logkit.d
fi

if [[ ! -d "$LOGKIT_CONFIG_PATH" ]]; then
   mkdir "$LOGKIT_CONFIG_PATH"
fi

if [[ ! -f "${LOGKIT_CONFIG_PATH}/logkit.conf" ]]; then
   mv /app/logkit.conf "${LOGKIT_CONFIG_PATH}"
fi

sed -i "s@LOGKIT_CONFS_PATH@${LOGKIT_CONFS_PATH}@g" "${LOGKIT_CONFIG_PATH}/logkit.conf"

sed -i "s/LOGKIT_BIND_HOST/${LOGKIT_BIND_HOST}/g" "${LOGKIT_CONFIG_PATH}/logkit.conf"

sed -i "s/LOGKIT_DEBUG_LEVEL/${LOGKIT_DEBUG_LEVEL}/g" "${LOGKIT_CONFIG_PATH}/logkit.conf"

mkdir -p $LOGKIT_CONFS_PATH

/app/logkit -f "${LOGKIT_CONFIG_PATH}/logkit.conf"
