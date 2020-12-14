import bisect
import itertools
import sys
import random
import math
import json
from tqdm import tqdm
from elastic_app_search import Client
import time

sys.path.append('../XpertiseData/')
from pathlib import Path


client = Client('e633da92fd5c4462a16e73f4acb7e1b5', 'private-u2w9ujnfomnvitzmxejtiqci')
engine_name = 'cspaper'


def run():
    infile = './XpertiseData/paper.json'
    readfile = open(infile,"r")
    line = readfile.readline()
    while(line):
        documents = []
        num = 0
        while num < 100:
            jsondata = json.loads(line)
            try:
                jsondata['conference'] = jsondata['venue']
                del(jsondata['venue'])
                num += 1
                documents.append(jsondata)
            except KeyError:
                pass
            except IndexError:
                pass
            finally:
                line = readfile.readline() 
                if not line:
                    break
        
        if documents:
            index_document_results = client.index_documents(engine_name, documents)
            time.sleep(10)
            if not index_document_results:
                print(index_document_results)


if __name__ == '__main__':
    run()
            
            