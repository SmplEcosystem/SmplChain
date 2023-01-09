
// copy latest  genesis.json

let file = 'genesis.json'
if (process.argv[2]) file = process.argv[2];

let genesis = require(`${file}`)

genesis.app_state.bank.balances.forEach(account => {
    let output = `${account.address} `;
    account.coins.forEach(coin => {
        let amount = coin.amount;
        let denom = coin.denom;
        if (denom === 'smpl') {
            amount += "000000000"
            denom = 'garden'
        } else if (denom === 'usdse') {
            amount += "000000"
            denom = 'busdse'
        }
        output += `${amount}${denom}`;
        console.log(output)
    })
    // console.log("address",account.address)
    // console.log("amount",account.coins[0].amount)
});