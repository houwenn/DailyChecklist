# DailyChecklist - 日常待办事项管理系统

这是一个基于Go语言和Gin框架开发的待办事项管理系统，支持创建、查看、更新、删除和存档待办事项等功能。

## 项目特性

- 📝 **创建待办事项**：快速添加新的待办事项
- 📋 **查看待办列表**：浏览所有活跃的待办事项
- ✅ **更新待办状态**：标记待办事项为已完成/未完成
- 🗑️ **删除待办事项**：移除不需要的待办事项
- 🗄️ **存档功能**：将已完成的待办事项存档，保持列表整洁
- 📊 **统计功能**：查看待办事项的统计数据和历史趋势

## 技术栈

- **后端**：Go 1.13
- **Web框架**：Gin 1.7.7
- **ORM框架**：GORM 1.9.16
- **数据库**：MySQL
- **配置管理**：go-ini
- **前端**：HTML、CSS、JavaScript

## 项目结构

```
├── .gitignore           # Git忽略文件
├── conf/                # 配置文件目录
│   └── config.ini       # 应用程序配置文件
├── controller/          # 控制器目录
│   └── controller.go    # 请求处理逻辑
├── dao/                 # 数据访问层
│   └── mysql.go         # 数据库连接管理
├── go.mod               # Go模块依赖
├── go.sum               # 依赖版本锁定
├── main.go              # 程序入口文件
├── models/              # 数据模型目录
│   └── todo.go          # Todo数据模型及操作
├── routers/             # 路由配置目录
│   └── routers.go       # 路由定义与注册
├── setting/             # 配置管理目录
│   └── setting.go       # 配置结构与加载
├── static/              # 静态资源目录
│   ├── css/             # 样式文件
│   ├── fonts/           # 字体文件
│   └── js/              # JavaScript文件
└── templates/           # 模板文件目录
    ├── archive.html     # 存档页面模板
    └── index.html       # 主页面模板
```

## 快速开始

### 环境要求

- Go 1.13或更高版本
- MySQL 5.7或更高版本

### 安装步骤

1. **克隆项目**

```bash
git clone https://github.com/your-username/DailyChecklist.git
cd DailyChecklist
```

2. **配置数据库**

首先创建数据库：

```sql
CREATE DATABASE DailyChecklist;
```

然后修改`conf/config.ini`文件中的数据库配置：

```ini
[mysql]
user = root
password = 123456
host = 127.0.0.1
port = 3306
db = checklist
```

3. **安装依赖**

```bash
go mod tidy
```

4. **运行项目**

```bash
go run main.go
```

默认情况下，服务将在 http://localhost:9001 启动。如果需要使用自定义配置文件，可以通过命令行参数指定：

```bash
go run main.go /path/to/config.ini
```

## API接口说明

### 基础路径

所有API接口都以 `/v1` 开头

### 待办事项接口

| 方法 | 路径 | 功能 |
|------|------|------|
| GET  | /    | 查看主页面 |
| GET  | /archive | 查看存档页面 |
| POST | /v1/todo | 创建新的待办事项 |
| GET  | /v1/todo | 获取所有活跃的待办事项 |
| PUT  | /v1/todo/:id | 更新指定ID的待办事项 |
| DELETE | /v1/todo/:id | 删除指定ID的待办事项 |
| POST | /v1/todo/:id/archive | 将指定ID的待办事项存档 |
| GET  | /v1/todo/archived | 获取所有已存档的待办事项 |
| GET  | /v1/todo/stats | 获取待办事项统计数据 |
| GET  | /v1/todo/history | 获取历史统计数据 |

### 请求与响应示例

#### 创建待办事项

**请求**:
```json
POST /v1/todo
Content-Type: application/json

{
  "title": "学习Go语言",
  "status": false
}
```

**响应**:
```json
{
  "id": 1,
  "title": "学习Go语言",
  "status": false,
  "is_archived": false,
  "created_at": "2023-01-01T10:00:00Z",
  "updated_at": "2023-01-01T10:00:00Z",
  "archived_at": null
}
```

## 配置说明

配置文件位于`conf/config.ini`，主要配置项包括：

```ini
port = 9001        # 服务端口
release = false    # 是否为发布模式（true/false）

[mysql]            # 数据库配置
user = root        # 数据库用户名
password = 123456  # 数据库密码
host = 127.0.0.1   # 数据库主机
port = 3306        # 数据库端口
db = checklist     # 数据库名
```

## 部署说明

### 编译项目

```bash
go build -o dailychecklist main.go
```

### 运行编译后的二进制文件

```bash
./dailychecklist
```

### 生产环境配置

在生产环境中，建议将`config.ini`中的`release`设置为`true`，以启用Gin的发布模式，提高性能。

## License

[MIT](https://opensource.org/licenses/MIT)

## 致谢

- [Gin](https://github.com/gin-gonic/gin) - 高性能Go Web框架
- [GORM](https://github.com/jinzhu/gorm) - 优秀的Go ORM库
- [go-ini](https://github.com/go-ini/ini) - Go语言的INI文件解析库