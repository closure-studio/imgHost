# 文件上传 API 文档

## 1. 简介
该 API 用于上传文件至服务器，支持自定义文件名或自动生成 UUID 文件名。

---

## 2. 请求方式
- **请求方法**：`POST`
- **接口地址**：`https://object-storage.arknights.app/api/upload`
- **请求头**：
  - `Content-Type: multipart/form-data`

---

## 3. 请求参数

| 参数名   | 类型      | 必填 | 说明 |
|----------|----------|------|--------------------------------|
| file     | `file`   | 是   | 需要上传的文件                 |
| filename | `string` | 否   | 自定义文件名，不传则自动生成 UUID |

---

## 4. 示例请求

### 4.1 上传文件并指定文件名
```bash
curl --request POST \
  --url http://object/api/upload \
  --header 'Content-Type: multipart/form-data' \
  --form file=@file.png \
  --form filename=test.png

------------------------

{
  "code": 1,
  "data": "test.png",
  "message": "upload success"
}


### 4.2 上传文件但不指定文件名（系统自动生成 UUID）
```bash
curl --request POST \
  --url https://object-storage.arknights.app/api/upload \
  --header 'Content-Type: multipart/form-data' \
  --form file=@file.png \

------------------------

{
  "code": 1,
  "data": "5e8b7f8d-92a6-4b3d-9a14-8f2e1c89c0a4",
  "message": "upload success"
}


