import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgMint } from "./types/smplchain/smplusdse/tx";
import { MsgBurn } from "./types/smplchain/smplusdse/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/smplchain.smplusdse.MsgMint", MsgMint],
    ["/smplchain.smplusdse.MsgBurn", MsgBurn],
    
];

export { msgTypes }