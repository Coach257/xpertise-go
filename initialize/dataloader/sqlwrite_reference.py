import pymysql

conn = pymysql.connect( host='101.132.227.56',
                        port=3306,
                        user='root',
                        passwd='@buaa21',
                        charset='utf8',
                        db = 'xpertise_db',)
crusor = conn.cursor()

# create table paper_references
crusor.execute('CREATE TABLE `paper_references`( \
               `paper_id` VARCHAR(10) NOT NULL,\
                `paper_title` VARCHAR(400) NOT NULL,\
               `reference_id` VARCHAR(10) NOT NULL,\
               `reference_paper_title` VARCHAR(400) NOT NULL,\
               INDEX `ID` USING BTREE (`paper_id`))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of paper_references
f = open('csdata/reference.txt','r', encoding='UTF-8')
index = 0
while True:
    line = f.readline()
    if len(line) == 0 or len(line) ==1:
        break
    data = line[:-1].split('\t')
    print(data)
    crusor.execute('INSERT INTO paper_references\
                    VALUES (%s, %s, %s, %s)', (data[0], data[1], data[2], data[3]))
    index += 1
    if not (index % 5000):
        print (("paper_reference",index))
        conn.commit()
f.close()
conn.commit()
