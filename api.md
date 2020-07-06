GET /cms/brands
```json

```
GET /cms/brand/{id}
```json
{
    "id": 1,
    "name": "string, 品牌名称",
    "desc": "string, 品牌描述",
    "logo_url": "string, 品牌logo图片"
}
```
POST /cms/brand

**req**
```json
{
  
    "name": "string, 品牌名称",
    "desc": "string, 品牌描述",
    "logo_url": "string, 品牌logo图片"
}

```

**res**
```json
{
    "code": "200",
    "msg": "OK",
    "result": {
    }
}
```
PUT /cms/brand/{id}

**req**
```json
{
  
    "name": "string, 品牌名称",
    "desc": "string, 品牌描述",
    "logo_url": "string, 品牌logo图片"
}
```
**res**
```json
{
    "code": "200",
    "msg": "OK",
    "result": {
    }
}
```
POST /cms/category

**req**

```json
 {
  
    "name": "string, 品牌名称",
    "desc": "string, 品牌描述",
    "pic_url": "string, 分类图片",
    "path": "string, 分类地址{pid}-{child_id}-..."
}
```
**res**
```json
{
    "code": "200",
    "msg": "OK",
    "result": {
    }
}
```
POST /cms/spu
```json

```
**res**
```json
{
    "code": "200",
    "msg": "OK",
    "result": {
    }
}

```

PUT /cms/spu/{id}