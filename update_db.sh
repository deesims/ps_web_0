#!bin/bash

SCRIPT=`realpath $0`
SCRIPTPATH=`dirname $SCRIPT`

sudo -u postgres -H sh -c "psql -d postgres -a -f $SCRIPTPATH/database_dump.sql"