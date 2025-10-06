# DailyChecklist - 日常待办事项管理系统

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.13+-00ADD8?style=flat-square&logo=go)
![Gin Version](https://img.shields.io/badge/Gin-1.7.7-00A86B?style=flat-square&logo=gin)
![GORM](https://img.shields.io/badge/GORM-1.9.16-00C853?style=flat-square)
![MySQL](https://img.shields.io/badge/MySQL-5.7+-4479A1?style=flat-square&logo=mysql)
![License](https://img.shields.io/badge/License-MIT-yellow?style=flat-square)

一个基于 Go 语言和 Gin 框架开发的轻量级待办事项管理系统，旨在帮助你高效地管理日常任务，保持工作与生活的井井有条。

</div>

## ✨ 核心功能

- 📝 **快速创建**：轻松添加新的待办事项，记录每一个灵感。
- 📋 **任务列表**：清晰展示所有待办任务，一目了然。
- ✅ **状态管理**：灵活切换任务的完成与未完成状态。
- 🗑️ **便捷删除**：一键移除不再需要的任务。
- 🗄️ **智能存档**：将已完成的任务归档，保持主列表整洁。
- 📊 **数据统计**：可视化查看任务完成情况与历史趋势。

## 🛠️ 技术栈

- **后端语言**: Go
- **Web 框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **前端**: HTML / CSS / JavaScript

## 🚀 快速体验

1.  **克隆项目**
    `
    bash
    git clone https://github.com/your-username/DailyChecklist.git
    cd DailyChecklist
    `
3.  **配置数据库** (修改 `conf/config.ini` 文件中的连接信息)
4.  **安装依赖并运行**
    `
    bash
    go mod tidy && go run main.go
    `
6.  **访问应用**: 打开浏览器访问
    `  
    http://localhost:9001
    `

## 📄 许可证

本项目基于 [MIT](LICENSE) 许可证开源，欢迎贡献代码或提出建议。
