# TBox 全栈项目实施计划

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 构建 TBox 软件商城的完整全栈系统，包括桌面端客户端、客户端后端服务、管理后台前端和管理后台后端服务。

**Architecture:** 采用前后端分离架构。两个 Go 后端（backend 和 admin-backend）共享 MySQL 数据库，使用 Gin + GORM 框架，Redis 做缓存和会话管理。桌面端使用 Electron + Vue3，管理后台使用 Vue3 SPA。四个子系统独立部署，通过 RESTful API 通信。

**Tech Stack:** Go 1.21+, Gin, GORM, MySQL, Redis, Vue3, TypeScript, Vite, Electron, Element Plus

---

## 子系统划分

本项目分为4个独立子系统，可并行开发：

### 子系统 1: Backend（客户端后端服务）
- 目录: `backend/`
- 职责: 为桌面端客户端提供 API 服务
- 功能: 用户注册/登录、软件商城浏览/搜索/详情、软件下载、软件上传、用户空间、评价系统

### 子系统 2: Admin-Backend（管理后台后端服务）
- 目录: `admin-backend/`
- 职责: 为管理后台前端提供 API 服务
- 功能: 管理员认证、仪表盘、软件管理、审核中心、内容运营、用户管理、数据统计、系统配置

### 子系统 3: Desktop（桌面端客户端）
- 目录: `desktop/`
- 职责: TBox 桌面端软件
- 功能: 用户界面、软件浏览下载安装、个人空间、设置管理

### 子系统 4: Admin-Frontend（管理后台前端）
- 目录: `admin-frontend/`
- 职责: TBox 后台运营管理平台 Web 前端
- 功能: 仪表盘、软件管理、审核中心、内容运营、用户管理、数据统计、系统管理

---

## 子系统 1: Backend 后端服务

### 项目结构

```
backend/
├── cmd/
│   └── server/
│       └── main.go                 # 入口文件
├── config/
│   ├── config.go                   # 配置结构定义与加载
│   └── config.yaml                 # 配置文件
├── internal/
│   ├── middleware/
│   │   ├── auth.go                 # JWT 认证中间件
│   │   ├── cors.go                 # 跨域中间件
│   │   └── logger.go               # 日志中间件
│   ├── model/
│   │   ├── user.go                 # 用户模型
│   │   ├── software.go             # 软件模型
│   │   ├── software_screenshot.go  # 软件截图模型
│   │   ├── software_version.go     # 软件版本模型
│   │   ├── software_review.go      # 软件评价模型
│   │   ├── user_space.go           # 用户空间模型
│   │   ├── category.go             # 分类模型
│   │   └── tag.go                  # 标签模型
│   ├── handler/
│   │   ├── auth.go                 # 认证相关 handler
│   │   ├── user.go                 # 用户相关 handler
│   │   ├── software.go             # 软件相关 handler
│   │   ├── category.go             # 分类相关 handler
│   │   ├── download.go             # 下载相关 handler
│   │   ├── upload.go               # 上传相关 handler
│   │   ├── review.go               # 评价相关 handler
│   │   └── space.go                # 个人空间 handler
│   ├── service/
│   │   ├── auth.go                 # 认证服务
│   │   ├── user.go                 # 用户服务
│   │   ├── software.go             # 软件服务
│   │   ├── category.go             # 分类服务
│   │   ├── download.go             # 下载服务
│   │   ├── upload.go               # 上传服务
│   │   ├── review.go               # 评价服务
│   │   └── space.go                # 个人空间服务
│   ├── repository/
│   │   ├── user.go                 # 用户数据访问
│   │   ├── software.go             # 软件数据访问
│   │   ├── category.go             # 分类数据访问
│   │   ├── review.go               # 评价数据访问
│   │   └── space.go                # 空间数据访问
│   └── router/
│       └── router.go               # 路由定义
├── pkg/
│   ├── database/
│   │   ├── mysql.go                # MySQL 连接
│   │   └── redis.go                # Redis 连接
│   ├── response/
│   │   └── response.go             # 统一响应格式
│   ├── utils/
│   │   ├── jwt.go                  # JWT 工具
│   │   ├── hash.go                 # 密码哈希工具
│   │   └── validator.go            # 验证工具
│   └── email/
│       └── email.go                # 邮件发送
├── go.mod
└── go.sum
```

