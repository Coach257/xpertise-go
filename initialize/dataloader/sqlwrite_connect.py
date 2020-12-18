import pymysql

conn = pymysql.connect( host='101.132.227.56',
                        port=3306,
                        user='root',
                        passwd='@buaa21',
                        charset='utf8',
                        db = 'xpertise_db',)
crusor = conn.cursor()

# create table connections
crusor.execute('CREATE TABLE `connections`( \
               `author1_id` VARCHAR(10) ,\
                `author1_name` VARCHAR(100) ,\
               `author2_id` VARCHAR(10) ,\
                `author2_name` VARCHAR(100) ,\
                `father_id` VARCHAR(10),\
               `paper_id` VARCHAR(10) ,\
               `paper_title` VARCHAR(400))\
               ENGINE = InnoDB,\
               DEFAULT CHARACTER SET = utf8mb4')
conn.commit()

#insert data of connections
f = open('csdata/connect.txt','r', encoding='UTF-8')
index = 0
while True:
    line = f.readline()
    if len(line) == 0 or len(line) ==1:
        break
    data = line[:-1].split('\t')
    # print(data)
    crusor.execute('INSERT INTO connections\
                    VALUES (%s, %s, %s, %s, %s, %s, %s)', (data[0], data[1], data[2], data[3], data[4], data[5], data[6]))
    index += 1
    if not (index % 1000):
        print (("connection",index))
        conn.commit()
f.close()
conn.commit()
