import sys
import json
import tqdm

sys.path.append('../')
from pathlib import Path

data_path = Path('data')
save_data_path = Path('XpertiseData')

def load_affiliations():
    load_data_path = data_path.joinpath('affiliations.txt')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()
    data = {}
    while(line):
        linedata = line.strip().split('\t')
        data[linedata[0]] = linedata[1]
        # jsondata = {
        #     "id":linedata[0],
        #     "name":linedata[1],
        #     "authors":[],
        #     "n_pubs":0,
        #     "n_citation":0,
        #     "pubs":[],
        # }
        line = read_file.readline()
    #print(line)
    return data

def load_authors():
    load_data_path = Path('XpertiseData').joinpath('author.json')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()
    data = []
    while(line):
        data.append(json.loads(line))
        line = read_file.readline()
    return data

def load_conferences():
    load_data_path = data_path.joinpath('conferences.txt')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()
    data = []
    while(line):
        linedata = line.strip().split('\t')
        jsondata = {
            "id":linedata[0],
            "name":linedata[1]
        }
        data.append(jsondata)
        line = read_file.readline()
    return data

def load_paper_author_affiliation():
    load_data_path = data_path.joinpath('paper_author_affiliation.txt')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()
    data = []
    while(line):
        linedata = line.strip().split('\t')
        jsondata = {
            "paperid":linedata[0],
            "authorid":linedata[1],
            "affiliationid":linedata[2],
            "authorsequence":linedata[3],
        }
        data.append(jsondata)
        line = read_file.readline()
    return data

def load_paper_reference():
    load_data_path = data_path.joinpath('paper_reference.txt')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()
    data = []
    while(line):
        linedata = line.strip().split('\t')
        jsondata = {
            "paperid":linedata[0],
            "referenceid":linedata[1]
        }
        data.append(jsondata)
        line = read_file.readline()
    return data

def load_papers():
    data = {}
    load_data_path = Path('XpertiseData').joinpath('paper.json')
    read_file = open(load_data_path,"r",encoding="utf-8")
    line = read_file.readline()
    while(line):
        jsondata = json.loads(line)
        data[jsondata['id']]={
            'title':jsondata['title'],
            'n_citation':jsondata['n_citation']
        }
        line = read_file.readline()
    return data

def can_add_author(add_id, author_list):
    for author in author_list:
        if author['id'] == add_id:
            return False
    return True

def can_add_paper(add_id, paper_list):
    for paper in paper_list:
        if paper['id'] == add_id:
            return False
    return True

def run():
    affiliations_data = load_affiliations()
    authors_data = load_authors()
    # conferences_data = load_conferences()
    # paper_author_affiliation_data = load_paper_author_affiliation()
    # paper_reference_data = load_paper_reference()
    # paper_data = load_papers()

    out_paper_data = []
    data_name = 'newauthor.json'
    out_data_path = save_data_path.joinpath(data_name)
    with open(out_data_path,'w',encoding='utf-8') as fout:
        for data in authors_data:
            new_orgs = []
            for id in data['orgs']:
                new_orgs.append({
                    'id':id,
                    'name':affiliations_data[id]
                })
            data['orgs'] = new_orgs
            fout.write(json.dumps(data)+"\n")
        

    


if __name__ == '__main__':
    run()
    