### Task 1: 项目初始化与基础框架

**Files:**
- Create: `backend/go.mod`, `backend/cmd/server/main.go`
- Create: `backend/config/config.go`, `backend/config/config.yaml`
- Create: `backend/pkg/database/mysql.go`, `backend/pkg/database/redis.go`
- Create: `backend/pkg/response/response.go`
- Create: `backend/internal/router/router.go`
- Create: `backend/internal/middleware/cors.go`, `backend/internal/middleware/logger.go`

- [ ] **Step 1: 初始化 Go 模块**
```bash
cd backend && go mod init tbox-backend
```

- [ ] **Step 2: 创建配置文件和配置加载**

config.yaml:
```yaml
server:
  port: 8080
  mode: debug

database:
  host: localhost
  port: 3306
  username: root
  password: root
  dbname: tbox
  charset: utf8mb4
  max_idle_conns: 10
  max_open_conns: 100

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0

jwt:
  secret: tbox-secret-key
  expire: 720h

log:
  level: debug
  format: json

email:
  host: smtp.example.com
  port: 587
  username: ""
  password: ""
  from: noreply@tbox.com
```

config.go: 使用 viper 加载 yaml 配置

- [ ] **Step 3: 创建数据库连接（MySQL + Redis）**
- [ ] **Step 4: 创建统一响应格式**
- [ ] **Step 5: 创建中间件（CORS、Logger）**
- [ ] **Step 6: 创建路由框架和 main.go 入口**
- [ ] **Step 7: 安装依赖并验证编译通过**
```bash
go mod tidy && go build ./cmd/server/
```
- [ ] **Step 8: Commit**

### Task 2: 数据模型定义

**Files:**
- Create: `backend/internal/model/*.go` (所有模型文件)

- [ ] **Step 1: 创建 User 模型**（含 status, last_login_at 等扩展字段）
- [ ] **Step 2: 创建 Software 模型**（含 is_recommended, is_on_shelf 等扩展字段）
- [ ] **Step 3: 创建 SoftwareScreenshot, SoftwareVersion, SoftwareReview 模型**
- [ ] **Step 4: 创建 UserSpace, Category, Tag 模型**
- [ ] **Step 5: 在 main.go 中添加 AutoMigrate**
- [ ] **Step 6: 验证编译通过**
- [ ] **Step 7: Commit**

### Task 3: 用户认证模块（注册/登录）

**Files:**
- Create: `backend/pkg/utils/jwt.go`, `backend/pkg/utils/hash.go`
- Create: `backend/pkg/email/email.go`
- Create: `backend/internal/middleware/auth.go`
- Create: `backend/internal/repository/user.go`
- Create: `backend/internal/service/auth.go`
- Create: `backend/internal/handler/auth.go`

API 端点：
- POST /api/v1/auth/register - 邮箱注册
- POST /api/v1/auth/login - 邮箱登录
- POST /api/v1/auth/send-code - 发送验证码
- POST /api/v1/auth/reset-password - 重置密码
- POST /api/v1/auth/logout - 退出登录

- [ ] **Step 1: 创建 JWT 和密码哈希工具**
- [ ] **Step 2: 创建邮件发送服务**
- [ ] **Step 3: 创建 JWT 认证中间件**
- [ ] **Step 4: 创建用户 Repository**
- [ ] **Step 5: 创建认证 Service**（注册、登录、验证码、重置密码，登录失败5次锁定）
- [ ] **Step 6: 创建认证 Handler**
- [ ] **Step 7: 注册路由**
- [ ] **Step 8: 验证编译通过**
- [ ] **Step 9: Commit**

