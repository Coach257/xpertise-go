# API Doc

## User

### 注册

```go
填写注册信息，创建新用户
request url:api/v1/user/register 
request method:POST
request header:不需要token
request body:
	{
		"username":string,
		"password":string,
		"password2":string,
		"email":string,
		"info":string
	}
response:
	{
		"success":bool,
		"message":string,
	}
```

### 登录

```go
填写登录信息，登录账号
request url:api/v1/user/login 
request method:POST
request header:不需要token
request body:
	{
		"username":string,
		"email":string,
		"password":string
	}
username和email只需填写其一即可
response:
如果成功，返回
	{
   		"success":bool,
		"message":string, 
    	"data":
        		{
            		"token":string,
                    "user_id":int,
                    "username":string,
                    "email":string,
                    "user_type":int,
                    "info":string,
                    "interdiction":bool
        		}
    
	}
user_type 默认为1，表示普通用户；为0则表示管理员
interdiction 默认为false，表示未禁言；为true则表示用户被禁言
如果失败，返回
	{
		"success":bool,
		"message":string,
	}
```

#### 登录状态验证

1. request header中必须包含"token"字段

2. request body中必须包含"user_id"字段（保持大小写、下划线一致）

3. 在router中添加 auth.JwtAuth()

   ```go
   import "xpertise-go/user/auth"
   
   userV1.POST("/reset/account_info", auth.JwtAuth(),userController.ResetAccountInfo)
   ```

4. 如果需要使用token中的信息，请在对应的controller函数中添加

   ```go
   import(
   		"xpertise-go/user/auth"
   		jwtgo "github.com/dgrijalva/jwt-go"
   )	
   
   claims := c.MustGet("claims").(*auth.CustomClaims)
   ```

   即可得到返回的struct指针

   claims中包含的信息如下：

   ```go
   {
   	UserID 	  uint64
   	Username  string
   	Email     string
   	/*还有一些默认字段*/
   }
   ```

   只推荐使用UserID字段，由于token在用户登录时创建，其他字段在用户修改信息后不会得到立刻更新（目前没有在用户修改信息后就立刻重新生成token）

### 修改密码

```go
填写邮箱、密码，填写新密码并确认，修改密码
request url:api/v1/user/reset/password
request method:POST
request header:不需要token
request body:
	{
		"email":string,
		"password":string,
		"new_password":string,
		"new_password2":string,
	}
response:
	{
		"success":bool,
		"message":string,
	}
```

### 修改用户信息

```go
填写相关字段，修改用户信息
request url:api/v1/user/reset/account_info
request method:POST
request header:需要token
request body:
	{
		"user_id":int,
		"username":string,
		"email":string,
		"info":string,
	}
response:
	{
		"success":bool,
		"message":string,
	}
```

### 创建收藏夹

```go
创建收藏夹，用于归类收藏的文献
request url:api/v1/user/reset/password
request method:POST
request header:需要token
request body:
	{
		"user_id":int,
		"username":string,
		"email":string,
		"info":string,
	}
response:
	{
		"success":bool,
		"message":string,
	}
```

### 添加收藏

```go
添加文献到用户收藏夹
request url:api/v1/user/folder/add
request method:POST
request header:需要token
request body:
	{
		"user_id":int,
		"folder_id":int,
		"doc_id":int,
		"doc_info":string,
	}
response:
	{
		"success":bool,
		"message":string,
	}
```


## Branch

### 创建评论
```go
填写评论信息，创建新评论
request url:api/v1/branch/create 
request method:POST
request header:需要token
request body:
	{
		"UserID":int,
		"DocID":int,
		"Content":string,
	}
response:
	{
		"success":bool,
		"message":string,
	}
```


### 删除评论
```go
删除评论
request url:api/v1/branch/delete
request method:DELETE
request header:不需要token
request body:
	{
		
		"id":int,
		
	}
response:
	{
		"success":bool,
		"message":string,
	}
```

### 点赞评论
```go
对具体评论点赞
request url:api/v1/branch/thumbup
request method:POST
request header:需要token
request body:
	{
		"UserID":int,
		"ComID":int,
		
	}
response:
	{
		"success":bool,
		"message":string,
	}
```

### 踩评论
```go
对具体评论踩
request url:api/v1/branch/thumbdown
request method:POST
request header:需要token
request body:
	{
		"UserID":int,
		"ComID":int,
		
	}
response:
	{
		"success":bool,
		"message":string,
	}
```

### 取消赞
```go
取消对具体评论的点赞
request url:api/v1/branch/revert_thumbup
request method:POST
request header:需要token
request body:
	{
		"UserID":int,
		"ComID":int,
		
	}
response:
	{
		"success":bool,
		"message":string,
	}
```

### 取消踩
```go
取消对具体评论的踩
request url:api/v1/branch/revert_thumbdown
request method:POST
request header:需要token
request body:
	{
		"UserID":int,
		"ComID":int,
		
	}
response:
	{
		"success":bool,
		"message":string,
	}
```


### 查找评论的赞总数
```go
获得该评论的赞总数
request url:api/v1/branch/query/thumbup
request method:GET
request header:需要token
request body:
	{
		
		"id":int,
		
	}
response:
	{
		"like":int,
	}
```

### 查找评论的踩总数
```go
获得该评论的踩总数
request url:api/v1/branch/query/thumbdown
request method:GET
request header:需要token
request body:
	{
		
		"id":int,
		
	}
response:
	{
		"dislike":int,
	}
```

## Search

### 搜索作者

```python
# 搜索作者，返回作者的个人信息
def search_author(request):
    author_name = request.POST['author_name']
    author = next(scholarly.search_author(author_name))
    res = {'affiliation': author.affiliation,
           'citedby': author.citedby,
           'email': author.email,
           'filled': author.filled,
           'id': author.id,
           'interests': author.interests,
           'name': author.name,
           'url_picture': author.url_picture}
```

[搜索作者，返回作者的文献信息](https://github.com/fredrike/googlescholar-api)

Example: http://cse.bth.se/~fer/googlescholar-api/googlescholar.php?user=vJjq9LwAAAAJ

从url中获取参数author.id，返回json信息如下：

```json
{
 "total_citations": 58,
 "citations_per_year": {
  "2012 ": 1 ,
  "2013 ": 7 ,
  "2014 ": 13 ,
  "2015 ": 10 ,
  "2016 ": 23 ,
  "2017 ": 2 
 },
 "publications": [
  {
    "title": "Privacy threats related to user profiling in online social networks",
    "authors": "F Erlandsson, M Boldt, H Johnson",
    "venue": "Privacy, Security, Risk and Trust (PASSAT), 2012 International Conference on ..., 2012 ",
    "citations": 18,
    "year": 2012 
  },
  {
    "title": "SIN: A Platform to Make Interactions in Social Networks Accessible",
    "authors": "SFW Roozbeh Nia, Fredrik Erlandsson, Prantik",
    "venue": "ASE International Conference on Social Informatics, 2012 ",
    "citations": 10,
    "year": 2012
  }
 ]
}
```

### 搜索文献

```python
# 搜索文献（可多篇，指定个数）
def search_pubs(request):
    pub_name = request.POST['pub_name'] # 文献名
    num = int(request.POST['num']) # 希望检索多少篇文献

    # 每篇文献所含的内容
    content = {'bib': pub.bib,
        'citations_link': pub.citations_link,
        'filled': pub.filled,
        'source': pub.source,
        'url_add_sclib': pub.url_add_sclib,
        'url_scholarbib': pub.url_scholarbib}
```

