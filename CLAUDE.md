# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

TBox 是一个软件分发与管理平台，包含四个子项目共享同一仓库：

| 目录 | 说明 | 技术栈 | 端口 |
|------|------|--------|------|
| `backend/` | 用户端后端 API | Go 1.21 + Gin + GORM + Redis | 8080 |
| `admin-backend/` | 运营管理后台 API | Go 1.21 + Gin + GORM + Redis | 8081 |
| `desktop/` | 桌面客户端 | Vue 3 + TypeScript + Vite + Electron | 5173 (dev) |
| `admin-frontend/` | 管理后台前端 | Vue 3 + TypeScript + Vite + ECharts | 3000 (dev) |

两个后端共享同一 MySQL 数据库 `tbox`，Redis 分别使用 DB 0（backend）和 DB 1（admin-backend）。

## Common Commands

### Backend / Admin-Backend (Go)

```bash
# 运行
cd backend && go run ./cmd/server/main.go
cd admin-backend && go run ./cmd/server/main.go

# 构建
cd backend && go build -o server ./cmd/server/
cd admin-backend && go build -o server ./cmd/server/
```

### Desktop (Electron + Vue)

```bash
cd desktop
npm run dev              # Vue 开发服务器
npm run electron:dev     # Electron 开发模式
npm run electron:build   # 打包桌面应用
npm run build            # 生产构建（含类型检查）
```

### Admin Frontend (Vue)

```bash
cd admin-frontend
npm run dev    # 开发服务器（端口 3000，自动代理到 admin-backend:8081）
npm run build  # 生产构建（含类型检查）
```

## Architecture

### 后端分层结构（backend/ 和 admin-backend/ 结构相同）

```
cmd/server/main.go          # 入口
config/config.yaml           # YAML 配置（端口、DB、Redis、JWT、上传等）
internal/
  router/router.go           # API 路由定义（/api/v1/...）
  middleware/                 # 中间件（auth JWT 验证、CORS、日志）
  handler/                   # HTTP 请求处理 → 调用 service
  service/                   # 业务逻辑 → 调用 repository
  repository/                # 数据访问层（GORM）
  model/                     # 数据库模型定义
pkg/                         # 共享工具（数据库连接、邮件、响应格式化）
```

请求流：`Router → Middleware → Handler → Service → Repository → MySQL/Redis`

### 前端结构（desktop/ 和 admin-frontend/ 结构相同）

```
src/
  api/          # Axios API 客户端模块
  router/       # Vue Router 路由配置
  stores/       # Pinia 状态管理
  views/        # 页面组件
  components/   # 可复用组件
  styles/       # 全局样式
```

Desktop 额外包含 `electron/` 目录（main.ts 主进程 + preload.ts）。

### 关键业务模块

- **用户认证**：邮箱+密码登录，JWT Token，Redis 存储会话
- **软件市场**：软件 CRUD、分类/标签、版本管理、下载追踪
- **评价系统**：用户评分与评论
- **个人空间**：用户上传软件（分片上传，最大 2GB）
- **运营管理**：数据看板、内容审核、用户管理、审计日志、权限控制

## Development Notes

- 配置文件位于各子项目的 `config/config.yaml`，包含数据库、Redis、JWT 等连接信息
- 需求文档在 `docs/` 目录：`需求文档-细化版.md`（用户端）、`后台运营管理平台-需求文档.md`（管理端）
- 原型文件在 `docs/prototype/`：`index.html`（用户端）、`admin.html`（管理端）
- admin-frontend 的 Vite 配置已设置 `/api` 代理到 `http://localhost:8081`