### Task 4: 用户个人中心模块

**Files:**
- Create: `backend/internal/service/user.go`
- Create: `backend/internal/handler/user.go`

API 端点：
- GET /api/v1/user/profile - 获取个人信息
- PUT /api/v1/user/profile - 更新个人信息
- PUT /api/v1/user/avatar - 更新头像
- PUT /api/v1/user/password - 修改密码
- DELETE /api/v1/user/account - 注销账户

- [ ] **Step 1-5: 创建用户 Service 和 Handler，注册路由**
- [ ] **Step 6: Commit**

### Task 5: 软件商城模块（浏览/搜索/详情）

**Files:**
- Create: `backend/internal/repository/software.go`, `backend/internal/repository/category.go`
- Create: `backend/internal/service/software.go`, `backend/internal/service/category.go`
- Create: `backend/internal/handler/software.go`, `backend/internal/handler/category.go`

API 端点：
- GET /api/v1/software/home - 首页数据（轮播+推荐+排行）
- GET /api/v1/software/list - 软件列表（分类筛选+排序+分页）
- GET /api/v1/software/:id - 软件详情
- GET /api/v1/software/search - 搜索软件
- GET /api/v1/categories - 获取分类列表
- GET /api/v1/software/:id/versions - 版本历史
- GET /api/v1/software/:id/reviews - 评价列表

- [ ] **Step 1-6: 创建 Repository, Service, Handler，注册路由**
- [ ] **Step 7: Commit**

### Task 6: 软件下载模块

**Files:**
- Create: `backend/internal/service/download.go`
- Create: `backend/internal/handler/download.go`

API 端点：
- GET /api/v1/software/:id/download - 获取下载链接（支持断点续传 Range header）
- POST /api/v1/software/:id/download-count - 记录下载次数
- GET /api/v1/software/:id/check-update - 检查更新

- [ ] **Step 1-4: 创建下载 Service 和 Handler**
- [ ] **Step 5: Commit**

### Task 7: 软件上传与个人空间模块

**Files:**
- Create: `backend/internal/repository/space.go`
- Create: `backend/internal/service/upload.go`, `backend/internal/service/space.go`
- Create: `backend/internal/handler/upload.go`, `backend/internal/handler/space.go`

API 端点：
- GET /api/v1/space/overview - 空间概览
- GET /api/v1/space/software - 我的软件列表
- POST /api/v1/space/software - 上传软件信息
- PUT /api/v1/space/software/:id - 更新软件信息
- POST /api/v1/upload/init - 初始化分片上传
- POST /api/v1/upload/chunk - 上传分片
- POST /api/v1/upload/complete - 完成上传合并
- POST /api/v1/space/software/:id/submit-audit - 提交审核
- POST /api/v1/space/software/:id/withdraw - 撤回审核

- [ ] **Step 1-6: 创建 Repository, Service, Handler，注册路由**
- [ ] **Step 7: Commit**

### Task 8: 评价模块

**Files:**
- Create: `backend/internal/repository/review.go`
- Create: `backend/internal/service/review.go`
- Create: `backend/internal/handler/review.go`

API 端点：
- POST /api/v1/software/:id/reviews - 发表评价
- PUT /api/v1/reviews/:id - 修改评价
- DELETE /api/v1/reviews/:id - 删除评价

- [ ] **Step 1-4: 创建 Repository, Service, Handler**
- [ ] **Step 5: Commit**

---

## 子系统 2: Admin-Backend 管理后台后端

### 项目结构

