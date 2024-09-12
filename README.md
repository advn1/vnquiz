# vnquiz command line interface app

This is a simple command-line quiz application written in Go. The quiz reads questions and answers from a CSV file and allows the user to answer within a specified time limit. The project supports colored output using the `fatih/color` package.

## Features
- Load quiz questions and answers from a CSV file.
- Set a time limit for the question.
- Colorful.
- Immediate answer checking.

## Usage

### Build
To build the application, you need to have Go installed. Run the following command in the project directory to build the binary:

```bash
go build -o vnquiz
```

This will create an executable file named `vnquiz` (or `vnquiz.exe` on Windows).

### Command-line Arguments

| Flag  | Type   | Description| Default           |
|-------|--------|------------|-------------------|
| `-csv` | string | Path to the CSV file containing quiz questions and answers. | `problems.csv`    |
| `-limit` | int64  | Time limit for each quiz question in seconds.          | `5`               |

### Running the Quiz

You can run the quiz by specifying the CSV file and the time limit for answering the questions. Example command:

```bash
vnquiz -csv=problems.csv -limit=10
```

- The `-csv` flag allows you to specify a custom CSV file path.
- The `-limit` flag sets the time limit for answering each question in seconds.

If no arguments are provided, the default CSV file (`problems.csv`) will be used, and the time limit will be set to 5 seconds.

### CSV File Format

The CSV file should be in the following format:

```
question,answer
5+5,10
10-2,8
Go is developed by?,Google
```

Each line contains a question and its corresponding answer, separated by a comma.

### Example Output

Running the quiz will look like this:

```
Question #1: 5+5 = 10
CORRECT
--------------------------------------------------
Question #2: 10-2 = 9
INCORRECT
--------------------------------------------------
Question #3: Go is developed by? 
Time Out
--------------------------------------------------
Results: 1 out of 3
```

### Custom Usage Message

The application also provides a custom usage message when running the following:

```bash
vnquiz -h
```

The output will display the available flags and their descriptions.

## Dependencies

This project uses the [github.com/fatih/color](https://github.com/fatih/color) package for colorized output. You can install it by running:

```bash
go get github.com/fatih/color
```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/advn1/vnquiz.git
   ```

2. Navigate to the project directory:

   ```bash
   cd vnquiz
   ```

3. Install the dependencies:

   ```bash
   go get github.com/fatih/color
   ```

4. Build the application:

   ```bash
   go build -o vnquiz
   ```

5. Run the quiz:

   ```bash
   ./vnquiz -csv=problems.csv -limit=10
   ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.