# ClickHouse 连接配置

## 连接配置字段

| 字段名 | 说明 | 示例 |
|--------|------|------|
| host | 主机地址 | localhost / clickhouse.example.com |
| port | HTTP端口 | 8123 |
| native_port | Native协议端口 | 9000 |
| username | 用户名 | default |
| password | 密码 | - |
| database | 数据库名 | default |
| secure | 启用HTTPS | true / false |
| verify | 验证SSL证书 | true / false |
| timeout | 超时(秒) | 30 |
| compress | 启用压缩 | true |
| compression_codec | 压缩编解码 | lz4 / zstd |
| pool_size | 连接池大小 | 10 |
| max_execution_time | 最大执行时间(秒) | 60 |
| settings | 会话设置 | readonly=1 |

## 环境配置模板

### 开发环境 (dev)

```yaml
env: dev
host: localhost
port: 8123
native_port: 9000
username: default
password: ${CLICKHOUSE_DEV_PASSWORD}
database: myapp_dev
secure: false
verify: false
timeout: 30
compress: true
compression_codec: lz4
pool_size: 5
max_execution_time: 60
```

### 测试环境 (test)

```yaml
env: test
host: test-clickhouse.example.com
port: 8123
native_port: 9000
username: test_user
password: ${CLICKHOUSE_TEST_PASSWORD}
database: myapp_test
secure: true
verify: true
timeout: 60
compress: true
compression_codec: lz4
pool_size: 10
max_execution_time: 120
```

### 生产环境 (prod)

```yaml
env: prod
# 集群配置
hosts:
  - host: ch-node1.example.com
    port: 8123
  - host: ch-node2.example.com
    port: 8123
  - host: ch-node3.example.com
    port: 8123
username: prod_user
password: ${CLICKHOUSE_PROD_PASSWORD}
database: myapp_prod
secure: true
verify: true
timeout: 120
compress: true
compression_codec: zstd
pool_size: 20
max_execution_time: 300
settings:
  max_memory_usage: 10000000000
  max_threads: 8
```

## 连接字符串格式

### HTTP接口
```
http://username:password@host:port/database
https://username:password@host:port/database
```

### Native协议 (TCP)
```
clickhouse://username:password@host:native_port/database
clickhouses://username:password@host:native_port/database
```

### Python (clickhouse-connect)
```python
import clickhouse_connect

client = clickhouse_connect.get_client(
    host='localhost',
    port=8123,
    username='default',
    password='password',
    database='myapp_db',
    compress=True,
    settings={'max_execution_time': 60}
)

# 执行查询
result = client.query('SELECT * FROM table LIMIT 10')
print(result.result_rows)
```

### Python (clickhouse-driver)
```python
from clickhouse_driver import Client

client = Client(
    host='localhost',
    port=9000,
    user='default',
    password='password',
    database='myapp_db',
    compression=True,
    settings={'max_execution_time': 60}
)

# 执行查询
result = client.execute('SELECT * FROM table LIMIT 10')
```

### Node.js (@clickhouse/client)
```javascript
import { createClient } from '@clickhouse/client';

const client = createClient({
    host: 'http://localhost:8123',
    username: 'default',
    password: 'password',
    database: 'myapp_db',
    request_timeout: 60000,
    max_open_connections: 10
});

const resultSet = await client.query({
    query: 'SELECT * FROM table LIMIT 10',
    format: 'JSONEachRow'
});
const rows = await resultSet.json();
```

## 常用命令

```bash
# 连接 ClickHouse (使用 clickhouse-client)
clickhouse-client -h host --port 9000 -u username --password password -d database

# 执行查询
clickhouse-client -q "SELECT * FROM system.tables LIMIT 10"

# 导出数据
clickhouse-client -q "SELECT * FROM table FORMAT CSV" > output.csv

# 导入数据
clickhouse-client -q "INSERT INTO table FORMAT CSV" < input.csv
```

## 性能建议

1. **批量插入**: 使用批量插入而非逐条插入，推荐每批 10,000-100,000 行
2. **分区键**: 合理设计表分区键，避免扫描过多分区
3. **索引**: 使用主键和跳数索引优化查询性能
4. **物化视图**: 对常用聚合查询创建物化视图
5. **数据类型**: 使用合适的数据类型，如用 Int32 替代 Int64 如果数据范围允许

## 注意事项

1. ClickHouse 擅长分析型查询，不适合高频小事务
2. 更新删除操作是异步的，且有性能开销
3. 分布式表需要配合分布式引擎使用
