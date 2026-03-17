# Redis 连接配置

## 连接配置字段

| 字段名 | 说明 | 示例 |
|--------|------|------|
| host | 主机地址 | localhost / redis.example.com |
| port | 端口号 | 6379 |
| password | 密码 | - |
| database | 数据库索引 | 0 |
| username | 用户名 (Redis 6+) | default |
| ssl | 启用SSL/TLS | true / false |
| ssl_ca_cert | CA证书路径 | /path/to/ca.crt |
| connection_timeout | 连接超时(秒) | 5 |
| socket_timeout | 套接字超时(秒) | 5 |
| pool_size | 连接池大小 | 10 |
| max_connections | 最大连接数 | 50 |
| retry_on_timeout | 超时重试 | true |
| health_check_interval | 健康检查间隔(秒) | 30 |

## 环境配置模板

### 开发环境 (dev)

```yaml
env: dev
host: localhost
port: 6379
password: ${REDIS_DEV_PASSWORD}
database: 0
ssl: false
connection_timeout: 5
socket_timeout: 5
pool_size: 5
max_connections: 20
retry_on_timeout: true
health_check_interval: 30
```

### 测试环境 (test)

```yaml
env: test
host: test-redis.example.com
port: 6379
password: ${REDIS_TEST_PASSWORD}
database: 0
ssl: true
connection_timeout: 5
socket_timeout: 10
pool_size: 10
max_connections: 50
retry_on_timeout: true
health_check_interval: 30
```

### 生产环境 (prod)

```yaml
env: prod
# Sentinel 模式
sentinels:
  - host: sentinel1.example.com
    port: 26379
  - host: sentinel2.example.com
    port: 26379
  - host: sentinel3.example.com
    port: 26379
master_name: mymaster
password: ${REDIS_PROD_PASSWORD}
database: 0
ssl: true
ssl_ca_cert: /path/to/ca.crt
connection_timeout: 10
socket_timeout: 10
pool_size: 20
max_connections: 100
retry_on_timeout: true
health_check_interval: 30
```

## 连接字符串格式

### 单机模式 (redis-py)
```
redis://[:password@]host:port/database
# 或带SSL
rediss://[:password@]host:port/database
```

### 集群模式
```
redis://host1:port1,host2:port2,host3:port3
```

### Python (redis-py)
```python
import redis

# 单机模式
r = redis.Redis(
    host='localhost',
    port=6379,
    password='password',
    db=0,
    decode_responses=True,
    socket_connect_timeout=5,
    socket_timeout=5,
    max_connections=10
)

# Sentinel 模式
from redis.sentinel import Sentinel
sentinel = Sentinel([
    ('sentinel1.example.com', 26379),
    ('sentinel2.example.com', 26379),
    ('sentinel3.example.com', 26379)
])
master = sentinel.master_for('mymaster', socket_timeout=5)
```

### Node.js (ioredis)
```javascript
const Redis = require('ioredis');

// 单机模式
const redis = new Redis({
    host: 'localhost',
    port: 6379,
    password: 'password',
    db: 0,
    retryStrategy: (times) => Math.min(times * 50, 2000)
});

// Sentinel 模式
const redis = new Redis({
    sentinels: [
        { host: 'sentinel1.example.com', port: 26379 },
        { host: 'sentinel2.example.com', port: 26379 }
    ],
    name: 'mymaster',
    password: 'password'
});
```

## 常用命令

```bash
# 连接 Redis
redis-cli -h host -p port -a password

# 测试连接
redis-cli ping

# 查看信息
redis-cli info

# 监控命令
redis-cli monitor

# 慢查询日志
redis-cli slowlog get 10
```

## 性能建议

1. **连接池**: 生产环境务必使用连接池，避免频繁创建连接
2. **超时设置**: 合理设置 socket_timeout，防止连接挂死
3. **健康检查**: 启用健康检查，及时发现失效连接
4. **Pipeline**: 批量操作使用 pipeline 减少网络往返
