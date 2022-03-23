local p_dtalk = Proto("dtalk", "dtalk layer protocal");

local f_data = ProtoField.bytes("dtalk.data", "Data")
local f_type = ProtoField.string("dtalk.type", "Type")

p_dtalk.fields = { f_data,f_type }

local protobuf_dissector = Dissector.get("protobuf")

DissectorTable.new('dtalk_dis', 'dtalk_dis', ftypes.UINT32, base.DEC, p_dtalk)

local msgtypes = {
    [1] = 'chat33.comet.AuthMsg',
    [2] = 'chat33.comet.AuthReply',
    [3] = 'chat33.comet.Heartbeat',
    [4] = 'chat33.comet.HeartbeatReply',
    [5] = 'chat33.comet.Disconnect',
    [6] = 'chat33.comet.DisconnectReply',
    [7] = 'dtalk.proto.Proto',
    [8] = 'dtalk.proto.Proto',
    [9] = 'dtalk.proto.Proto',
    [10] = 'dtalk.proto.Proto'
}

local option_name = {
    [0] = 'Undefined',
    [1] = 'Auth',
    [2] = 'AuthReply',
    [3] = 'Heartbeat',
    [4] = 'HeartbeatReply',
    [5] = 'Disconnect',
    [6] = 'DisconnectReply',
    [7] = 'SendMsg',
    [8] = 'SendMsgReply',
    [9] = 'ReceiveMsg',
    [10] = 'ReceiveMsgReply',
    [14] = 'SyncMsgReq',
    [15] = 'SyncMsgReply'
}

--biz proto
local p_dtalk_biz = Proto("dtalk_biz", "dtalk biz layer protocal");

function p_dtalk.dissector(buf, pkt, tree)
    local subtree = tree:add(p_dtalk, buf())

    --Data
    subtree:add(f_data, buf())
    --Type
    local opt_type = tonumber(pkt.private["dtalk_opt_type"])
    subtree:add(f_type, option_name[opt_type])

    pkt.private["pb_msg_type"] = "message," .. msgtypes[opt_type]
    pcall(Dissector.call, protobuf_dissector, buf, pkt, subtree)

  
    local protobuf_field_table = DissectorTable.get("protobuf_field")
    protobuf_field_table:add("dtalk.proto.Proto.body", p_dtalk_biz)
end

function p_dtalk_biz.dissector(buf, pkt, tree)
    for id, name in pairs(DissectorTable.list()) do
        tree:add(id .. ':', name)
    end
    pkt.private["pb_msg_type"] = "message," .. 'dtalk.proto.CommonMsg'
    pcall(Dissector.call, protobuf_dissector, buf, pkt, tree)
end