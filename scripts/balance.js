
// copy latest  genesis.json

let file = 'genesis.json'
if (process.argv[2]) file = process.argv[2];

let genesis = require(`${file}`)
let smplTotal = 0;

genesis.app_state.bank.balances.forEach(account => {
    let output = `${account.address}`;
    let authAccount = genesis.app_state.auth.accounts.find(a => a.address === account.address);
    if(!authAccount || authAccount['@type'] !== '/cosmos.auth.v1beta1.BaseAccount') return;
    const coins = [];
    account.coins.forEach(coin => {
        let amount = coin.amount;
        let denom = coin.denom;

        if (denom === 'smpl') {
            smplTotal += parseInt(amount);
            amount += "000000000"
            denom = 'garden'
        } else if (denom === 'usdse') {
            amount += "000000"
            denom = 'busdse'
        } else {
            console.log('OTHER DENOM')
        }

        coins.push(`${amount}${denom}`);
    })
    console.log(`${coins.join('\t')}`)
    // console.log("address",account.address)
    // console.log("amount",account.coins[0].amount)
});

console.log('total smpl', smplTotal)