```
admin-backend/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   ├── config.go
│   └── config.yaml
├── internal/
│   ├── middleware/
│   │   ├── auth.go                 # 管理员 JWT 认证
│   │   ├── cors.go
│   │   ├── logger.go
│   │   └── permission.go           # 权限校验中间件
│   ├── model/
│   │   ├── admin.go                # 管理员模型
│   │   ├── role.go                 # 角色模型
│   │   ├── permission.go           # 权限模型
│   │   ├── operation_log.go        # 操作日志模型
│   │   ├── banner.go               # Banner 模型
│   │   ├── announcement.go         # 公告模型
│   │   ├── report.go               # 举报模型
│   │   ├── feedback.go             # 用户反馈模型
│   │   ├── recommendation.go       # 推荐配置模型
│   │   ├── system_config.go        # 系统配置模型
│   │   └── shared.go               # 共享模型（引用 backend 的表）
│   ├── handler/
│   │   ├── auth.go                 # 管理员认证
│   │   ├── dashboard.go            # 仪表盘
│   │   ├── software.go             # 软件管理
│   │   ├── category.go             # 分类管理
│   │   ├── tag.go                  # 标签管理
│   │   ├── banner.go               # Banner 管理
│   │   ├── recommendation.go       # 推荐管理
│   │   ├── announcement.go         # 公告管理
│   │   ├── audit.go                # 软件审核
│   │   ├── review_audit.go         # 评论审核
│   │   ├── report.go               # 举报处理
│   │   ├── user_mgmt.go            # 用户管理
│   │   ├── feedback.go             # 用户反馈
│   │   ├── statistics.go           # 数据统计
│   │   ├── admin_mgmt.go           # 管理员管理
│   │   ├── role.go                 # 角色权限
│   │   ├── operation_log.go        # 操作日志
│   │   └── system_config.go        # 系统配置
│   ├── service/
│   │   ├── (与 handler 一一对应)
│   ├── repository/
│   │   ├── (与数据模型一一对应)
│   └── router/
│       └── router.go
├── pkg/
│   ├── database/
│   │   ├── mysql.go
│   │   └── redis.go
│   ├── response/
│   │   └── response.go
│   └── utils/
│       ├── jwt.go
│       └── hash.go
├── go.mod
└── go.sum
```

### Task 1-3: 基础框架、数据模型、管理员认证
（结构同 backend，但端口 8081，JWT 密钥独立，使用管理员表）

### Task 4: 仪表盘模块
- GET /api/v1/dashboard/stats - 核心指标
- GET /api/v1/dashboard/download-trend - 下载趋势
- GET /api/v1/dashboard/user-trend - 用户增长趋势
- GET /api/v1/dashboard/category-distribution - 分类分布
- GET /api/v1/dashboard/top-software - 热门 Top10
- GET /api/v1/dashboard/todos - 待办事项

### Task 5: 软件管理模块
- GET /api/v1/software - 软件列表（筛选+分页+排序）
- GET /api/v1/software/:id - 软件详情
- PUT /api/v1/software/:id - 编辑软件
- PUT /api/v1/software/:id/on-shelf - 上架
- PUT /api/v1/software/:id/off-shelf - 下架
- DELETE /api/v1/software/:id - 删除
- CRUD /api/v1/categories - 分类管理
- CRUD /api/v1/tags - 标签管理

### Task 6: 内容运营模块
- CRUD /api/v1/banners - Banner 管理
- CRUD /api/v1/recommendations - 推荐管理
- CRUD /api/v1/announcements - 公告管理

### Task 7: 审核中心模块
- GET /api/v1/audits - 审核列表
- GET /api/v1/audits/:id - 审核详情
- POST /api/v1/audits/:id/approve - 通过
- POST /api/v1/audits/:id/reject - 拒绝
- CRUD /api/v1/review-audits - 评论审核
- CRUD /api/v1/reports - 举报处理

### Task 8: 用户管理模块
- GET /api/v1/users - 用户列表
- GET /api/v1/users/:id - 用户详情
- PUT /api/v1/users/:id/disable - 禁用
- PUT /api/v1/users/:id/enable - 启用
- PUT /api/v1/users/:id/space - 调整空间
- CRUD /api/v1/feedbacks - 用户反馈

