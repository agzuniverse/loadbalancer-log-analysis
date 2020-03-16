import gzip
import os
from datetime import datetime


def getFileNamesWithDateTime(filenames):
    data = []
    for filename in filenames:
        d = filename.split("_")[4]
        dt = datetime(int(d[:4]), int(d[4:6]), int(d[6:8]), int(d[9:11]), int(d[11:13]))
        data.append({"filename": filename, "datetime": dt})
    return data


def parser():
    data = []
    filenames = os.listdir("logs/")
    filenames_with_datetime = getFileNamesWithDateTime(filenames)
    filenames_with_datetime.sort(key=lambda f: f["datetime"])

    currFileData = []
    prevTimestamp = filenames_with_datetime[0]["datetime"]
    for f in filenames_with_datetime:
        if f["datetime"] != prevTimestamp:
            data.append(currFileData)
            currFileData = []
            prevTimestamp = f["datetime"]
        with gzip.open("logs/" + f["filename"], "r",) as logs:
            for line in logs:
                params = line.split()
                if len(params) > 1:
                    curr = {
                        "timestamp": params[1].decode("utf-8").split("T")[1][:-8],
                        "target_processing_time": float(params[6]),
                        "target": params[4].decode("utf-8"),
                        "responseCode": int(params[8]),
                    }
                    currFileData.append(curr)
    data.append(currFileData)
    return data, filenames_with_datetime

