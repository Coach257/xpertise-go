from django.core.wsgi import get_wsgi_application
from django.http import HttpResponse
from django.urls import path
import sys
from django.conf import settings
from scholarly import scholarly
import json
settings.configure(
    DEBUG=True,
    ROOT_URLCONF=__name__,
)


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
    return HttpResponse(json.dumps(res), content_type="application/json")

# 搜索文献（可多篇，指定个数）
def search_pubs(request):
    pub_name = request.POST['pub_name'] # 文献名
    num = int(request.POST['num']) # 希望检索多少篇文献
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



urlpatterns = [
    path('api/search_author', search_author),
    path('api/search_pubs', search_pubs)
]

application = get_wsgi_application()
if __name__ == "__main__":
    from django.core.management import execute_from_command_line
    execute_from_command_line(sys.argv)
