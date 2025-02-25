# 更新日志

## dev

### Fix

- 修复 job state 变化时仍然更新了 job progress 的问题，对之前的逻辑无影响，主要用于支持 partial sync (selectdb/ccr-syncer#124)
- 修复 get_lag 接口中不含 lag 的问题 (selectdb/ccr-syncer#126)
- 修复下游 restore 时未清理 orphan tables/partitions 的问题 (selectdb/ccr-syncer#128)
- **修复下游删表后重做 snapshot 时 dest meta cache 过期的问题 (selectdb/ccr-syncer#132)**

### Feature

- 支持 partial sync，减少需要同步的数据量 (selectdb/ccr-syncer#125)
- 添加参数 `allowTableExists`，允许在下游 table 存在时，仍然创建 ccr job（如果 schema 不一致，会自动删表重建）(selectdb/ccr-syncer#136)

### Improve

- 如果下游表的 schema 不一致，则将表移动到 RecycleBin 中（之前是强制删除）(selectdb/ccr-syncer#137)

## 2.0.14/2.1.5

### Fix

- 过滤已经删除的 partitions，避免 full sync，需要 doris 2.0.14/2.1.5 (selectdb/ccr-syncer#117)
- 过滤已经删除的 tables，避免 full sync (selectdb/ccr-syncer#123)
- 兼容 doris 3.0 alternative json name，doris 3.0 必须使用该版本的 CCR syncer (selectdb/ccr-syncer#121)
- 修复 list jobs 接口在高可用环境下不可用的问题 (selectdb/ccr-syncer#120)

## 2.0.11

对应 doris 2.0.11。

### Feature

- 支持以 postgresql 作为 ccr-syncer 的元数据库 (selectdb/ccr-syncer#77)
- 支持 insert overwrite 相关操作 (selectdb/ccr-syncer#97,selectdb/ccr-syncer#99)

### Fix

- 修复 drop partition 后因找不到 partition id 而无法继续同步的问题 (selectdb/ccr-syncer#82)
- 修复高可用模式下接口无法 redirect 的问题 (selectdb/ccr-syncer#81)
- 修复 binlog 可能因同步失败而丢失的问题 (selectdb/ccr-syncer#86,selectdb/ccr-syncer#91)
- 修改 connect 和 rpc 超时时间默认值，connect 默认 10s，rpc 默认 30s (selectdb/ccr-syncer#94,selectdb/ccr-syncer#95)
- 修复 view 和 materialized view 使用造成空指针问题 (selectdb/ccr-syncer#100)
- 修复 add partition sql 错误的问题 (selectdb/ccr-syncer#99)


## 2.1.3/2.0.3.10

### Fix

- 修复因与上下游 FE 网络中断而触发 full sync 的问题

### Feature

- 增加 `/job_progress` 接口用于获取 JOB 进度
- 增加 `/job_details` 接口用于获取 JOB 信息
- 保留 job 状态变更的各个时间点，并在 `/job_progress` 接口中展示

### Fix

- 修复若干 keywords 没有 escape 的问题

## 2.0.3.9

配合 doris 2.0.9 版本

### Feature

- 添加选项以启动 pprof server
- 允许配置 rpc 合 connection 超时

### Fix

- restore 每次重试时使用不同的 label 名
- update table 失败时（目标表不存在）会触发快照同步
- 修复同步 sql 中包含关键字的问题
- 如果恢复时碰到表 schema 发生变化，会先删表再重试恢复

## 0.5

### 支持高可用
- 现在可以部署多个Syncer节点来保证CCR功能的高可用。
- db是Syncer集群划分的依据，同一个集群下的Syncer共用一个db。
- Syncer集群采用对称设计，每个Syncer都会相对独立的执行被分配到的job。在某个Syncer节点down掉后，它的jobs会依据负载均衡算法被分给其他Syncer节点。

## 0.4
* 增加 enable_db_binlog.sh 方便用户对整库开启binlog

## 0.3

### LOG

- 更新日志格式，提高日志可读性，现在日志的格式如下，其中hook只会在 `log_level > info`的时候打印：

  ```bash
  #        time         level        msg                  hooks
  [2023-07-18 16:30:18] TRACE This is trace type. ccrName=xxx line=xxx
  [2023-07-18 16:30:18] DEBUG This is debug type. ccrName=xxx line=xxx
  [2023-07-18 16:30:18]  INFO This is info type. ccrName=xxx line=xxx
  [2023-07-18 16:30:18]  WARN This is warn type. ccrName=xxx line=xxx
  [2023-07-18 16:30:18] ERROR This is error type. ccrName=xxx line=xxx
  [2023-07-18 16:30:18] FATAL This is fatal type. ccrName=xxx line=xxx
  ```
- 现在可以指定log的等级和log文件的路径  
  `--log_level <level>`：  
  level可以是trace、debug、info、warn、error、fatal，log的数量依次递减。默认值为 `info`  
  
  `--log_dir </PATH/TO/LOG/FILE>`：  
  log文件路径包括路径+文件名，如：/var/myfile.log，默认值为 `log/ccr-syncer.log`  
  
  例：  

  ```bash
  sh start_syncer.sh --daemon --log_level trace --log_dir /PATH/TO/LOG/FILE
  ```
- 非守护进程状态下会在日志打印到终端的同时利用tee将其保存在 `log_dir`中
- 在日志中屏蔽了用户的敏感信息

### BD

- 现在可以指定syncer持久化DB的文件路径  
  `--db_dir </PATH/TO/DB/FILE>`：  
  DB文件路径包括路径+文件名，如：/var/myccr.db，默认值为 `db/ccr.db`  
  
  例：  

  ```bash
  sh start_syncer.sh --daemon --db_dir /PATH/TO/DB/FILE
  ```
