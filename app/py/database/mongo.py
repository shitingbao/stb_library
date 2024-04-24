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
collection = db["open_douyin_video"]
# 插入一条数据
user = {
    "client_key": "gg",
    "open_id": "gg",
    "video_id": "7354986998525971762",
    "account_id": "jinnianjiang",
    "avatar": "https://p3-pc.douyinpic.com/aweme/100x100/aweme-avatar/tos-cn-avt-0015_2a1bde80e55c63839504708b35d9db0a.jpeg",
    "aweme_id": 101439950544,
    "cover": "https://p9-pc-sign.douyinpic.com/image-cut-tos-priv/638f8701049c74addb50ca5bb8d319e3~tplv-dy-resize-origshort-autoq-75:330.jpeg?x-expires=2016090000&x-signature=dfYTH3ATSda%2BkaIpf7h7MFKzpMA%3D&from=3213915784&s=PackSourceEnum_AWEME_DETAIL&se=false&sc=cover&biz_tag=pcweb_cover&l=202311231715111188867F61F7B90437D2",
    "create_time": 1713165747.0,
    "is_reviewed": True,
    "is_top": True,
    "item_id": "@9VwLz/SbSsRjJm+la4EnQs791W3oOPyLP5xzqAqvL1IVZvD560zdRmYqig357zEBQ6FqXeUjNu73aykgLEF16w==",
    "media_type": 4,
    "nickname": "\u91d1\u5e74\u848b\u8001\u677f\u5a18",
    "product_id": "3549646024421326789",
    "product_img": "",
    "product_name": "",
    "share_url": "https://www.iesdouyin.com/share/video/7081549492109036815/?region=CN&mid=7081549541203315486&u_code=16i5b8hc8&did=MS4wLjABAAAANwkJuWIRFOzg5uCpDRpMj4OX-QryoDgn-yYlXQnRwQQ&iid=MS4wLjABAAAANwkJuWIRFOzg5uCpDRpMj4OX-QryoDgn-yYlXQnRwQQ&with_sec_did=1&titleType=title&share_sign=3JRd0MfNphH0XZI5BNKFVLO0mswZ0j7Zp2xw_IAH84c-&share_version=230200&ts=1703901600&from_aid=1128&from_ssr=1",
    "statistics": {
        "comment_count": 506,
        "digg_count": 4322,
        "download_count": 75,
        "forward_count": 0,
        "play_count": 855238,
        "share_count": 429,
    },
    "title": "\u4eca\u5929\u8001\u677f\u5a18\u5e26\u5927\u5bb6\u6765\u770b\u814c\u5236\u4e2d\u7684\u8089\u8089@DOU+\u5c0f\u52a9\u624b @\u8001\u677f\u5a18\u7684\u963f\u5bbf\uff08xue\uff09",
    "video_status": 1,
    "is_update_cover": 1,
}
try:
    collection.insert_one(user)
    print("User inserted successfully.")
except Exception as e:
    print("Failed to insert user:", e)

# # 查询集合中的所有数据
try:
    users = collection.find({"client_key": "gg"})
    print("All users:")
    for user in users:
        user["_id"] = str(user["_id"])
        json_doc = json.dumps(user, indent=4)
        print(json_doc)
except Exception as e:
    print("Failed to fetch users:", e)

try:
    # 更新符合条件的第一个文档
    result = collection.update_one(
        {"client_key": "gg"}, {"$set": {"open_id": "gga", "video_id": "aagg"}}
    )
    print("Document updated:", result)
except Exception as e:
    print("Failed to update document:", e)

try:
    # 删除符合条件的第一个文档
    result = collection.delete_one({"client_key": "gg"})
    print("Document deleted:", result)
except Exception as e:
    print("Failed to delete document:", e)

# 关闭连接
client.close()
