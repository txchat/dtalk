# 以最新swagger为准
[http://172.16.101.107:8888/group/swagger/index.html](http://172.16.101.107:8888/group/swagger/index.html)

## 群服务


测试地址 172.16.101.107:8888/group/app



## 已完成接口

| 路由                                  | 说明       |
| ------------------------------------- | ---------- |
| `POST` URL: /app/create-group         | 创建群     |
| `POST` URL:/app /invite-group-members | 邀请新群友 |
| `POST` URL: /app/group-info           | 群信息     |
| `POST` URL: /app/group-list           | 群列表     |
| `POST` URL: /app/group-member-list    | 群成员列表 |
| `POST`URL:/app/group-member-info      | 群成员信息 |
| `POST` URL:/app/group-remove          | 踢群成员   |
| `POST` URL:/app/group-exit            | 退出群     |
| `POST` URL:/app/group-disband         | 解散群     |




### 创建群+

`POST` URL: /app/create-group



**Herder**

`FZM-SIGNATURE` = token



**请求参数：**

| 参数名    | 必选  | 类型     | 说明       |
| --------- | ----- | -------- | ---------- |
| name      | false | string   | 群名称     |
| avatar    | false | string   | 群头像 url |
| introduce | false | string   | 群简介     |
| memberIds | false | []string | 新群员 id  |



**请求参数示例**

```json
{
    "name": "test-group-1",
    "avatar": "",
    "memberIds": [
       "member-1", "member-2"
    ]
}
```



**返回参数(data)：**

| 参数名     | 类型         | 说明                                                         |
| ---------- | ------------ | ------------------------------------------------------------ |
| id         | int          | 群id                                                         |
| markId     | string       | 群短 id(暂时没用, 后面可以供搜索加群使用)                    |
| name       | string       | 群名称                                                       |
| avatar     | string       | 群头像 url                                                   |
| introduce  | string       | 群简介                                                       |
| owner      | MemberInfo   | 群主信息                                                     |
| members    | []MemberInfo | 群成员信息list                                               |
| memberNum  | int          | 群总人数                                                     |
| maximum    | int          | 群成员人数上限                                               |
| status     | int          | 群状态，0=正常 1=封禁 2=解散                                 |
| createTime | int          | 群创建时间                                                   |
| joinType   | int          | 加群方式，0=允许任何方式加群，1=群成员邀请加群，2=群主和管理员邀请加群 |
| muteType   | int          | 禁言， 0=所有人可以发言， 1=群主和管理员可以发言             |



**MemberInfo 参数类型**

| 参数名     | 类型   | 说明                                       |
| ---------- | ------ | ------------------------------------------ |
| memberId   | string | 用户id                                     |
| memberName | string | 用户群昵称                                 |
| memberType | int    | 用户角色，2=群主，1=管理员，0=群员，3=退群 |



**返回参数示例：**

```json
{
    "result": 0,
    "message": "",
    "data": {
        "id": 127043701116506112,
        "markId": "00351854",
        "name": "test-group-1",
        "avatar": "",
        "introduce": "",
        "owner": {
            "memberId": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
            "memberName": "",
            "memberType": 2
        },
        "members": [
            {
                "memberId": "member-1",
                "memberName": "",
                "memberType": 0
            },
            {
                "memberId": "member-2",
                "memberName": "",
                "memberType": 0
            }
        ],
        "memberNum": 3,
        "maximum": 200,
        "status": 0,
        "createTime": 1621230378707,
        "joinType": 0,
        "muteType": 0
    }
}
```



### 邀请新群友+

`POST` URL: /app/invite-group-members



**Herder**

`FZM-SIGNATURE` = token



**请求参数：**

| 参数名       | 必选 | 类型     | 说明       |
| ------------ | ---- | -------- | ---------- |
| id           | true | int64    | 群ID       |
| newMemberIds | true | []string | 被邀请人ID |



**请求参数示例**

```json
{
    "id": 127043701116506112,
    "newMemberIds": [
        "member-5",
        "member-6"
    ]
}
```



**返回参数(data)：**

| 参数名     | 类型         | 说明              |
| ---------- | ------------ | ----------------- |
| id         | int          | 群id              |
| memberNum  | int          | 群总人数          |
| inviter    | MemberInfo   | 邀请人信息        |
| newMembers | []MemberInfo | 被邀请人信息 list |



**MemberInfo 参数类型**

| 参数名     | 类型   | 说明                                       |
| ---------- | ------ | ------------------------------------------ |
| memberId   | string | 用户id                                     |
| memberName | string | 用户群昵称                                 |
| memberType | int    | 用户角色，2=群主，1=管理员，0=群员，3=退群 |



**返回参数示例：**

```json
{
    "result": 0,
    "message": "",
    "data": {
        "id": 125290793882619904,
        "memberNum": 5,
        "inviter": {
            "memberId": "chenhongyu",
            "memberName": "",
            "memberType": 0
        },
        "newMembers": [
            {
                "memberId": "member-4",
                "memberName": "",
                "memberType": 2
            }
        ]
    }
}
```



### 群信息+

`POST` URL: /app/group-info



**Herder**

`FZM-SIGNATURE` = token



**请求参数：**

| 参数名 | 必选 | 类型 | 说明 |
| ------ | ---- | ---- | ---- |
| id     | true | int  | 群ID |



**请求参数示例**

```json
{
    "id":127082931377147904
}
```



**返回参数(data)：**

| 参数名     | 类型         | 说明                                                         |
| ---------- | ------------ | ------------------------------------------------------------ |
| id         | int          | 群id                                                         |
| markId     | string       | 群短 id(暂时没用, 后面可以供搜索加群使用)                    |
| name       | string       | 群名称                                                       |
| avatar     | string       | 群头像 url                                                   |
| introduce  | string       | 群简介                                                       |
| owner      | MemberInfo   | 群主信息                                                     |
| person     | MemberInfo   | 本人信息                                                     |
| members    | []MemberInfo | 所有群成员信息list                                           |
| memberNum  | int          | 群总人数                                                     |
| maximum    | int          | 群成员人数上限                                               |
| status     | int          | 群状态，0=正常 1=封禁 2=解散                                 |
| createTime | int          | 群创建时间                                                   |
| joinType   | int          | 加群方式，0=允许任何方式加群，1=群成员邀请加群，2=群主和管理员邀请加群 |
| muteType   | int          | 禁言， 0=所有人可以发言， 1=群主和管理员可以发言             |



**MemberInfo 参数类型**

| 参数名     | 类型   | 说明                                       |
| ---------- | ------ | ------------------------------------------ |
| memberId   | string | 用户id                                     |
| memberName | string | 用户群昵称                                 |
| memberType | int    | 用户角色，0=群主，1=管理员，2=群员，3=退群 |



**返回参数示例：**

```json
{
    "result": 0,
    "message": "",
    "data": {
        "id": 127043701116506112,
        "markId": "",
        "name": "",
        "avatar": "",
        "introduce": "",
        "owner": {
            "memberId": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
            "memberName": "",
            "memberType": 2
        },
        "members": [
            {
                "memberId": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
                "memberName": "",
                "memberType": 2
            },
            {
                "memberId": "member-1",
                "memberName": "",
                "memberType": 0
            },
            {
                "memberId": "member-2",
                "memberName": "",
                "memberType": 0
            },
            {
                "memberId": "member-5",
                "memberName": "",
                "memberType": 0
            },
            {
                "memberId": "member-6",
                "memberName": "",
                "memberType": 0
            }
        ],
        "memberNum": 5,
        "maximum": 200,
        "status": 0,
        "createTime": 1621230378707,
        "joinType": 0,
        "muteType": 0
    }
}
```





### 群列表+

`GET` | `POST` URL: /app/group-list



**Herder**

`FZM-SIGNATURE` = token



**请求参数：**

| 参数名 | 必选 | 类型 | 说明 |
| ------ | ---- | ---- | ---- |
|        |      |      |      |



**请求参数示例**

```json

```



**返回参数(data)：**

| 参数名 | 类型        | 说明        |
| ------ | ----------- | ----------- |
| groups | []GroupInfo | 群信息 list |



**GroupInfo：**

| 参数名     | 类型       | 说明                                                         |
| ---------- | ---------- | ------------------------------------------------------------ |
| id         | int        | 群id                                                         |
| markId     | string     | 群短 id(暂时没用, 后面可以供搜索加群使用)                    |
| name       | string     | 群名称                                                       |
| avatar     | string     | 群头像 url                                                   |
| introduce  | string     | 群简介                                                       |
| owner      | MemberInfo | 群主信息                                                     |
| memberNum  | int        | 群总人数                                                     |
| maximum    | int        | 群成员人数上限                                               |
| status     | int        | 群状态，0=正常 1=封禁 2=解散                                 |
| createTime | int        | 群创建时间                                                   |
| joinType   | int        | 加群方式，0=允许任何方式加群，1=群成员邀请加群，2=群主和管理员邀请加群 |
| muteType   | int        | 禁言， 0=所有人可以发言， 1=群主和管理员可以发言             |




**返回参数示例：**

```json
{
    "result": 0,
    "message": "",
    "data": {
        "groups": [
            {
                "id": 127043701116506112,
                "markId": "00351854",
                "name": "test-group-1",
                "avatar": "",
                "introduce": "",
                "owner": {
                    "memberId": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
                    "memberName": "",
                    "memberType": 2
                },
                "memberNum": 5,
                "maximum": 200,
                "status": 0,
                "createTime": 1621230378707,
                "joinType": 0,
                "muteType": 0
            },
            {
                "id": 127067012814868480,
                "markId": "62025607",
                "name": "test-group-1",
                "avatar": "",
                "introduce": "",
                "owner": {
                    "memberId": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
                    "memberName": "",
                    "memberType": 2
                },
                "memberNum": 5,
                "maximum": 200,
                "status": 0,
                "createTime": 1621235936650,
                "joinType": 0,
                "muteType": 0
            },
            {
                "id": 127082931377147904,
                "markId": "88951481",
                "name": "test-group-1",
                "avatar": "",
                "introduce": "",
                "owner": {
                    "memberId": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
                    "memberName": "",
                    "memberType": 2
                },
                "memberNum": 11,
                "maximum": 200,
                "status": 0,
                "createTime": 1621239731935,
                "joinType": 0,
                "muteType": 0
            },
            {
                "id": 127100214044528640,
                "markId": "43833969",
                "name": "test-group-1",
                "avatar": "",
                "introduce": "",
                "owner": {
                    "memberId": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
                    "memberName": "",
                    "memberType": 2
                },
                "memberNum": 3,
                "maximum": 200,
                "status": 0,
                "createTime": 1621243852439,
                "joinType": 0,
                "muteType": 0
            },
            {
                "id": 127485592593240064,
                "markId": "00865442",
                "name": "test-group-1",
                "avatar": "",
                "introduce": "",
                "owner": {
                    "memberId": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
                    "memberName": "",
                    "memberType": 2
                },
                "memberNum": 9,
                "maximum": 200,
                "status": 0,
                "createTime": 1621335733857,
                "joinType": 0,
                "muteType": 0
            }
        ]
    }
}
```





### 群成员列表+

`POST` URL: /app/group-member-list



**Herder**

`FZM-SIGNATURE` = token



**请求参数：**

| 参数名 | 必选 | 类型  | 说明 |
| ------ | ---- | ----- | ---- |
| id     | true | int64 | 群ID |



**请求参数示例**

```json
{
    "id": 127082931377147904
}
```



**返回参数(data)：**

| 参数名  | 类型         | 说明            |
| ------- | ------------ | --------------- |
| id      | int          | 群id            |
| members | []MemberInfo | 全部群成员 list |



**MemberInfo 参数类型**

| 参数名     | 类型   | 说明                                       |
| ---------- | ------ | ------------------------------------------ |
| memberId   | string | 用户id                                     |
| memberName | string | 用户群昵称                                 |
| memberType | int    | 用户角色，0=群主，1=管理员，2=群员，3=退群 |



**返回参数示例：**

```json
{
    "result": 0,
    "message": "",
    "data": {
        "id": 127043701116506112,
        "members": [
            {
                "memberId": "1FKxgaEh5fuSm7a35BfUnKYAmradowpiTR",
                "memberName": "",
                "memberType": 2
            },
            {
                "memberId": "member-1",
                "memberName": "",
                "memberType": 0
            },
            {
                "memberId": "member-2",
                "memberName": "",
                "memberType": 0
            },
            {
                "memberId": "member-5",
                "memberName": "",
                "memberType": 0
            },
            {
                "memberId": "member-6",
                "memberName": "",
                "memberType": 0
            }
        ]
    }
}
```





### 群成员信息+

URL: /app/group-member-info

`GET`



**Herder**

`FZM-SIGNATURE` = token



**请求参数：**

| 参数名   | 必选 | 类型   | 说明      |
| -------- | ---- | ------ | --------- |
| id       | true | int    | 群ID      |
| memberId | true | string | 群成员 ID |



**请求参数示例**

```json
{
    "id": 125290793882619904,
    "memberId": "123"
}
```



**返回参数(data)：**

| 参数名     | 类型       | 说明   |
| ---------- | ---------- | ------ |
| id         | int        | 群id   |
| newMembers | MemberInfo | 群成员 |



**MemberInfo 参数类型**

| 参数名     | 类型   | 说明                                       |
| ---------- | ------ | ------------------------------------------ |
| memberId   | string | 用户id                                     |
| memberName | string | 用户群昵称                                 |
| memberType | int    | 用户角色，0=群主，1=管理员，2=群员，3=退群 |



**返回参数示例：**

```json
{
    "result": 0,
    "message": "",
    "data": {
        "memberId": "member-1",
        "memberName": "",
        "memberType": 0
    }
}
```




### 踢人+

`POST` URL: /app/group-remove



### 退群+

`POST` URL:/app/group-exit



### 解散群

URL: /disband-group

`PUT`




### 更新群头像

URL: /update-group-avatar

`PUT`



### 更新群名称

URL: /update-group-name

`PUT`



### 更新个人群昵称(自己改自己的, 群主和管理员改所有人)

URL: /update-group-member-name

`PUT`



### 更新群简介

URL: /update-group-introduce

`PUT`



### 更新加群权限设置

URL: /update-group-join-type

`PUT`





### 更新群成员类型

URL: /update-group-member-type

`PUT`



### 转移群给群成员

URL: /update-group-owner

`PUT`







### 更新群状态

URL: /update-group-status

`PUT`



### 更新群成员上限

URL: /update-group-maximum

`PUT`