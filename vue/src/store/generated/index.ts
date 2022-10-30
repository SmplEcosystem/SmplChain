// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import SmplchainRoles from './smplchain.roles'
import SmplchainSmplchain from './smplchain.smplchain'
import SmplchainSmplusdse from './smplchain.smplusdse'


export default { 
  SmplchainRoles: load(SmplchainRoles, 'smplchain.roles'),
  SmplchainSmplchain: load(SmplchainSmplchain, 'smplchain.smplchain'),
  SmplchainSmplusdse: load(SmplchainSmplusdse, 'smplchain.smplusdse'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}