import sys
import json
from pathlib import Path

fa = {}

def findfa(x):
    if fa[x]==x:
         return x
    fa[x]=findfa(fa[x])
    return fa[x]

if __name__ == '__main__':
    
    f_out = open("connect3.txt","w+")
    

    load_data_path = Path('XpertiseData').joinpath('paper.json')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()
    connect = {}
    
    data = {}
    withpaper = {}

    string=""

    while(line):
        dict_str = json.loads(line)
        papertitle = dict_str["title"]
        paperid = dict_str["id"]
        #author = dict_str["authors"][0]["name"]
        authorlist = dict_str["authors"]

        
        for author in authorlist:
            if author["id"] not in fa:
                fa[author["id"]] = author["id"]
            if author["id"] not in connect:
                connect[author["id"]]=[]
            if author["id"] not in data:
                data[author["id"]]=author
            if author["id"] not in withpaper:
                withpaper[author["id"]]=[]
            for author2 in authorlist:
                if author2["id"] not in data:
                    data[author2["id"]]=author2
                if author2["id"] not in fa:
                    fa[author2["id"]] = findfa(author["id"])
                else:
                    fa[findfa(author2["id"])] = findfa(author["id"])
                    fa[author2["id"]] = findfa(author2["id"])
                if (author["id"]!=author2["id"])and(author2["id"] not in connect[author["id"]]):
                    connect[author["id"]].append(author2["id"])
                    withpaper[author["id"]].append([author2["id"],paperid,papertitle])
            #print(author)
        line = read_file.readline()
    #print(connect)
    #print(fa)
    #print(withpaper)

    count = 0

    for authorA in connect:
        for authorB in connect[authorA]:
            fa[authorB] = findfa(authorB)
            fa[authorA] = findfa(authorA)
            fa[fa[authorB]] = findfa(authorA)
            

    for authorA in withpaper:
        for another in withpaper[authorA]:
            authorB = another[0]
            if (authorA<authorB[0]):
                #string = string+str(authorA)+"\t"+str(data[authorA]["name"])+"\t"+str(authorB)+"\t"+str(data[authorB]["name"])+"\t"+str(findfa(authorA))+"\t"+str(another[1])+"\t"+str(another[2])+"\n"
                string = str(authorA)+"\t"+str(data[authorA]["name"])+"\t"+str(authorB)+"\t"+str(data[authorB]["name"])+"\t"+str(findfa(authorA))+"\t"+str(another[1])+"\t"+str(another[2])+"\n"
                count = count +1 
                if fa[authorA] != fa[authorB]:
                    print(count)
                f_out.write(string)
                #print(authorA,data[authorA]["name"],authorB,data[authorB]["name"],fa[authorA],another[1],another[2],sep="\t")
 
    f_out.write(string)
    
    print("hello")
    
    f_out.close()