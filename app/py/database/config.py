import yaml


def getConfig(fileName):
    with open(fileName, "r") as f:
        res = yaml.load(f, Loader=yaml.FullLoader)
        return res