### Task 9: 数据统计模块
- GET /api/v1/stats/downloads - 下载统计
- GET /api/v1/stats/users - 用户统计
- GET /api/v1/stats/software-ranking - 软件排行
- GET /api/v1/stats/export - 导出 Excel

### Task 10: 系统管理模块
- CRUD /api/v1/admins - 管理员管理
- CRUD /api/v1/roles - 角色管理
- GET /api/v1/permissions - 权限列表
- PUT /api/v1/roles/:id/permissions - 配置权限
- GET /api/v1/operation-logs - 操作日志
- GET /api/v1/system-configs - 系统配置
- PUT /api/v1/system-configs - 更新配置

---

## 子系统 3: Desktop 桌面端

### 项目结构

```
desktop/
├── electron/
│   ├── main.ts                     # Electron 主进程
│   ├── preload.ts                  # 预加载脚本
│   └── tray.ts                     # 系统托盘
├── src/
│   ├── main.ts                     # Vue 入口
│   ├── App.vue                     # 根组件
│   ├── api/
│   │   ├── request.ts              # axios 封装
│   │   ├── auth.ts                 # 认证 API
│   │   ├── user.ts                 # 用户 API
│   │   ├── software.ts             # 软件 API
│   │   ├── download.ts             # 下载 API
│   │   ├── upload.ts               # 上传 API
│   │   └── space.ts                # 空间 API
│   ├── router/
│   │   └── index.ts                # 路由配置
│   ├── stores/
│   │   ├── auth.ts                 # 认证状态
│   │   ├── download.ts             # 下载管理状态
│   │   ├── upload.ts               # 上传管理状态
│   │   └── settings.ts             # 设置状态
│   ├── views/
│   │   ├── auth/
│   │   │   └── LoginRegister.vue   # 登录注册页
│   │   ├── home/
│   │   │   └── HomePage.vue        # 首页
│   │   ├── category/
│   │   │   └── CategoryBrowse.vue  # 分类浏览页
│   │   ├── software/
│   │   │   └── SoftwareDetail.vue  # 软件详情页
│   │   ├── download/
│   │   │   └── DownloadManager.vue # 下载管理页
│   │   ├── install/
│   │   │   └── InstallRecord.vue   # 安装记录页
│   │   ├── space/
│   │   │   └── PersonalSpace.vue   # 个人空间页
│   │   ├── profile/
│   │   │   └── UserProfile.vue     # 个人中心页
│   │   └── settings/
│   │       └── SettingsPage.vue    # 设置页
│   ├── components/
│   │   ├── layout/
│   │   │   ├── AppLayout.vue       # 主布局（侧边栏+顶栏+内容区）
│   │   │   ├── Sidebar.vue         # 侧边栏
│   │   │   └── TopBar.vue          # 顶部导航栏
│   │   ├── software/
│   │   │   ├── SoftwareCard.vue    # 软件卡片
│   │   │   └── BannerCarousel.vue  # 轮播组件
│   │   ├── upload/
│   │   │   └── UploadDialog.vue    # 上传弹窗
│   │   └── common/
│   │       └── SearchBox.vue       # 搜索框
│   └── styles/
│       └── global.css              # 全局样式
├── index.html
├── package.json
├── tsconfig.json
├── vite.config.ts
└── electron-builder.json5
```

### Task 1: 项目初始化
- 使用 Vite 创建 Vue3+TS 项目
- 安装依赖: vue-router, pinia, axios, element-plus, electron, electron-builder
- 配置 Electron 主进程

### Task 2: 主布局框架
- 创建 AppLayout, Sidebar, TopBar
- 配置路由

### Task 3: 登录注册页
### Task 4: 首页（轮播+分类标签+热门推荐+排行榜）
### Task 5: 分类浏览页
### Task 6: 软件详情页
### Task 7: 下载管理页
### Task 8: 安装记录页
### Task 9: 个人空间页（含上传弹窗）
### Task 10: 个人中心页
### Task 11: 设置页
### Task 12: Electron 集成（窗口控制、系统托盘、本地数据库）

