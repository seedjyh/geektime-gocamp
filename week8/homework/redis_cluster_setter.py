# -*- coding: UTF-8 -*-
import getopt
import random
import string
import sys
from rediscluster import StrictRedisCluster


def write_random_string_to_redis_cluster(key_number, value_size):
    """
    连接redis集群，指定数量、指定长度的string value。
    :param key_number: key的数量
    :param value_size: 值的长度
    :return:
    """
    redis_nodes = [{'host': '127.0.0.1', 'port': 8001},
                   {'host': '127.0.0.1', 'port': 8002},
                   {'host': '127.0.0.1', 'port': 8003},
                   {'host': '127.0.0.1', 'port': 8004},
                   {'host': '127.0.0.1', 'port': 8005},
                   {'host': '127.0.0.1', 'port': 8006}]
    try:
        redis_conn = StrictRedisCluster(startup_nodes=redis_nodes)
        for i in range(key_number):
            key = "%010d" % i
            value = generate_random_string(value_size)
            redis_conn.set(key, value)
    except Exception as e:
        print("ERR:", e)
        sys.exit(1)


def generate_random_string(length):
    """
    生成随机字符串。最多前26个字符是随机的。
    :param length: 期望生成的字符串长度
    :return: 随机字符串
    """
    alphabet = string.ascii_lowercase
    if length <= 26:
        return ''.join(random.sample(alphabet, length))
    else:
        return ''.join(random.sample(alphabet, len(alphabet))) + 'a' * (length - 26)


if __name__ == "__main__":
    key_number = None
    value_size = None
    try:
        opts, args = getopt.getopt(sys.argv[1:], "n:d:", ["key_number=", "value_size="])
        for opt, arg in opts:
            if opt in ["-n", "--key_number"]:
                key_number = int(arg)
            elif opt in ["-d", "--value_size"]:
                value_size = int(arg)
    except getopt.GetoptError:
        print("ERR: empty arg value, require: -n 10000 -s 20")
        sys.exit(2)
    except ValueError:
        print("ERR: invalid arg value, require: -n 10000 -s 20")
        sys.exit(2)
    if key_number is None:
        key_number = 10000
        print("Use default key number")
    if value_size is None:
        value_size = 2
        print("Use default value size")

    print(">>", key_number, value_size)
    write_random_string_to_redis_cluster(key_number, value_size)
    print("All setting done!")
