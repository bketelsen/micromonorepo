#!/usr/bin/env bash

NET_IF=$(netstat -rn | awk '/^0.0.0.0/ {thif=substr($0,74,10); print thif;} /^default.*UG/ {thif=substr($0,65,10); print thif;}')

NET_IP=$(ifconfig ${NET_IF} | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1')

export NODE1=${NET_IP}

docker volume create --name etcd-data
export ETCD_DATA_DIR="etcd-data"

ETCD_VERSION=v3.5.21
REGISTRY=quay.io/coreos/etcd
# available from v3.2.5
REGISTRY=gcr.io/etcd-development/etcd

docker run \
    -p 2379:2379 \
    -p 2380:2380 \
    --volume=${ETCD_DATA_DIR}:/etcd-data \
    --name etcd ${REGISTRY}:${ETCD_VERSION} \
    /usr/local/bin/etcd \
    --data-dir=/etcd-data --name node1 \
    --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://0.0.0.0:2380 \
    --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://0.0.0.0:2379 \
    --initial-cluster node1=http://${NODE1}:2380