---

## 子系统 4: Admin-Frontend 管理后台前端

### 项目结构

```
admin-frontend/
├── src/
│   ├── main.ts
│   ├── App.vue
│   ├── api/
│   │   ├── request.ts              # axios 封装
│   │   ├── auth.ts                 # 管理员认证 API
│   │   ├── dashboard.ts            # 仪表盘 API
│   │   ├── software.ts             # 软件管理 API
│   │   ├── category.ts             # 分类管理 API
│   │   ├── tag.ts                  # 标签管理 API
│   │   ├── banner.ts               # Banner API
│   │   ├── recommendation.ts       # 推荐管理 API
│   │   ├── announcement.ts         # 公告管理 API
│   │   ├── audit.ts                # 审核 API
│   │   ├── review-audit.ts         # 评论审核 API
│   │   ├── report.ts               # 举报 API
│   │   ├── user.ts                 # 用户管理 API
│   │   ├── feedback.ts             # 反馈 API
│   │   ├── statistics.ts           # 统计 API
│   │   ├── admin.ts                # 管理员管理 API
│   │   ├── role.ts                 # 角色权限 API
│   │   ├── operation-log.ts        # 操作日志 API
│   │   └── system-config.ts        # 系统配置 API
│   ├── router/
│   │   └── index.ts
│   ├── stores/
│   │   ├── auth.ts
│   │   └── permission.ts
│   ├── views/
│   │   ├── login/
│   │   │   └── LoginPage.vue
│   │   ├── dashboard/
│   │   │   └── DashboardPage.vue
│   │   ├── software/
│   │   │   ├── SoftwareList.vue
│   │   │   ├── SoftwareDetail.vue
│   │   │   ├── CategoryManage.vue
│   │   │   └── TagManage.vue
│   │   ├── content/
│   │   │   ├── BannerManage.vue
│   │   │   ├── RecommendManage.vue
│   │   │   └── AnnouncementManage.vue
│   │   ├── audit/
│   │   │   ├── SoftwareAudit.vue
│   │   │   ├── ReviewAudit.vue
│   │   │   └── ReportHandle.vue
│   │   ├── user/
│   │   │   ├── UserList.vue
│   │   │   ├── UserDetail.vue
│   │   │   └── FeedbackManage.vue
│   │   ├── statistics/
│   │   │   ├── DownloadStats.vue
│   │   │   ├── UserStats.vue
│   │   │   └── SoftwareRanking.vue
│   │   └── system/
│   │       ├── AdminManage.vue
│   │       ├── RolePermission.vue
│   │       ├── OperationLog.vue
│   │       └── SystemConfig.vue
│   ├── components/
│   │   ├── layout/
│   │   │   ├── AdminLayout.vue
│   │   │   ├── SideMenu.vue
│   │   │   └── TopNav.vue
│   │   └── common/
│   │       └── SearchFilter.vue
│   └── styles/
│       └── global.css
├── index.html
├── package.json
├── tsconfig.json
└── vite.config.ts
```

### Task 1: 项目初始化
### Task 2: 主布局框架（侧边栏+顶栏+面包屑）
### Task 3: 管理员登录页
### Task 4: 仪表盘页（指标卡片+图表）
### Task 5: 软件管理页（列表+详情+分类+标签）
### Task 6: 内容运营页（Banner+推荐+公告）
### Task 7: 审核中心页（软件审核+评论审核+举报）
### Task 8: 用户管理页（用户列表+详情+反馈）
### Task 9: 数据统计页（下载统计+用户统计+排行）
### Task 10: 系统管理页（管理员+角色权限+日志+配置）

---

## 执行策略

四个子系统完全独立，使用并行 Agent 团队同时开发。每个 Agent 负责一个子系统的完整开发。
