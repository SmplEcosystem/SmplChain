
// copy latest  genesis.json 
let genesis = require("./genesis.json")

genesis.app_state.bank.balances.forEach(account => {
    console.log("address",account.address)
    console.log("amount",account.coins[0].amount)
});