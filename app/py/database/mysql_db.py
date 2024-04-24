import mysql.connector
from mysql.connector import errorcode

from config import getConfig


def getMysqlDb(user, password, host, database):
    # 打开数据库连接
    try:
        db = mysql.connector.connect(
            user=user, password=password, host=host, database=database
        )
        return db

    except mysql.connector.Error as err:
        raise (err)


def create_task(db, name, price, status):
    try:
        cursor = db.cursor()
        sql = "INSERT INTO test (name, price,status) VALUES (%s, %s, %s)"
        val = (name, price, status)
        cursor.execute(sql, val)
        db.commit()
        print("Task created successfully.")
    except Exception as e:
        db.rollback()
        print("Failed to create task:", e)


def delete_task(db, id):
    try:
        cursor = db.cursor()
        sql = "DELETE FROM test WHERE id = %s"
        cursor.execute(sql, (id,))
        db.commit()
        print("Task deleted successfully.")
    except Exception as e:
        db.rollback()
        print("Failed to delete task:", e)


def update_task(db, id, name):
    try:
        cursor = db.cursor()
        sql = "UPDATE test SET name = %s WHERE id = %s"
        val = (name, id)
        cursor.execute(sql, val)
        db.commit()
        print("Task updated successfully.")
    except Exception as e:
        db.rollback()
        print("Failed to update task:", e)


def ListTsk(db):
    try:
        # 使用cursor()方法获取操作游标
        cursor = db.cursor()

        # 使用execute方法执行SQL语句
        cursor.execute("SELECT * from test limit 3")

        # 使用 fetchone() 方法获取一条数据
        data = cursor.fetchall()

        # print("data:", data)

        for da in data:
            print(da)

    except Exception as e:
        print("Exception db==:", e)


def main():
    conf = getConfig("../configs/config_example.yaml")

    mysqlConf = conf["data"]["database"]
    print(conf["data"]["database"])

    try:
        db = getMysqlDb(
            user=mysqlConf["user"],
            password=mysqlConf["password"],
            host=mysqlConf["host"],
            database=mysqlConf["database"],
        )

        # ListTsk(db)
        # create_task(db, "test", 18, 1)
        # update_task(db, 5, "gg")
        delete_task(db, 1)
        db.close()
    except Exception as e:
        print("getMysqlDb db==:", e)

    # 关闭数据库连接


main()
