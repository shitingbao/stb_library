from config import getConfig
import pymongo
import json

conf = getConfig("../configs/config_example.yaml")

mongoConf = conf["data"]["mongo"]
print("config file:", conf["data"]["mongo"]["host"])

myclient = pymongo.MongoClient("mongodb://localhost:27017/")
mydb = myclient["runoobdb"]


# 连接 MongoDB 数据库
try:
    # MongoDB 服务器地址
    host = mongoConf["host"]
    # MongoDB 服务器端口
    port = mongoConf["port"]
    # 数据库用户名
    username = mongoConf["username"]
    # 数据库密码
    password = mongoConf["password"]
    # 数据库名称
    dbname = mongoConf["dbname"]
    client = pymongo.MongoClient(host, port, username=username, password=password)
    db = client[dbname]
    print("Connected to MongoDB:", db)
except Exception as e:
    print("Failed to connect to MongoDB:", e)

# 在数据库中创建一个集合（类似于关系型数据库中的表）
collection_name = "open_douyin_video"

# 插入一条数据
# user = {"name": "John", "age": 30, "city": "New York"}
# try:
#     db[collection_name].insert_one(user)
#     print("User inserted successfully.")
# except Exception as e:
#     print("Failed to insert user:", e)

# 查询集合中的所有数据
try:
    users = db[collection_name].find({"client_key": "ff"})
    print("All users:")
    for user in users:
        user["_id"] = str(user["_id"])
        json_doc = json.dumps(user, indent=4)
        print(json_doc)
except Exception as e:
    print("Failed to fetch users:", e)

# 关闭连接
client.close()
