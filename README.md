<div align="center">

# 🎁 Gift Wishlist

**一个轻量、美观的礼物愿望清单 Web 应用**

朋友们可以浏览你的愿望清单并直接跳转购买链接，你在后台管理物品、配置收货地址——单文件部署，开箱即用。

<br>

[![CI/CD](https://img.shields.io/github/actions/workflow/status/transnull/giftlist/docker-publish.yml?style=flat-square&logo=github-actions&label=CI%2FCD)](https://github.com/transnull/giftlist/actions)
[![Docker Pulls](https://img.shields.io/docker/pulls/transnull/giftlist?style=flat-square&logo=docker&color=2496ED)](https://hub.docker.com/r/transnull/giftlist)
[![Image Size](https://img.shields.io/docker/image-size/transnull/giftlist/latest?style=flat-square&logo=docker&label=image%20size&color=2496ED)](https://hub.docker.com/r/transnull/giftlist)
[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![Vue](https://img.shields.io/badge/Vue-3-4FC08D?style=flat-square&logo=vuedotjs)](https://vuejs.org)
[![SQLite](https://img.shields.io/badge/SQLite-embedded-003B57?style=flat-square&logo=sqlite)](https://sqlite.org)
[![License](https://img.shields.io/github/license/transnull/giftlist?style=flat-square)](LICENSE)

<br>

[🐳 Docker 部署](#-docker-部署推荐) · [🏗️ 本地构建](#️-本地构建) · [🖼️ 截图预览](#️-截图预览) · [📡 API 文档](#-api-文档)

**示例网站**：[礼物愿望清单](https://giftlist.transnull.cn)


</div>

---

## ✨ 功能特性

| | 特性 | 说明 |
|---|------|------|
| 🛍️ | **礼物卡片展示** | 响应式网格布局，图片统一裁剪，完美适配 PC / 平板 / 手机 |
| 🔒 | **后台管理** | 密码保护，支持物品增删改查、站点名称与引言（含收货地址）配置 |
| 🏷️ | **优先级 & 类别** | 高 / 中 / 低优先级彩色标签，支持分类筛选与多维度排序 |
| 💰 | **价格参考** | 显示估计费用，一键跳转购买链接，新标签页打开 |
| 🗃️ | **SQLite 嵌入** | 零依赖数据库，数据以单文件持久化，迁移极简 |
| 📦 | **单文件部署** | 前端编译产物通过 `go:embed` 嵌入二进制，无需额外文件 |
| 🐳 | **Docker 支持** | 多阶段构建，最终镜像 < 25 MB，支持 `amd64` / `arm64` |

---

## 🖼️ 截图预览

### PC 端

<table>
  <tr>
    <td align="center"><b>公开首页 — 礼物卡片网格</b></td>
  </tr>
  <tr>
    <td><a href="https://imgchr.com/i/peV75N9"><img src="https://s41.ax1x.com/2026/03/16/peV75N9.png" alt="首页截图"></a></td>
  </tr>
  <tr>
    <td align="center"><b>礼物详情页</b></td>
  </tr>
  <tr>
    <td><a href="https://imgchr.com/i/peV7bjK"><img src="https://s41.ax1x.com/2026/03/16/peV7bjK.png" alt="详情页截图"></a></td>
  </tr>
  <tr>
    <td align="center"><b>后台管理页</b></td>
  </tr>
  <tr>
    <td><a href="https://imgchr.com/i/peV7xNd"><img src="https://s41.ax1x.com/2026/03/16/peV7xNd.png" alt="后台管理截图"></a></td>
  </tr>
</table>

### 移动端

<table>
  <tr>
    <td align="center"><b>首页</b></td>
    <td align="center"><b>详情页</b></td>
    <td align="center"><b>后台管理</b></td>
  </tr>
  <tr>
    <td><a href="https://imgchr.com/i/peVHZNj"><img src="https://s41.ax1x.com/2026/03/16/peVHZNj.png" alt="移动端首页"></a></td>
    <td><a href="https://imgchr.com/i/peVHVEQ"><img src="https://s41.ax1x.com/2026/03/16/peVHVEQ.png" alt="移动端详情"></a></td>
    <td><a href="https://imgchr.com/i/peVHAHg"><img src="https://s41.ax1x.com/2026/03/16/peVHAHg.png" alt="移动端后台"></a></td>
  </tr>
</table>

---

## 🐳 Docker 部署（推荐）

### 方式一：docker compose（最简单）

```bash
# 1. 下载 docker-compose.yml
curl -O https://raw.githubusercontent.com/transnull/giftlist/main/docker-compose.yml

# 2. 创建 .env 文件，配置你的密码
cat > .env <<EOF
ADMIN_PASSWORD=your_secure_password
SESSION_SECRET=a_long_random_secret_string
EOF

# 3. 启动服务
docker compose up -d

# 查看日志
docker compose logs -f
```

> 访问 → **http://localhost:8080**

### 方式二：docker run

```bash
docker run -d \
  --name giftlist \
  --restart unless-stopped \
  -p 8080:8080 \
  -v giftlist_data:/data \
  -e ADMIN_PASSWORD=your_secure_password \
  -e SESSION_SECRET=a_long_random_secret_string \
  transnull/giftlist:latest
```

### 更新镜像

```bash
# docker compose 方式
docker compose pull && docker compose up -d

# docker run 方式
docker pull transnull/giftlist:latest && docker restart giftlist
```

---

## ⚙️ 环境变量

| 变量 | 必填 | 默认值 | 说明 |
|------|:----:|--------|------|
| `ADMIN_PASSWORD` | ✅ | `admin123` | 后台登录密码，**生产环境必须修改** |
| `SESSION_SECRET` | 推荐 | 随机生成 | Session 加密密钥，未设置则重启后需重新登录 |
| `PORT` | ❌ | `8080` | 服务监听端口 |
| `DB_PATH` | ❌ | `/data/giftlist.db` | SQLite 数据库文件路径（Docker 中请挂载 `/data` 目录） |

---

## 🌐 页面路由

| 页面 | 路径 | 说明 |
|------|------|------|
| 公开首页 | `/` | 朋友们浏览礼物卡片 |
| 礼物详情 | `/item/:id` | 查看详情，跳转购买链接 |
| 后台登录 | `/admin/login` | 输入管理密码 |
| 后台管理 | `/admin` | 管理物品、配置站点信息 |

---

## 🏗️ 本地构建

### 前置要求

- [Go 1.21+](https://golang.org/dl/)（需要 CGO / GCC 支持）
- [Node.js 18+](https://nodejs.org/)

### Linux / macOS

```bash
git clone https://github.com/transnull/giftlist.git
cd giftlist

# 构建前端
cd frontend && npm ci && npm run build && cd ..

# 构建 Go 二进制
CGO_ENABLED=1 go build -o giftlist .

# 运行
ADMIN_PASSWORD=your_password ./giftlist
```

### Windows

```bat
git clone https://github.com/transnull/giftlist.git
cd giftlist

cd frontend && npm ci && npm run build && cd ..

set CGO_ENABLED=1
go build -o giftlist.exe .

set ADMIN_PASSWORD=your_password
giftlist.exe
```

### 本地 Docker 构建

```bash
docker build -t giftlist .
docker run -p 8080:8080 -e ADMIN_PASSWORD=your_password giftlist
```

---

## 🔧 Nginx 反向代理

将 Gift Wishlist 挂载到域名并启用 HTTPS：

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass         http://127.0.0.1:8080;
        proxy_set_header   Host              $host;
        proxy_set_header   X-Real-IP         $remote_addr;
        proxy_set_header   X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header   X-Forwarded-Proto $scheme;
    }
}
```

配合 Certbot 一键启用 HTTPS：

```bash
certbot --nginx -d your-domain.com
```

---

## 📁 项目结构

```
giftlist/
├── main.go                 # 入口：路由注册 + go:embed 静态文件
├── Dockerfile              # 多阶段构建（node → go → 精简运行镜像）
├── docker-compose.yml      # 一键部署配置
├── .env.example            # 环境变量示例
│
├── handlers/
│   ├── auth.go             # 登录 / 登出 / 认证检查
│   ├── config.go           # 站点配置 API
│   └── items.go            # 物品 CRUD API
│
├── models/
│   ├── db.go               # SQLite 初始化 & 建表
│   ├── config.go           # 配置模型
│   └── item.go             # 物品模型（分页 / 筛选 / 排序）
│
├── middleware/
│   └── auth.go             # Session 认证中间件
│
├── static/                 # 前端编译产物（嵌入二进制，无需单独部署）
│
└── frontend/               # Vue 3 + Tailwind CSS 源码
    └── src/
        ├── views/          # 页面组件（首页、详情、后台登录、后台管理）
        ├── components/     # 通用组件（GiftCard 等）
        ├── router/         # 路由配置（含导航守卫）
        └── stores/         # Pinia 状态管理（认证）
```

---

## 📡 API 文档

所有接口统一前缀 `/api`，需认证的接口通过 Cookie Session 鉴权。

### 认证接口

| 方法 | 路径 | 认证 | 说明 |
|------|------|:----:|------|
| `POST` | `/api/login` | ❌ | 登录，请求体 `{"password":"..."}` |
| `POST` | `/api/logout` | ❌ | 登出，清除 Session |
| `GET` | `/api/auth/check` | ❌ | 检查当前登录状态 |

### 配置接口

| 方法 | 路径 | 认证 | 说明 |
|------|------|:----:|------|
| `GET` | `/api/config` | ❌ | 获取站点配置（站点名 + 引言） |
| `PUT` | `/api/config` | ✅ | 更新配置 `{"site_name":"...","intro_text":"..."}` |

### 物品接口

| 方法 | 路径 | 认证 | 说明 |
|------|------|:----:|------|
| `GET` | `/api/items` | ❌ | 物品列表（支持分页、筛选、排序） |
| `GET` | `/api/items/:id` | ❌ | 获取单个物品详情 |
| `POST` | `/api/items` | ✅ | 新增物品 |
| `PUT` | `/api/items/:id` | ✅ | 更新物品 |
| `DELETE` | `/api/items/:id` | ✅ | 删除物品 |

**`GET /api/items` 查询参数：**

| 参数 | 类型 | 默认 | 说明 |
|------|------|:----:|------|
| `page` | number | `1` | 页码 |
| `limit` | number | `10` | 每页数量 |
| `category` | string | — | 类别模糊匹配 |
| `priority` | string | — | 优先级筛选：`high` / `medium` / `low` |
| `sort` | string | `created_at` | 排序字段：`created_at` / `estimated_cost` |
| `order` | string | `desc` | 排序方向：`asc` / `desc` |

---

## 🤝 参与贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本仓库
2. 创建功能分支：`git checkout -b feature/amazing-feature`
3. 提交改动：`git commit -m 'feat: add amazing feature'`
4. 推送分支：`git push origin feature/amazing-feature`
5. 发起 Pull Request

---

## 📄 License

[MIT License](LICENSE) © [transnull](https://github.com/transnull)

---

<div align="center">

Made with ❤️ by [transnull](https://github.com/transnull)

**如果这个项目对你有帮助，欢迎点个 ⭐ Star！**

</div>
