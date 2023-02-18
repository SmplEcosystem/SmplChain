#!/usr/bin/env bash

if [ -d "$HOME/.SmplChain" ]; then
  echo "$HOME/.SmplChain exists, exiting"
  exit 1
fi

if [ "" == "$1" ]; then
  echo "Please pass a node moniker"
  exit 1
fi

if [ "" == "$2" ]; then
  echo "Please pass an admin account"
  exit 1
fi

file=smpl-chaind-state.txt

denomMetadata=$(<denom-metadata.json)

SmplChaind init $1

chainName="cosmos:smpl-chain-testing"

genesisFile="$HOME/.SmplChain/config/genesis.json"
sed -i '/\[api\]/,+3 s/enable = false/enable = true/' "$HOME"/.SmplChain/config/app.toml
sed -i 's/"os"/"test"/' "$HOME"/.SmplChain/config/client.toml
sed -i "s/SmplChain/$chainName/" "$HOME"/.SmplChain/config/client.toml
jq ".chain_id = \"$chainName\"" "$genesisFile" >temp.json && mv temp.json "$genesisFile"
jq '.chain_id = "cosmos:smpl-chain-testing"' "$genesisFile" >temp.json && mv temp.json "$genesisFile"
jq '.app_state.mint.minter.inflation = "0"' "$genesisFile" > temp.json && mv temp.json "$genesisFile"
jq '.app_state.mint.params.mint_denom = "garden"' "$genesisFile" > temp.json && mv temp.json "$genesisFile"
jq '.app_state.mint.params.inflation_rate_change = "0"' "$genesisFile" > temp.json && mv temp.json "$genesisFile"
jq '.app_state.mint.params.inflation_max = "0"' "$genesisFile" > temp.json && mv temp.json "$genesisFile"
jq '.app_state.mint.params.inflation_min = "0"' "$genesisFile" > temp.json && mv temp.json "$genesisFile"

jq '.app_state.crisis.constant_fee.denom = "garden"' "$genesisFile" > temp.json && mv temp.json "$genesisFile"
jq '.app_state.staking.params.bond_denom = "garden"' "$genesisFile" > temp.json && mv temp.json "$genesisFile"
jq '.app_state.gov.deposit_params.min_deposit[0].denom = "garden"' "$genesisFile" > temp.json && mv temp.json "$genesisFile"
jq ".app_state.roles.adminaccount = \"$2\"" "$genesisFile" > temp.json && mv temp.json "$genesisFile"

jq ".app_state.bank.denom_metadata = $denomMetadata" "$genesisFile" > temp.json && mv temp.json "$genesisFile"

while read -r i; do
  echo $i
  j=$(sed 's/\s.*$//' <<<"$i")
  r=$(sed 's/^\w* //' <<<"$i")
  SmplChaind add-genesis-account "$j" "$r"
done <$file

cp -a keyring-test "$HOME"/.SmplChain/

SmplChaind gentx validator2 100000000000garden

SmplChaind collect-gentxs