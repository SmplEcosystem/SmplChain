import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgAdd } from "./types/smplchain/roles/tx";
import { MsgRemove } from "./types/smplchain/roles/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/smplchain.roles.MsgAdd", MsgAdd],
    ["/smplchain.roles.MsgRemove", MsgRemove],
    
];

export { msgTypes }