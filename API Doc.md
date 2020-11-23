# API Doc

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

