# Instructions

## Request count per Target graph

- Extract the `logs.tar.gz` file to in the project folder.
- Install the required dependencies by `pip install -r requirements.txt`

  - Before you do this, it is recommended to [setup and activate a Python virtualenv](https://uoa-eresearch.github.io/eresearch-cookbook/recipe/2014/11/26/python-virtual-env/).

- Run `python request_count_per_target.py`.
- The graph will be generated as an `png` image in the project folder.

## Outline

The `parser.py` file contains code to parse the files, extract the relevant information, and sort them on the basis of timestamps.

The `request_count_per_target.py` file uses the parse function to get the required data and generates the graph using Python's `matplotlib` library.

## The OLD folder

Initially I tried to do this using Golang, and the `OLD` folder contains the code from this attempt. It generates the request count per status code graphs and the latency percentile graphs, but as bar graphs and not line graphs.

This code is not optimized and works only on a single log file named `log.log` placed in that folder.

To run this code, the steps are:

- `cd` into OLD folder.
- Copy any single log file into the folder and rename it `log.log`.
- Run `go build --mod=vendor` to generate a binary called `logs_analysis`
- Run `./logs_analysis` to generate the graphs as images in the same folder.

The reasons why I switched from Go to Python was the plotting libraries for Golang are very underdeveloped and less flexible compared to python.

In fact, I had to modify the original [Plot](https://github.com/gonum/plot) library to include labels for barcharts, which is the reason why the dependecies are vendored. (I did submit a [PR upstream](https://github.com/gonum/plot/pull/585) to fix this so hopefully no one has to do the same in the future.)
