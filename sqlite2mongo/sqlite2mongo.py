import sqlite3
from pymongo import MongoClient

sqlite_name = 'D:\work\sqlite\商丘最新办法.pam'
print("sqlite_name: ", sqlite_name)
# 连接MongoDB
client = MongoClient('mongodb://jinrun:123456@localhost:27017/')
db = client['jinrun_db']


# 连接SQLite数据库
try:
    conn = sqlite3.connect(sqlite_name)
    print("conn ok")
except Exception as e:
    print("sqlite3.connect error: ", e)
    exit()

cursor = conn.cursor()

# 查询所有表名
cursor.execute("SELECT name FROM sqlite_master WHERE type='table';")

# 获取查询结果
tables = cursor.fetchall()

# 打印所有表名
for table in tables:
    print("----------table:", table[0])
    cursor.execute(f'SELECT * FROM {table[0]}')
    rows = cursor.fetchall()
    db[table[0]].drop()
    collection = db[table[0]]
    # 对于表中的每一行数据
    for row in rows:
        # 创建一个字典来存储数据
        data = {}
        for idx, column in enumerate(cursor.description):
            print(column[0],'=', row[idx])
            data[column[0]] = row[idx]

        # 将数据插入到 MongoDB 集合中
        collection.insert_one(data)

# 关闭连接
conn.close()
client.close()
