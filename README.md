# Statistic

Go program for the maths-skills project in 01Founders that gives a statistical analysis of a list of numbers provided in
a randomly generated file. It returns the **Average** (mean), **Median**, **Variance**, and **Standard
Deviation**.

## How to run

Run the file by typing `go run . <filename>` where filename is the file containing the list of numbers.

## How to test

> To be able to run this script you need to have installed locally
> [`docker-engine`](https://docs.docker.com/engine/install/)

Run the script `./run.sh math-skills`. This will create a file 'data.txt' that contains a randomly generated list of
numbers, and return the expected values of **Average**, **Median**, **Variance**, and **Standard Deviation**.

Run the statistic file with `go run . data.txt` and compare the values.
