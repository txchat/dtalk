# 应用方消息子协议    

**协议格式：**

| 参数名     | 必选  | 类型 | 说明       |
| :-----     | :---  | :--- | :---       |
|eventType  | true  | int32 | 事件类型 |
|body  | true  | byte[] | 消息体 |

## 事件类型
| 类型     | 说明  | 
| :-----     | :---  |
| 0 | 普通消息 |
| 1 | 消息回复 |
| 2 | 通知信令 |


## 普通消息
**body描述（encode:proto序列化）**

| 参数名     | 必选  | 类型 | 说明       |
| :-----     | :---  | :--- | :---       |
|channelType  | true  | int32    | 消息通道 |
|logId  | true  | int64    | 消息id |
|msgId  | true | string | 客户端序列 |
|from  | true | string | 发送者 |
|target  | true | string | 接收者 |
|msgType  | true | int32 | 消息类型 |
|msg  | true | binary | 消息体 |
|datetime  | true | uint64 | 事件戳，ms |
|source  | true | Source | 来源，详见如下 |

Source结构：
```
message Source {
    int32 channelType=1;
    SourceUser from=2;
    SourceUser target=3;
}

message SourceUser {
    string id=1;
    string name=2;
}
```

### 消息通道
| 类型     | 说明  | 
| :-----     | :---  |
| 0 | 单聊 |
| 1 | 群聊 |

### 消息类型
| 类型     | 说明  | 
| :-----     | :---  |
| 0 | 系统消息 |
| 1 | 文本消息 |
| 2 | 音频消息 |
| 3 | 图片消息 |
| 4 | 视频消息 |
| 5 | 文件消息 |
| 6 | 卡片消息 |
| 7 | 通知消息 |
| 8 | 合并转发 |

注意：msg 必须可以反序列化为msgType相对应的结构体  
具体协议参考 `pkg/proto/api.proto`

```
message TextMsg {
    string content = 1;
    repeated string mentions = 2;
}

message AudioMsg {
    string mediaUrl = 1;
    int32 time = 2;
}

message ImageMsg {
    string mediaUrl = 1;
    int32 height = 2;
    int32 width = 3;
}

message VideoMsg {
    string mediaUrl = 1;
    int32 time = 2;
    int32 height = 3;
    int32 width = 4;
}

message FileMsg {
    string mediaUrl = 1;
    string name = 2;
    string md5 = 3;
    int64 size = 4;
}

message CardMsg {
    string bank = 1;
    string name = 2;
    string account = 3;
}

message NoticeMsg {
    AlertType type = 1;
    bytes body = 2;
}

message ForwardMsg {
    repeated ForwardItem items = 1;
}

message ForwardItem {
    string avatar=1;
    string name=2;
    int32 msgType=3;
    bytes msg=4;
    uint64 datetime=5;
}
```

#### 通知消息
##### 通知类型
| 类型     | 说明  | 
| :-----     | :---  |
| 0 | 修改群名 |
| 1 | 加群 |
| 2 | 退群 |
| 3 | 踢群 |
| 4 | 删群 |
| 5 | 群禁言模式更改 |
| 6 | 更改禁言名单 |

注意：body必须可以反序列化为AlertType相对应的结构体  
具体协议参考 `pkg/proto/api.proto`

```
message AlertUpdateGroupName {
    int64 group = 1;
    string operator = 2;
    string name = 3;
}

message AlertSignInGroup {
    int64 group = 1;
    string inviter = 2;
    repeated string members = 3;
}

message AlertSignOutGroup {
    int64 group = 1;
    string operator = 2;
}

message AlertKickOutGroup {
    int64 group = 1;
    string operator = 2;
    repeated string members = 3;
}

message AlertDeleteGroup {
    int64 group = 1;
    string operator = 2;
}

message AlertUpdateGroupMuted {
    int64 group = 1;
    string operator = 2;
    MuteType type = 3;
}

message AlertUpdateGroupMemberMutedTime {
    int64 group = 1;
    string operator = 2;
    repeated string members = 3;
}
```

### Receive_Reply协议格式
| 参数名     | 必选  | 类型 | 说明       |
| :-----     | :---  | :--- | :---       |
|eventType  | true  | int32 | 事件类型 |
|body  | true  | byte[] | 消息体 |

当eventType为0时，body必须可以被反序列化为CommonMsgAck
```
message CommonMsgAck {
    int64 logId = 2;
    uint64 datetime = 8;
}
```

## 通知信令
**body描述（encode:proto序列化）**

| 参数名     | 必选  | 类型 | 说明       |
| :-----     | :---  | :--- | :---       |
|action  | true  | ActionType | 信令类型 |
|body  | true | binary | 消息体 |
```
//alert msg define
message NotifyMsg {
    ActionType action = 1;
    bytes body = 2;
}
```
### 信令类型
| 类型     | 说明  | 
| :-----     | :---  |
| 0 | 送达 |
| 1 | 加群 |
| 2 | 退群 |
| 3 | 删群 |
| 20 | 更新加群权限 |
| 21 | 更新群加好友权限 |
| 22 | 更新群禁言类型 |
| 23 | 更新群成员 |
| 24 | 更新禁言列表 |
| 25 | 更新群名 |
| 26 | 更新群头像 |

注意：body必须可反序列化为ActionType对应的结构体。  
具体参考 `pkg/proto/api.proto`

```
message ActionReceived {
    repeated uint64 logs = 1;
}

message ActionSignInGroup {
    repeated string uid = 1;
    int64 group = 2;
    uint64 time = 3;
}

message ActionSignOutGroup {
    repeated string uid = 1;
    int64 group = 2;
    uint64 time = 3;
}

message ActionDeleteGroup {
    int64 group = 1;
    uint64 time = 2;
}
```