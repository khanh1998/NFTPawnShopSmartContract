#!/bin/bash

echo "**********************************************" ${MONGO_MONGO1_HOST}
echo "Waiting for startup.."
sleep 10
echo "done"

echo SETUP.sh time now: `date +"%T" `
mongosh --host ${MONGO_MONGO1_HOST}:${MONGO_MONGO1_PORT}  -u ${MONGO_INITDB_ROOT_USERNAME} -p ${MONGO_INITDB_ROOT_PASSWORD} <<EOF
var cfg = {
    "_id": "${MONGO_REPLICA_SET_NAME}",
    "protocolVersion": 1,
    "version": 1,
    "members": [
        {
            "_id": 0,
            "host": "${MONGO_MONGO1_HOST}:${MONGO_MONGO1_PORT}",
            "priority": 2
        },
        {
            "_id": 1,
            "host": "${MONGO_MONGO2_HOST}:${MONGO_MONGO2_PORT}",
            "priority": 0
        },
        {
            "_id": 2,
            "host": "${MONGO_MONGO3_HOST}:${MONGO_MONGO3_PORT}",
            "priority": 0,
        }
    ]
};
rs.initiate(cfg, { force: true });
db.getMongo().setReadPref('primaryPreferred');
rs.status();
EOF