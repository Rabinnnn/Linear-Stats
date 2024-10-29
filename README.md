# LINEAR-STATS

## Description
This is a program built on the previously done 'math-skills' program. It focusses on Linear Regression Line and the Pearson Correlation Coefficient to print various statistical calculations. The program is created using Golang.

## File System

## 1. maths
The directory has 2 files:
* average.go: contains the function calculates the mean of values provided.
* linearRegression.go: contains the function that calculates linear regression.

## 2. data.txt
The file contains sample dataset in the following format:
189
113
121
114
145
110
...

## 3. main.go
This is the entry point for the program. It calculates and displays the Linear Regression Line and Pearson Correlation Coefficient. The output should be in the following format:

Linear Regression Line: y = <value>x + <value>
Pearson Correlation Coefficient: <value>

The values in between the single angle quotation marks (< >) should be a decimal number. The values for the Linear Regression Line should have 6 decimal places, while the Pearson Correlation Coefficient value should have 10 decimal places.

## Execute
* Clone the repo:

```bash
$ git clone https://learn.zone01kisumu.ke/git/rotieno/linear-stats.git
```

* Navigate to the directory:

```bash
$ cd linear-stats
```
* To run the program use the following command:

```bash
$ go run main.go data.txt

``` 
* You can use the following tester to check if your program is working correctly: https://assets.01-edu.org/stats-projects/stat-bin-dockerized.zip

* Once you've downloaded it and extracted the file, navigate to the 'stat-bin' directory, copy the 'bin' folder and paste it in your program's directory.

* Run the tester using the command below:
```bash
$ ./bin/linear-stats
```
Then run your program using:

```bash
$ go run main.go data.txt
``` 
The results should be the same. Repeat the steps 3 more times to test new data sets.

### Author

- [Rabin Otieno](https://github.com/Rabinnnn)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

