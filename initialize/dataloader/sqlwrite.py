import pymysql

conn = pymysql.connect( host='101.132.227.56',
                        port=3306,
                        user='root',
                        passwd='@buaa21',
                        charset='utf8',
                        db = 'xpertise_db',)
crusor = conn.cursor()

#crusor.execute('create database EELAB1 character set UTF8mb4 collate utf8mb4_general_ci')

#create table papers
crusor.execute('CREATE TABLE `papers`( \
               `paper_id` VARCHAR(10) NOT NULL,\
               `title` VARCHAR(400) NULL,\
               `paper_publish_year` VARCHAR(5) NULL,\
               `conference_id` VARCHAR(10) NULL,\
               PRIMARY KEY (`paper_id`),\
               INDEX `ID` USING BTREE (`paper_id`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4',)
conn.commit()

#insert data of papers
f = open('csdata/papers.txt','r', encoding='UTF-8')
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
               `author_id` VARCHAR(10) NOT NULL,\
               `author_name` VARCHAR(100) NULL,\
               PRIMARY KEY (`author_id`),\
               INDEX `ID` USING BTREE (`author_id`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of authors
f = open('csdata/authors.txt','r', encoding='UTF-8')
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
               `conference_id` VARCHAR(10) NOT NULL,\
               `conference_name` VARCHAR(10) NULL,\
               PRIMARY KEY (`conference_id`),\
               INDEX `ID` USING BTREE (`conference_id`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of conferences
f = open('csdata/conferences.txt','r', encoding='UTF-8')
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
               `affiliation_id` VARCHAR(10) NOT NULL,\
               `affiliation_name` VARCHAR(150) NULL,\
               PRIMARY KEY (`affiliation_id`),\
               INDEX `ID` USING BTREE (`affiliation_id`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of affiliations
f = open('csdata/affiliations.txt','r', encoding='UTF-8')
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
crusor.execute('CREATE TABLE `paper_references`( \
               `paper_id` VARCHAR(10) NOT NULL,\
               `reference_id` VARCHAR(10) NOT NULL,\
               PRIMARY KEY (`paper_id`,`reference_id`),\
               INDEX `ID` USING BTREE (`paper_id`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of paper_reference
f = open('csdata/paper_reference.txt','r', encoding='UTF-8')
index = 0
while True:
    line = f.readline()
    if len(line) == 0 or len(line) ==1:
        break
    data = line[:-1].split('\t')
    # print(data)
    crusor.execute('INSERT INTO paper_references\
                    VALUES (%s, %s)', (data[0],data[1]))
    index += 1
    if not (index % 5000):
        print (("paper_reference",index))
        conn.commit()
f.close()
conn.commit()

#create table paper_author_affiliation
crusor.execute('CREATE TABLE `paper_author_affiliations`( \
               `paper_id` VARCHAR(10) NOT NULL,\
               `author_id` VARCHAR(10) NULL,\
               `affiliation_id` VARCHAR(10) NULL,\
               `author_sequence` VARCHAR(3) NOT NULL,\
               PRIMARY KEY (`paper_id`,`author_sequence`),\
               INDEX `ID` USING BTREE (`paper_id`,`author_sequence`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of paper_author_affiliation
f = open('csdata/paper_author_affiliation.txt','r', encoding='UTF-8')
index = 0
while True:
    line = f.readline()
    if len(line) == 0 or len(line) ==1:
        break
    data = line[:-1].split('\t')
    # print(data)
    if data[2] == 'None':
        crusor.execute('INSERT INTO paper_author_affiliations\
                       VALUES (%s, %s, NULL, %s)', (data[0],data[1],data[3]))
    else:
        crusor.execute('INSERT INTO paper_author_affiliations\
                       VALUES (%s, %s, %s, %s)', (data[0],data[1],data[2],data[3]))
    index += 1
    if not (index % 5000):
        print (("paper_author_affiliation",index))
        conn.commit()
f.close()
conn.commit()