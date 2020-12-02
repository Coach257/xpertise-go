# Xpertise-backend

## 运行

- Django backend
  - Run in `localhost:8000`

```bash
pip install django # default latest version

# django-backend 单页django应用
python search.py runserver

# pybackend # 完整django项目
python manage.py runserver
```

- Flask backend
  - Run in `localhost:8080`

```bash
pip install flask
python main.py
```

- Gin backend
  - Run in `localhost:8080`

```bash
go mod tidy
go run main.go
```

## 代码规范

### 命名

- 每个变量/对象的命名具有语义，不使用a、b这样的命名
- 驼峰式书写，大小写分隔各个单词，如tokenType、functionTable、localParam...

### Router

- 遵循restful规范，我们从中选取4个：GET POST PUT DELETE
  - GET：查询/获取数据，对数据库无修改
  - POST：需要对数据库中添加条目
  - PUT：对若干字段更新，条目数不变
  - DELETE：删除元素

- 对API的命名仿照已有的几个。v1代表v1版本。

### Controller

- 使用`PostForm()`来获取前端传来的数据
- url中如`/user/query/:id`，从url中获取参数的方法：`Param()`。
- 由于默认类型均为string，需要进行类型转换，使用：strconv.ParseUint()，对应得到`uint32`、`uint64`等
- Example:

```go
sid, _ := strconv.ParseUint(c.Param("id"), 0, 64) //:id获取参数
sid, _ := strconv.ParseUint(c.PostForm("sid"), 0, 64)
```

添加含list类型的对象，代码示例如Portal：

```go
func CreateADocument(c *gin.Context) {
	doc := dao.Document{
		DocID:       1,
		AuthorList:  []string{"wyh", "yp"},
		TypeList:    []string{},
		Abstract:    "ahahha",
		CiteList:    []string{"cite1", "cite2"},
		PublishTime: time.Now(),
		Source:      "test",
		Original:    "abcjfdksleinxsijfiengklsjakl",
	}

	if err := server.CreateADocument(&doc); err != nil {
		c.JSON(0, gin.H{"message": err})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}
```

### Server

对于查询，查一个（明确知道仅会返回一个对象的）和查多个的参数写法是不同的：

Example:

```go
// 查一个
func QueryDocumentByID(id uint64) (doc dao.Document) {
	dao.DB.First(&doc, id)
	return doc
}
// 查多个
func QueryAllStudents() (students []*dao.Student) {
	dao.DB.Find(&students)
	return students
}
```

### Model

- 以Document表为例，需要为每一个字段添加在json中的别名（遵循snake命名，全小写，下划线分隔）

```go
// Document Model.
type Document struct {
	DocID       uint64  `gorm:"primary_key;" json: "xxx"`
	AuthorList  StrList `gorm:"type:varchar(200)" json:"author_list"`
	TypeList    StrList `gorm:"type:varchar(200)" json:"type_list"`
	Abstract    string  `gorm:"size:150" json:"xxx"`
	CiteList    StrList `gorm:"type:varchar(500)" json:"cite_list"`
	PublishTime time.Time `json:"xxx"`
	Source      string `gorm:"size:30" json:"xxx"`
	Original    string `gorm:"size:100" json:"xxx"`
}
```

- 另外，现在已可声明外键，欢迎使用！

- 建表语句统一使用`AutoMigrate()`

```go
DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
```

- 若对表进行了修改，如添加了新的字段，原先的数据不会删除，会有空字段。此时需考虑把表drop掉（在datagrip/navicat中，或使用`Drop()`函数等）

## More

- 在搞定swagger之前，每写一个API，请维护API Doc！
- 多在群里交流！