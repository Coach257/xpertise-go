# Scholarly

[Project description]([https://pypi.org/project/scholarly/)

[scholarly’s documentation](https://scholarly.readthedocs.io/en/latest/?badge=latest)

[scholarly-python-package](https://github.com/scholarly-python-package/scholarly)

## Scholar 文献

### 返回数据格式

```javascript
'bib':{
        'abstract':'Humans can judge from vision alone whether an object is ',
        'author': 'SA Cholewiak and RW Fleming and M Singh',
        'eprint': 'https://jov.arvojournals.org/article.aspx?articleID=2213254',
        'title': 'Perception of physical stability and center of mass of 3-D '
                'objects',
        'url': 'https://jov.arvojournals.org/article.aspx?articleID=2213254',
        'venue': 'Journal of vision',
        'year': ' 2015'
},
'citedby': 19,
'filled': False,
'id_scholarcitedby': '15736880631888070187',
'source': 'scholar',
'url_scholarbib': 'https://scholar.googleusercontent.com/sch...'
```

### 搜文献

```python
def search_pubs(self,
                    query: str, patents: bool = True,
                    citations: bool = True, year_low: int = None,
                    year_high: int = None):
    """Searches by query and returns a generator of Publication objects"""

search_query = scholarly.search_pubs('Perception of physical stability and center of mass of 3D objects')
                print(next(search_query))
```

### 搜单篇文献

```python
def search_single_pub(self, pub_title: str, filled: bool = False):
        """Search by scholar query and return a single Publication object"""
```

### 根据url搜文献

```python
def search_pubs_custom_url(self, url: str):
        """Search by custom URL and return a generator of Publication objects
        URL should be of the form '/scholar?q=...'"""
        return self.__nav.search_publications(url)
```

## Author 作者

### 返回数据格式

```javascript
{'affiliation': 'Bank of Canada', //所属机构
 'citedby': 28, //总被引次数
 'email': '@bankofcanada.ca',
 'filled': False,
 'id': 'fS-9YqEAAAAJ',
 'interests': ['Macroeconomics', 'Financial Intermediation'], //研究兴趣
 'name': 'Martin Kuncl',
 'url_picture': 'https://scholar.google.com/citations?view_op=medium_photo&user=fS-9YqEAAAAJ'}
```

### 搜作者

```python
def search_author(self, name: str):
        """Search by author name and return a generator of Author objects"""

search_query = scholarly.search_author('Marty Banks, Berkeley')
print(next(search_query))
```

### 根据id搜作者

```python
def search_author_id(self, id: str, filled: bool = False):
        """Search by author id and return a single Author object"""

search_query = scholarly.search_author_id('EmD_lTEAAAAJ')
print(search_query)
```

### 根据关键词搜作者

```python
def search_keyword(self, keyword: str):
        """Search by keyword and return a generator of Author objects"""

search_query = scholarly.search_keyword('Haptics')
print(next(search_query))
```

### 根据url搜作者

```python
def search_author_custom_url(self, url: str):
        """Search by custom URL and return a generator of Author objects
        URL should be of the form '/citation?q=...'"""
        return self.__nav.search_authors(url)
```