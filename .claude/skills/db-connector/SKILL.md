---
name: db-connector
description: 数据库连接信息管理工具，用于存储和查询 MySQL、MongoDB、Redis、ClickHouse 等数据库的连接配置。支持开发、测试、生产多环境管理，提供连接字符串生成和代码示例。当用户需要查询数据库连接信息、生成连接配置或了解数据库连接参数时使用此 skill。
---

# 数据库连接信息管理

## 功能概述

此 skill 帮助管理和查询各种数据库的连接配置信息，支持：

- **多种数据库**: MySQL、MongoDB、Redis、ClickHouse
- **多环境管理**: 开发(dev)、测试(test)、生产(prod)环境配置
- **连接字符串生成**: 自动生成各语言的连接代码
- **配置模板**: 提供标准化的连接配置字段说明

## 使用方式

### 1. 查询数据库连接配置

当用户询问某个数据库的连接信息时：

1. 根据数据库类型读取对应的参考文件：
   - MySQL: [references/mysql.md](references/mysql.md)
   - MongoDB: [references/mongodb.md](references/mongodb.md)
   - Redis: [references/redis.md](references/redis.md)
   - ClickHouse: [references/clickhouse.md](references/clickhouse.md)

2. 提取对应环境的配置模板
3. 根据需要提供连接字符串或代码示例

### 2. 连接配置字段说明

所有数据库连接配置都包含以下核心字段：

| 字段 | 说明 |
|------|------|
| host | 数据库主机地址 |
| port | 端口号 |
| username | 用户名 |
| password | 密码（建议使用环境变量）|
| database | 数据库名 |
| ssl/ssl_mode/secure | SSL/TLS 配置 |
| timeout | 连接超时设置 |
| pool_size | 连接池大小 |

### 3. 环境区分

- **dev**: 本地开发环境，通常禁用 SSL，连接池较小
- **test**: 测试环境，启用 SSL，中等连接池
- **prod**: 生产环境，高可用配置（集群/副本集），最大连接池，强制 SSL

## 参考文件索引

| 数据库 | 文件路径 | 内容 |
|--------|----------|------|
| MySQL | references/mysql.md | 连接字段、JDBC/各语言示例、环境配置 |
| MongoDB | references/mongodb.md | 单机/副本集配置、连接字符串、pymongo示例 |
| Redis | references/redis.md | 单机/Sentinel模式、redis-py/ioredis示例 |
| ClickHouse | references/clickhouse.md | HTTP/Native协议、Python客户端示例、集群配置 |

## 最佳实践

1. **密码管理**: 始终使用环境变量存储密码，不要硬编码
2. **连接池**: 生产环境合理设置连接池大小，避免资源耗尽
3. **超时设置**: 根据网络环境设置合理的连接和读取超时
4. **SSL/TLS**: 生产环境必须启用 SSL，验证证书
5. **环境隔离**: 严格区分不同环境的配置，避免误连生产数据库

## 安全提醒

- 不要将包含真实密码的配置文件提交到代码仓库
- 使用 `.env` 文件或密钥管理服务存储敏感信息
- 生产环境限制数据库访问IP白名单
- 定期轮换数据库密码
