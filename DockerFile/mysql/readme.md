## docker-compose up -d

## 检查从库的起始状态

  $ show slave status;

## 检查主库的状态

  $ show master status;
  
  记录 主数据库 binary-log 的 文件名称 和 数据同步起始位置。

    File: replicas-mysql-bin.000003
    Position: 154
## 从库配置主库信息

在 从数据库 上运行 主数据库 的相关配置 sql 进行主从关联

```
CHANGE MASTER TO
      MASTER_HOST='mysql-master',
      MASTER_USER='root',
      MASTER_PASSWORD='root',
      MASTER_LOG_FILE='replicas-mysql-bin.000003',
      MASTER_LOG_POS=154;
```
    
## 重新启动 slave 服务

  $ stop slave
  
  $ start slave
  
## 检查从库信息

  show slave status;
