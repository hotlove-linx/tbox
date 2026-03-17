# MySQL 连接配置

## 连接配置字段

| 字段名 | 说明 | 示例 |
|--------|------|------|
| host | 数据库主机地址 | localhost / 127.0.0.1 / mysql.example.com |
| port | 端口号 | 3306 |
| username | 用户名 | root / app_user |
| password | 密码 | - |
| database | 数据库名 | myapp_db |
| charset | 字符集 | utf8mb4 |
| ssl_mode | SSL模式 | disable / require / verify-ca / verify-full |
| connection_timeout | 连接超时(秒) | 30 |
| pool_size | 连接池大小 | 10 |

## 环境配置模板

### 开发环境 (dev)

```yaml
env: dev
host: 172.31.187.209
port: 3306
username: root
password: aimind@mysql2019
database: aimind_uc_graph
charset: utf8mb4
ssl_mode: disable
connection_timeout: 30
pool_size: 5
```

### 测试环境 (test)

```yaml
env: test
host: test-mysql.example.com
port: 3306
username: test_user
password: ${MYSQL_TEST_PASSWORD}
database: myapp_test
charset: utf8mb4
ssl_mode: require
connection_timeout: 30
pool_size: 10
```

### 生产环境 (prod)

```yaml
env: prod
host: prod-mysql.example.com
port: 3306
username: prod_user
password: ${MYSQL_PROD_PASSWORD}
database: myapp_prod
charset: utf8mb4
ssl_mode: verify-ca
connection_timeout: 60
pool_size: 20
```

## 连接字符串格式

### JDBC
```
jdbc:mysql://host:port/database?characterEncoding=utf8&useSSL=true
```

### Python (mysql-connector-python)
```python
import mysql.connector
conn = mysql.connector.connect(
    host="localhost",
    port=3306,
    user="username",
    password="password",
    database="myapp_db",
    charset="utf8mb4"
)
```

### Node.js (mysql2)
```javascript
const mysql = require('mysql2/promise');
const pool = mysql.createPool({
    host: 'localhost',
    port: 3306,
    user: 'username',
    password: 'password',
    database: 'myapp_db',
    waitForConnections: true,
    connectionLimit: 10
});
```

## 常用命令

```bash
# 连接 MySQL
mysql -h host -P port -u username -p

# 导出数据库
mysqldump -h host -u username -p database > backup.sql

# 导入数据库
mysql -h host -u username -p database < backup.sql
```
