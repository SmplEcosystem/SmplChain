// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import SmplchainSmplchain from './smplchain.smplchain'


export default { 
  SmplchainSmplchain: load(SmplchainSmplchain, 'smplchain.smplchain'),
  
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