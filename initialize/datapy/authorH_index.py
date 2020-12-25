import sys
import json
from pathlib import Path

if __name__ == '__main__':
    
    f_out = open("authorH_index.json","w+")
    

    paperjson = {}
    authorjson = {}
    h_index = {}
    year_citation = {}
    year_pubs = {}
    paper_with_xcitation_num = {}

    string = ""

    load_data_path = Path('.').joinpath('paperyearcitation.json')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()

    while(line):
        dict_str = json.loads(line)
        paperid = dict_str["id"]
        paperyear = dict_str["year"]
        n_citation = dict_str["n_citation"]
        if paperid not in paperjson:
            paperjson[paperid] = dict_str
        for author in dict_str["authors"]:
            if author["id"] not in year_pubs:
                year_pubs[author["id"]] = {}
            if author["id"] not in year_citation:
                year_citation[author["id"]] = {}
            if author["id"] not in paper_with_xcitation_num:
                paper_with_xcitation_num[author["id"]] = {}
            # year_pubs
            if paperyear not in year_pubs[author["id"]]:
                year_pubs[author["id"]][paperyear] = 0
            year_pubs[author["id"]][paperyear] = year_pubs[author["id"]][paperyear] + 1
            # year_citation
            for citationyear in dict_str["citation_by_year"]:
                if citationyear not in year_citation[author["id"]]:
                    year_citation[author["id"]][citationyear] = 0
                year_citation[author["id"]][citationyear] = year_citation[author["id"]][citationyear] + 1
            # paper_with_xcitation_num
            if n_citation not in paper_with_xcitation_num[author["id"]]:
                paper_with_xcitation_num[author["id"]][n_citation] = 0
            paper_with_xcitation_num[author["id"]][n_citation] = paper_with_xcitation_num[author["id"]][n_citation] + 1

        line = read_file.readline()

    load_data_path = Path('XpertiseData').joinpath('author.json')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()

    while(line):
        dict_str = json.loads(line)
        authorid = dict_str["id"]
        if authorid not in authorjson:
            authorjson[authorid] = dict_str
        line = read_file.readline()

    # count h_index
    for itemid in paper_with_xcitation_num:
        h_index[itemid] = 0
        for n in paper_with_xcitation_num[itemid]:
            if paper_with_xcitation_num[itemid][n]>=n:
                if h_index[itemid]<n:
                    h_index[itemid]= n

    for itemid in authorjson:
        if itemid in year_pubs:
            authorjson[itemid]["year_pubs"] = year_pubs[itemid]
        else:
            authorjson[itemid]["year_pubs"] = {}
        if itemid in year_citation:
            authorjson[itemid]["year_citation"] = year_citation[itemid]
        else:
            authorjson[itemid]["year_citation"] = {}
        authorjson[itemid]["h_index"] = h_index[itemid]
        f_out.write(json.dumps(authorjson[itemid])+"\n")
    
    
    print("hello")
    
    f_out.close()                        