# MongoDB 连接配置

## 连接配置字段

| 字段名 | 说明 | 示例 |
|--------|------|------|
| host | 主机地址 | localhost / mongodb.example.com |
| port | 端口号 | 27017 |
| username | 用户名 | app_user |
| password | 密码 | - |
| database | 数据库名 | myapp_db |
| authSource | 认证数据库 | admin |
| replicaSet | 副本集名称 | rs0 |
| ssl | 启用SSL | true / false |
| connection_timeout | 连接超时(ms) | 5000 |
| server_selection_timeout | 服务器选择超时(ms) | 5000 |
| pool_size | 连接池大小 | 10 |
| retry_writes | 重试写入 | true |

## 环境配置模板

### 开发环境 (dev)

```yaml
env: dev
host: localhost
port: 27017
username: dev_user
password: ${MONGODB_DEV_PASSWORD}
database: myapp_dev
authSource: admin
ssl: false
connection_timeout: 5000
server_selection_timeout: 5000
pool_size: 5
retry_writes: true
```

### 测试环境 (test)

```yaml
env: test
host: test-mongodb.example.com
port: 27017
username: test_user
password: ${MONGODB_TEST_PASSWORD}
database: myapp_test
authSource: admin
replicaSet: rs_test
ssl: true
connection_timeout: 5000
server_selection_timeout: 5000
pool_size: 10
retry_writes: true
```

### 生产环境 (prod)

```yaml
env: prod
hosts:
  - mongo1.example.com:27017
  - mongo2.example.com:27017
  - mongo3.example.com:27017
username: prod_user
password: ${MONGODB_PROD_PASSWORD}
database: myapp_prod
authSource: admin
replicaSet: rs_prod
ssl: true
ssl_ca_file: /path/to/ca.pem
connection_timeout: 10000
server_selection_timeout: 30000
pool_size: 50
retry_writes: true
max_pool_size: 100
min_pool_size: 10
```

## 连接字符串格式

### 单机模式
```
mongodb://username:password@host:port/database?authSource=admin
```

### 副本集模式
```
mongodb://username:password@host1:port1,host2:port2,host3:port3/database?replicaSet=rs0&authSource=admin&ssl=true
```

### Python (pymongo)
```python
from pymongo import MongoClient

client = MongoClient(
    host="localhost",
    port=27017,
    username="dev_user",
    password="password",
    authSource="admin",
    serverSelectionTimeoutMS=5000,
    maxPoolSize=10
)
db = client["myapp_db"]
```

### Node.js (mongoose)
```javascript
const mongoose = require('mongoose');

await mongoose.connect('mongodb://username:password@host:port/database', {
    authSource: 'admin',
    useNewUrlParser: true,
    useUnifiedTopology: true,
    serverSelectionTimeoutMS: 5000,
    maxPoolSize: 10
});
```

## 常用命令

```bash
# 连接 MongoDB
mongo mongodb://username:password@host:port/database

# 或使用 mongosh
mongosh "mongodb://username:password@host:port/database"

# 导出集合
mongodump --uri="mongodb://username:password@host:port/database" --collection=mycollection

# 导入集合
mongorestore --uri="mongodb://username:password@host:port/database" dump/
```
