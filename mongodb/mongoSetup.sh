#!/bin/bash

MONGODB1=mongo
MONGODB2=mongo1

echo "**********************************************" ${MONGODB1}
echo "Waiting for startup.."
sleep 30
echo "done"

echo SETUP.sh time now: `date +"%T" `
mongo --host ${MONGODB1}:27018 -u khanh -p handsome <<EOF
var cfg = {
    "_id": "rsmongo",
    "protocolVersion": 1,
    "version": 1,
    "members": [
        {
            "_id": 0,
            "host": "${MONGODB1}:27018",
            "priority": 2
        },
        {
            "_id": 1,
            "host": "${MONGODB2}:27019",
            "priority": 0
        },
    ]
};
rs.initiate(cfg, { force: true });
rs.secondaryOk();
db.getMongo().setReadPref('primary');
rs.status();
EOF