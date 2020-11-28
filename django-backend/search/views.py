from django.shortcuts import render
from django.http import HttpResponse
from scholarly import scholarly
import json
# Create your views here.

def index(request):
    return HttpResponse("Xpertise Scholar")

# 搜索作者，返回作者的个人信息
def search_author(request):
    # author_name = request.POST['author_name']
    author_name = 'Mark'
    author = next(scholarly.search_author(author_name))
    res = {'affiliation': author.affiliation,
           'citedby': author.citedby,
           'email': author.email,
           'filled': author.filled,
           'id': author.id,
           'interests': author.interests,
           'name': author.name,
           'url_picture': author.url_picture}
    return HttpResponse(json.dumps(res), content_type="application/json")

# 搜索文献（可多篇，指定个数）
def search_pubs(request):
    # pub_name = request.POST['pub_name'] # 文献名
    # num = int(request.POST['num']) # 希望检索多少篇文献
    pub_name = 'Math'
    num = 3
    i=0
    res=[]
    for pub in scholarly.search_pubs(pub_name):
        if(i==num): break 
        content = {'bib': pub.bib,
            'citations_link': pub.citations_link,
            'filled': pub.filled,
            'source': pub.source,
            'url_add_sclib': pub.url_add_sclib,
            'url_scholarbib': pub.url_scholarbib}
        res.append(content)
        i+=1
    return HttpResponse(json.dumps(res), content_type="application/json")