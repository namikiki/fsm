# FSM

一个使用Golang 编写的多用户多端实时`文件同步`项目

服务端 [github.com/namikiki/fsm](github.com/namikiki/fsm)

客户端 [github.com/namikiki/fsm_client](github.com/namikiki/fsm_client)


## 功能
- 用户管理: 账号注册 登录 更新密码
- 同步任务: 用户可创建同步任务 创建 暂停 继续以及删除
- 实时同步: 当客户端文件产生变化时，会被立即同步至云端 
- 多平台实时同步: 客户端支持Windows Linux MacOS， 同一账户的多个客户端在线时，文件变化实时同步至所有客户端
- 安全存储: 使用 Minio 存储同步文件  
- 差异同步检查: 当其中一个客户端离线后，此客户端文件产生的差异和其他在线客户端所产生的差异会在重新上线后进行差异检查和同步
- 同步过滤: 用户可以过滤不需要同步的文件或者文件夹

## 技术栈与运行需求
- Go 1.19
- Gin
- Redis 6.2
- MySQL 8.3
- GORM
- Minio

## 服务端编译与启动
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/yourproject.git
2. 编译
   ```sh
   go build .
3. 启动配置
   ```toml
   [DataBase]
   DSN = "databse dsn" # 数据库连接

   [Redis]
   Address = ""  # Redis 连接
   Password = "" # 没有密码则为空

   [Minio]
   Endpoint        = "play.min.io" # Minio #连接地址
   AccessKeyID     = "Q3AM3UQ867SPQQA43P2F"  
   SecretAccessKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
   UseSSL          = true #是否启用SSL
   
   [Develop]
   DevMod = true  # 是否启用开发模式
