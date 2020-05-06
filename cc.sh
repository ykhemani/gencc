#!/bin/bash

count=$1
json=$2

echo "################################################################################"
echo "# count is ${count}"
echo "# json is ${json}"
echo ""

go run main.go -count=${count} > ${json}

echo "################################################################################"
echo "# Generated:"
cat ${json}
echo ""

echo "################################################################################"
echo "# Calling Vault Transform Secret Engine"

curl -s \
  --header "X-Vault-Token: ${VAULT_TOKEN}" \
  --request POST \
  --data @${json} \
  ${VAULT_ADDR}/v1/transform/encode/payments | jq -r .
