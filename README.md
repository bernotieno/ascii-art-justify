# ascii-art-justify

ASCII-Art-Justify is a program written in Go that converts input strings into graphic representations using ASCII characters. The ascii representation is then aligned to either right, left, center or justify on the terminal

## Features

- Converts strings into ASCII art
- Supports numbers, letters, spaces, special characters, and newline characters ('\n')
- Utilizes specific graphical templates for ASCII representation
- Align ascii art generated on the terminal.

## Installation

1. Install [go](https://go.dev/doc) to your machine.

2. Clone the repository:

    ```
    git clone https://learn.zone01kisumu.ke/git/togondol/ascii-art-justify
    cd ascii-art-justify
    ```

## Usage

### To generate a right aligned ASCII art for a string with a default font type (standard.txt), run the following command:

```go
 go run . --align=<alignment> "input string"
```
Example:

```go
 go run . --align=right "Hello" 
```

Output:

![alt text](/images/right.png)                                                

### If you want to generate a left aligned Ascii art:

```go
 go run . --align=<alignment> "input string" banner
```
Example:

```go
 go run . --align=left "Hello there" thinkertoy
```

Output:

![alt text](/images/left.png)

### If you want to generate a centered alignment:

```go
 go run . --align=<alignment> "input string" banner
```

Example:

```go
 go run . --align=center "Hello Amigo" shadow
```

Output:

![alt text](/images/center.png)                             

### If you want to generate ascii art with alignment justify:

```go
go run . --align=<alignment> "input string" banner
```

Example:
```go
go run . --align=justify "the quick brown fox" standard
```

Output:
![alt text](/images/justify.png)

## Testing 
To run the tests present use the following command:

```
go test ./Ascii
```

## File Formats

- `standard.txt`: Standard ASCII character set
- `shadow.txt`: Shadowed ASCII character set
- `thinkertoy.txt`: ASCII character set with thinkertoy style

## Contributing

If you have suggestions for improvements, bug fixes, or new features, feel free to open an issue or submit a pull request.

## Author

This project was build and maintained by:

[Thadeus Ogondola](https://learn.zone01kisumu.ke/git/togondol/)

[Bernad Okumu](https://learn.zone01kisumu.ke/git/bernaotieno)

