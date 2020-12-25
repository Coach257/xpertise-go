import sys
import json
from pathlib import Path

if __name__ == '__main__':
    
    f_out = open("affiliationCount.json","w+")
    

    paperjson = {}
    authorjson = {}
    affiliationjson = {}
    year_citation = {}
    year_pubs = {}

    
    # 读入paper的数据
    load_data_path = Path('.').joinpath('paperyearcitation.json')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()

    while(line):
        dict_str = json.loads(line)
        paperid = dict_str["id"]
        if paperid not in paperjson:
            paperjson[paperid] = dict_str
        line = read_file.readline()

    # 读入author的数据
    load_data_path = Path('.').joinpath('authorH_index.json')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()

    while(line):
        dict_str = json.loads(line)
        authorid = dict_str["id"]
        if authorid not in authorjson:
            authorjson[authorid] = dict_str
        line = read_file.readline()


    # 读入机构json
    load_data_path = Path('XpertiseData').joinpath('affiliation.json')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()

    while(line):
        dict_str = json.loads(line)
        affiliationid = dict_str["id"]
        if affiliationid not in affiliationjson:
            affiliationjson[affiliationid] = dict_str

        if affiliationid not in year_citation:
            year_citation[affiliationid] = {}
        if affiliationid not in year_pubs:
            year_pubs[affiliationid] = {}
        
        for paper in dict_str["pubs"]:
            paperid = paper["id"]
            papertitle = paper["title"]
            n_citation = 0
            if paperid in paperjson:
                nowpaper = paperjson[paperid]
                # n_citation
                n_citation = nowpaper["n_citation"]
                # year_citation
                for year in nowpaper["citation_by_year"]:
                    if year not in year_citation[affiliationid]:
                        year_citation[affiliationid][year] = 0
                    year_citation[affiliationid][year] = year_citation[affiliationid][year] + nowpaper["citation_by_year"][year]
                # year_pubs
                year = nowpaper["year"]
                if year not in year_pubs[affiliationid]:
                    year_pubs[affiliationid][year] = 0
                year_pubs[affiliationid][year] = year_pubs[affiliationid][year] + 1

            paper["n_citation"] = n_citation

        for author in dict_str["authors"]:
            authorid = author["id"]
            n_citation = 0
            n_pubs = 0
            if authorid in authorjson:
                nowauthor = authorjson[authorid]
                n_citation = nowauthor["n_citation"]
                n_pubs = nowauthor["n_pubs"]
            author["n_citation"] = n_citation
            author["n_pubs"] = n_pubs

        line = read_file.readline()
    

    for itemid in affiliationjson:
        if itemid in year_pubs:
            affiliationjson[itemid]["year_pubs"] = year_pubs[itemid]
        else:
            affiliationjson[itemid]["year_pubs"] = {}
        if itemid in year_citation:
            affiliationjson[itemid]["year_citation"] = year_citation[itemid]
        else:
            affiliationjson[itemid]["year_citation"] = {}
        f_out.write(json.dumps(affiliationjson[itemid])+"\n")
    
    print("hello")
    
    f_out.close()                        