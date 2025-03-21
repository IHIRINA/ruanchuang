# ai-swindle-detecter-backend

## 启动说明

启动本程序之前，至少应该启动ai驱动的后端

## 配置文件说明

程序运行后会自动在data目录下生成**setting.json**，这是程序的配置文件，格式为json，对应字段类型及作用如下

| 名      | 类型   | 值                                                          |
| ------ | ---- | ---------------------------------------------------------- |
| port   | str  | 应为本后端的端口                                                   |
| aiport | int  | ai服务器的端口，写在py脚本里，默认为6666                                   |
| debug  | bool | 为true启动debug模式，debug模式下并不访问AI的api，而是直接返回中性作为调用ai的结果。用于测试用。 |

## api返回格式

常规api的返回格式均为json，它们都具有字段code，类型为int，api工作正常时为0，否则错误信息将保存在message字段中。其他字段由不同的api决定，字段code和message将在下面不再解释，而不同api的独特返回字段将会在不同api说明下进行解释。

例如，一个运行正常的api的返回值可能为

```json
{"code":0,"message":""}
```

一个非正常运行的api的返回值可能为

```json
{"code":1,"message":"错误信息"}
```

## 在线api调试

爱来自apipost

[Apipost-在线测试-ai-swindle-detecter-backend](https://doc.apipost.net/docs/3800632ab4ee000?locale=zh-cn)  密码：aipassword

仅支持调试，具体参数说明见下列

## 各类api说明

若说明该类api需要鉴权，则需要在请求头中带有Telephone字段和Password字段，分别是手机号和密码

| 错误码  | 说明           |
| ---- | ------------ |
| -200 | 鉴权错误，即用户密码错误 |

### 杂类

#### 服务器状态面板

类似于实时数据反馈，如相应速度，服务器占用率等等，是个可视化面板

| 地址       | 访问方法 | 参数个数 |
| -------- | ---- | ---- |
| /monitor | GET  | 0    |

### 鉴权类

#### 注册

注册一个新的账号

| 地址             | 访问方法 | 参数个数 |
| -------------- | ---- | ---- |
| /user/register | POST | 2    |

| 参数        | 类型  | 说明                                 |
| --------- | --- | ---------------------------------- |
| telephone | str | 用作类似username的作用，用户名，手机号+密码即可登录一个账户 |
| password  | str | 账户的密码                              |

| 错误码 | 说明                |
| --- | ----------------- |
| -1  | 参数不全              |
| 1   | 该telephone已经被注册过了 |

#### 登录

可用于判断账号密码对不对和判断用户是否存在

| 地址          | 访问方法 | 参数个数 |
| ----------- | ---- | ---- |
| /user/login | POST | 2    |

| 参数        | 类型  | 说明                                 |
| --------- | --- | ---------------------------------- |
| telephone | str | 用作类似username的作用，用户名，手机号+密码即可登录一个账户 |
| password  | str | 账户的密码                              |

| 错误码 | 说明    |
| --- | ----- |
| -1  | 参数不全  |
| 1   | 用户不存在 |
| 2   | 密码错误  |

### AI类

该类api需要鉴权

#### 调用AI判断语句类型

核心api，调用他那个ai模型

| 地址      | 访问方法 | 参数个数 | 返回字段个数 |
| ------- | ---- | ---- | ------ |
| /ai/run | POST | 1    | 2      |

| 参数   | 类型  | 说明                                        |
| ---- | --- | ----------------------------------------- |
| text | str | 用于判断的语句，例如：“不交保证金，不交会费，即可赚取零花钱，最适合宝妈和学生。” |

| 错误码 | 说明             |
| --- | -------------- |
| -1  | 参数不全           |
| 1   | AI服务器出了问题，无法调用 |

| 返回字段    | 类型  | 说明                    |
| ------- | --- | --------------------- |
| type    | str | 得出的类型，比如中性，网络交易及兼职诈骗等 |
| type_id | int | 0-3四个数字，对应上面的四个类型     |

### Data类

该类api需要鉴权

#### 添加一条数据

接收到通知并分类后调用该api存储数据。

| 地址        | 访问方法 | 参数个数 | 返回字段个数 |
| --------- | ---- | ---- | ------ |
| /data/add | POST | 4    | 0      |

| 参数        | 类型  | 说明                            |
| --------- | --- | ----------------------------- |
| package   | str | 接收到的通知的包名，例如：com.example.test |
| type      | str | ai鉴别的类型，例如：中性                 |
| text      | str | 消息的内容                         |
| telephone | str | 用户的电话                         |

| 错误码 | 说明                            |
| --- | ----------------------------- |
| -1  | 参数不全                          |
| 1   | 权限不足，即当前用户没权限往这个telephone上存数据 |

#### 获取用户的所有数据

根据telephone获取用户的**所有**数据

| 地址        | 访问方法 | 参数个数 | 返回字段个数 |
| --------- | ---- | ---- | ------ |
| /data/get | GET  | 1    | 1      |

| 参数        | 类型  | 说明                          |
| --------- | --- | --------------------------- |
| telephone | str | 要获取的用户的电话号码，该参数写在**query**中 |

| 错误码 | 说明                           |
| --- | ---------------------------- |
| -1  | 参数不全                         |
| 1   | 权限不足，即当前用户没权限查这个telephone的数据 |

| 返回字段           | 类型    | 说明                    |
| -------------- | ----- | --------------------- |
| data           | array | 所有数据                  |
| data.package   | str   | 包名，例如com.example.test |
| data.telephone | str   | 用户手机号                 |
| data.text      | str   | 消息内容                  |
| data.type      | str   | ai判定的类型，例如：中性         |

#### 获取用户数据（分页）

根据telephone按照分页获取用户数据

| 地址           | 访问方法 | 参数个数 | 返回字段个数 |
| ------------ | ---- | ---- | ------ |
| /data/cutget | POST | 3    | 2      |

| 参数        | 类型  | 说明                 |
| --------- | --- | ------------------ |
| telephone | str | 要获取的用户的电话号码        |
| cut       | int | 切割的数量，例如5则代表5个数据一页 |
| page      | int | 当前查询的页             |

| 错误码 | 说明                                     |
| --- | -------------------------------------- |
| -1  | 参数不全                                   |
| 1   | 权限不足，即当前用户没权限查这个telephone的数据           |
| 2   | 参数错误，可能是page不在可查询的范围内，page应当>0且<=pages |

| 返回字段           | 类型    | 说明                    |
| -------------- | ----- | --------------------- |
| pages          | int   | 按照cut分割的总页数           |
| data           | array | 当前页的数据                |
| data.package   | str   | 包名，例如com.example.test |
| data.telephone | str   | 用户手机号                 |
| data.text      | str   | 消息内容                  |
| data.type      | str   | ai判定的类型，例如：中性         |

### 关联类

该类api需要鉴权

目前关联类实现方法比较简陋，并无确认这一功能的实现。单方可以添加一个关联。需要注意的是，关联肯定是双向的，即用户1有权限看到用户2的数据的话，用户2必然有权限看到用户1的数据。

#### 添加关联

调用者并**不需要**调用两次add来完成双向添加的功能，只需调用一次即可建立双向的关联。

| 地址        | 访问方法 | 参数个数 | 返回字段个数 |
| --------- | ---- | ---- | ------ |
| /link/add | POST | 2    | 0      |

| 参数         | 类型  | 说明  |
| ---------- | --- | --- |
| telephone1 | str | 用户1 |
| telephone2 | str | 用户2 |

| 错误码 | 说明    |
| --- | ----- |
| 1   | 参数不全  |
| 2   | 关联已存在 |

#### 关联是否存在

| 地址          | 访问方法 | 参数个数 | 返回字段个数 |
| ----------- | ---- | ---- | ------ |
| /link/exist | POST | 2    | 1      |

| 参数         | 类型  | 说明  |
| ---------- | --- | --- |
| telephone1 | str | 用户1 |
| telephone2 | str | 用户2 |

| 错误码 | 说明   |
| --- | ---- |
| 1   | 参数不全 |

| 返回字段  | 类型   | 说明        |
| ----- | ---- | --------- |
| exist | bool | 二者的关联是否存在 |

#### 获取某一用户所有关联

| 地址        | 访问方法 | 参数个数 | 返回字段个数 |
| --------- | ---- | ---- | ------ |
| /link/get | POST | 1    | 1      |

| 参数        | 类型  | 说明  |
| --------- | --- | --- |
| telephone | str | 用户  |

| 错误码 | 说明              |
| --- | --------------- |
| 1   | 参数不全            |
| 2   | 每个用户只能查询自己的所有关联 |

| 返回字段 | 类型    | 说明                                          |
| ---- | ----- | ------------------------------------------- |
| data | array | 关联的数组，每个元素均为str，代表所关联的一个用户，array中的所有元素不会重复。 |