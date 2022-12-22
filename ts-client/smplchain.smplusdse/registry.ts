import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgBurn } from "./types/smplchain/smplusdse/tx";
import { MsgMint } from "./types/smplchain/smplusdse/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/smplchain.smplusdse.MsgBurn", MsgBurn],
    ["/smplchain.smplusdse.MsgMint", MsgMint],
    
];

export { msgTypes }