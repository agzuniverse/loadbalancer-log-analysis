import matplotlib.pyplot as plt
from parser import parser

NO_OF_TARGETS = 42

requestCounts = []
timestamps = []

print("Parsing data to plot graph, please wait.")
data, filenames_with_datetime = parser()
print("Parsing complete.")

for details in data:
    requestCount = len(details)
    requestCountPerTarget = requestCount / NO_OF_TARGETS
    requestCounts.append(requestCountPerTarget)
for f in filenames_with_datetime:
    if f["datetime"].strftime("%H:%M") not in timestamps:
        timestamps.append(f["datetime"].strftime("%H:%M"))

plt.figure(1, figsize=(13, 10), dpi=200)
plt.plot(timestamps, requestCounts)
plt.xlabel("Time")
plt.xticks(rotation=45)
plt.ylabel("Request Count")
plt.title("Request Count Per Target")
plt.savefig("requestCountPerTarget.png")
