# xpertise-scholar

## Go 后端

### 运行

- 进入到`initialize/gorm.go`对数据库配置进行相应修改
- 初次运行/导入第三方module时，通过`go mod tidy`安装项目package依赖
- 运行：`go run main.go`

### 普通部署

#### 准备

- Google Cloud Platform (下简称GCP) 账号
- `app.yaml` (存放于项目根目录)

```yaml
runtime: go115
```

- `cloudbuild.yaml` (存放于项目根目录)

```yaml
steps:
- name: "gcr.io/cloud-builders/gcloud"
  args: ["app", "deploy"]
timeout: "1600s"
```

#### 在GCP中创建Project

填写Project相关信息，此处go-service为该go后端项目。

#### 在Cloud Shell中手动部署

- 创建项目App Engine，选择对应环境(Go)
- 打开Cloud Shell
- `git clone your_repo`
- `cd your_repo`
- `go mod tidy` 安装依赖
- `gcloud app deploy` 部署

### CI/CD 部署

> 使用Cloud Build服务，持续集成，持续部署

- 进入到Cloud Build/Triggers，Create trigger
  - Name: xxx
  - Description: xxx
  - Event: Push to a branch
  - Source: Connect repository(GitHub)-Select your_repo
- 在Cloud Build/Settings中，ENABLE App Engine & Service Accounts

## Vue 前端

### 运行

- `git clone xpertise-frontend(our vue frontend repo)`
- `cd xpertise-frontend`
- `npm install` 安装环境
- `npm run serve` 运行

### 部署

与Go 后端类似。修改的地方如下：

- `app.yaml`

```yaml
runtime: nodejs10
handlers:
  # Serve all static files with urls ending with a file extension
- url: /(.*\..+)$ 
  static_files: dist/\1
  upload: dist/(.*\..+)$
  # catch all handler to index.html
- url: /.*
  static_files: dist/index.html
  upload: dist/index.html
```

- `cloudbuild.yaml`

```yaml
steps:
- name: node:10.15.1
  entrypoint: npm
  args: ["install"]
- name: node:10.15.1
  entrypoint: npm
  args: ["run", "build"]
- name: "gcr.io/cloud-builders/gcloud"
  args: ["app", "deploy"]
timeout: "1600s"
```

- App Engine环境为Node.js
