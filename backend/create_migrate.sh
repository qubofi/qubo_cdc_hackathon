#!/bin/bash
# run this script with first parameter as migration file name (snake case)
CURRENT_UNIX_TIMESTAMP=$EPOCHSECONDS
MIGRATION_NAME=$1
FINAL_MIGRATION_NAME="${CURRENT_UNIX_TIMESTAMP}_${MIGRATION_NAME}"
echo "Creating migration $MIGRATION_NAME"
echo  > "./infrastructure/database/migrations/${FINAL_MIGRATION_NAME}.up.sql"
echo  > "./infrastructure/database/migrations/${FINAL_MIGRATION_NAME}.down.sql"
