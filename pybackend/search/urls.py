from django.urls import path

from . import views

urlpatterns =[
    path('',views.index,name='index'),
    path('api/search_author',views.search_author,name='search_author'),
    path('api/search_pubs',views.search_pubs,name='search_pubs')
]