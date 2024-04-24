import mysql.connector
from mysql.connector import errorcode

from py.database.config import getConfig


def getMysqlDb(user, password, host, database):
    # 打开数据库连接
    try:
        db = mysql.connector.connect(
            user=user, password=password, host=host, database=database
        )
        return db

    except mysql.connector.Error as err:
        raise (err)


def ListTsk():
    try:
        db = getMysqlDb(
            user="root", password="4116bbDD#", host="47.99.104.79", database="mengma"
        )

        # 使用cursor()方法获取操作游标
        cursor = db.cursor()

        # 使用execute方法执行SQL语句
        cursor.execute("SELECT * from company_task limit 3")

        # 使用 fetchone() 方法获取一条数据
        data = cursor.fetchall()

        # print("data:", data)

        for da in data:
            print(da, ":", da[1])

        # 关闭数据库连接
        db.close()
    except Exception as e:
        print("Exception db==:", e)


conf = getConfig("./configs/config_example.yaml")

print(conf["data"]["database"])

# ListTsk()
