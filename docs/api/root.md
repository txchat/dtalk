# 总览
## 约定
- 数据类型`datetime`：时间戳，毫秒(ms)
- 所有 `http` 接口请求都是 `POST` 方法， 数据请求和返回结果 都是 `JSON` 格式
- Header: 接口里需要的 
  - `FZM-SIGNATURE`（地址）：格式`<signature>#<message>#<publicKey>`; signature格式：`<datetime>*[random string]`
  - `FZM-DEVICE`（设备类型）：枚举 iOS、Android、PC
  - `FZM-UUID`（设备mac）
  - `FZM-DEVICE-NAME`（设备名称）
  - `FZM-VERSION` （应用版本号）
- 返回结果的 `result` 为 `0` 代表成功，`data` 里面是具体数据，`result` 不为 `0` 代表失败，`message` 为失败信息