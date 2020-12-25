import sys
import json
from pathlib import Path

if __name__ == '__main__':
    
    f_out = open("paperyearcitation.json","w+")
    

    load_data_path = Path('XpertiseData').joinpath('paper.json')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()

    citation_by_year = {}
    papername = {}
    localjson = {}

    string = ""

    while(line):
        dict_str = json.loads(line)
        papertitle = dict_str["title"]
        paperid = dict_str["id"]
        paperyear = dict_str["year"]
        #author = dict_str["authors"][0]["name"]
        authorlist = dict_str["authors"]
        referencelist = dict_str["reference"]
        if paperid not in localjson:
            localjson[paperid] = dict_str

        if paperid not in papername:
            papername[paperid] = papertitle

        for referenceid in referencelist:
            if referenceid not in citation_by_year:
                citation_by_year[referenceid] = {}
            if paperyear not in citation_by_year[referenceid]:
                citation_by_year[referenceid][paperyear] = 0
            
            citation_by_year[referenceid][paperyear] = citation_by_year[referenceid][paperyear]+1
        
        line = read_file.readline()

    '''
    for itemid in citation_by_year:
        string = str(itemid) + "\t" 
        if itemid in papername:
            string = string +str(papername[itemid])+"\n"
        else:
            string = string +"Can't find paper title"+"\n"
        f_out.write(string)
        for year in citation_by_year[itemid]:
            string = str(year)+"\t" + str(citation_by_year[itemid][year])+"\n"
            f_out.write(string)

    count = 0
    '''
    
    for itemid in localjson:
        if itemid in citation_by_year:
            localjson[itemid]["citation_by_year"] = citation_by_year[itemid]
        else:
            localjson[itemid]["citation_by_year"] = {}
        f_out.write(json.dumps(localjson[itemid])+"\n")
    
    #f_out.write(json.dumps(localjson))
    #print(localjson)
    
    print("hello")
    
    f_out.close()