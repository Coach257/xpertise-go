import pymysql

conn = pymysql.connect( host='101.132.227.56',
                        port=3306,
                        user='root',
                        passwd='@buaa21',
                        charset='utf8',
                        db = 'xpertise_db',)
crusor = conn.cursor()
#
#crusor.execute('create database EELAB1 character set UTF8mb4 collate utf8mb4_general_ci')
#create table papers
crusor.execute('CREATE TABLE `papers`( \
               `PaperID` VARCHAR(10) NOT NULL,\
               `Title` VARCHAR(400) NULL,\
               `PaperPublishYear` VARCHAR(5) NULL,\
               `ConferenceID` VARCHAR(10) NULL,\
               PRIMARY KEY (`PaperID`),\
               INDEX `ID` USING BTREE (`PaperID`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4',)
conn.commit()

#insert data of papers
f = open('data/papers.txt','r', encoding='UTF-8')
index = 0
while True:
    line = f.readline()
    if len(line) == 0 or len(line) ==1:
        break
    data = line[:-1].split('\t')
    # print(data)
    crusor.execute('INSERT INTO papers\
                    VALUES (%s, %s, %s, %s)', (data[0],data[1],data[2],data[3]))
    index += 1
    if not (index % 5000):
        print (("paper",index))
        conn.commit()
f.close()
conn.commit()



#create table authors
crusor.execute('CREATE TABLE `authors`( \
               `AuthorID` VARCHAR(10) NOT NULL,\
               `AuthorName` VARCHAR(100) NULL,\
               PRIMARY KEY (`AuthorID`),\
               INDEX `ID` USING BTREE (`AuthorID`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of authors
f = open('data/authors.txt','r', encoding='UTF-8')
index = 0
while True:
    line = f.readline()
    if len(line) == 0 or len(line) ==1:
        break
    data = line[:-1].split('\t')
    # print(data)
    crusor.execute('INSERT INTO authors\
                    VALUES (%s, %s)', (data[0],data[1]))
    index += 1
    if not (index % 5000):
        print (("authors",index))
        conn.commit()
f.close()
conn.commit()

#create table conferences
crusor.execute('CREATE TABLE `conferences`( \
               `ConferenceID` VARCHAR(10) NOT NULL,\
               `ConferenceName` VARCHAR(10) NULL,\
               PRIMARY KEY (`ConferenceID`),\
               INDEX `ID` USING BTREE (`ConferenceID`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of conferences
f = open('data/conferences.txt','r', encoding='UTF-8')
index = 0
while True:
    line = f.readline()
    if len(line) == 0 or len(line) ==1:
        break
    data = line[:-1].split('\t')
    # print(data)
    crusor.execute('INSERT INTO conferences\
                    VALUES (%s, %s)', (data[0],data[1]))
    index += 1
    if not (index % 5000):
        print (("conferences",index))
        conn.commit()
f.close()
conn.commit()


#create table affiliations
crusor.execute('CREATE TABLE `affiliations`( \
               `AffiliationID` VARCHAR(10) NOT NULL,\
               `AffiliationName` VARCHAR(150) NULL,\
               PRIMARY KEY (`AffiliationID`),\
               INDEX `ID` USING BTREE (`AffiliationID`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of affiliations
f = open('data/affiliations.txt','r', encoding='UTF-8')
index = 0
while True:
    line = f.readline()
    if len(line) == 0 or len(line) ==1:
        break
    data = line[:-1].split('\t')
    # print(data)
    crusor.execute('INSERT INTO affiliations\
                    VALUES (%s, %s)', (data[0],data[1]))
    index += 1
    if not (index % 5000):
        print (("affiliations",index))
        conn.commit()
f.close()
conn.commit()

#create table paper_reference2
crusor.execute('CREATE TABLE `paper_reference2`( \
               `PaperID` VARCHAR(10) NOT NULL,\
               `ReferenceID` VARCHAR(10) NOT NULL,\
               PRIMARY KEY (`PaperID`,`ReferenceID`),\
               INDEX `ID` USING BTREE (`PaperID`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of paper_reference
f = open('data/paper_reference.txt','r', encoding='UTF-8')
index = 0
while True:
    line = f.readline()
    if len(line) == 0 or len(line) ==1:
        break
    data = line[:-1].split('\t')
    # print(data)
    crusor.execute('INSERT INTO paper_reference2\
                    VALUES (%s, %s)', (data[0],data[1]))
    index += 1
    if not (index % 5000):
        print (("paper_reference",index))
        conn.commit()
f.close()
conn.commit()

#create table paper_author_affiliation
crusor.execute('CREATE TABLE `paper_author_affiliation`( \
               `PaperID` VARCHAR(10) NOT NULL,\
               `AuthorID` VARCHAR(10) NULL,\
               `AffiliationID` VARCHAR(10) NULL,\
               `AuthorSequence` VARCHAR(3) NOT NULL,\
               PRIMARY KEY (`PaperID`,`AuthorSequence`),\
               INDEX `ID` USING BTREE (`PaperID`,`AuthorSequence`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of paper_author_affiliation
f = open('data/paper_author_affiliation.txt','r', encoding='UTF-8')
index = 0
while True:
    line = f.readline()
    if len(line) == 0 or len(line) ==1:
        break
    data = line[:-1].split('\t')
    # print(data)
    if data[2] == 'None':
        crusor.execute('INSERT INTO paper_author_affiliation\
                       VALUES (%s, %s, NULL, %s)', (data[0],data[1],data[3]))
    else:
        crusor.execute('INSERT INTO paper_author_affiliation\
                       VALUES (%s, %s, %s, %s)', (data[0],data[1],data[2],data[3]))
    index += 1
    if not (index % 5000):
        print (("paper_author_affiliation",index))
        conn.commit()
f.close()
conn.commit()