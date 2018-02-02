# cli-table [![Build Status](https://travis-ci.org/benweidig/cli-table.svg?branch=master)](https://travis-ci.org/benweidig/cli-table)

cli-table is a Go library for easy table-formatted output for CLIs.

## Example

Please see [example/main.go](example/main.go) for a full example.

```go
table := clitable.New()
table.AddRow("COL1", "COL2", "COL3", "COL4")
table.AddRow("This is", "the first", "row?", 42)
table.AddRow("A", "shorter", "row")
table.AddRow("And an even", "longer", "one", "with", "more", "columns than the first")

fmt.Println(table.String())
```

Output:
```sh
COL1        | COL2      | COL3 | COL4 |      |
This is     | the first | row? | 42   |      |
A           | shorter   | row  |      |      |
And an even | longer    | one  | with | more | columns than the first
```

## Options

You can configure the output by modifying the table struct:

| Option                 | Type   | Default | Description                                        |
| ---------------------- | ------ | ------- | -------------------------------------------------- |
| ColSeparator           | string | " \| "   | Vertical separator between columns                 |
| AdditionalRightPadding | int    | 0       | Adds an additional padding to the next column      |
| OuterRightBorder       | bool   | false   | Enable the outer right border                      |
| HeaderBorder           | bool   | false   | Enable the horizontal border beneath the first row |
| HeaderSeparator        | string | '-'     | The byte used for the horizontal header border     |

## Get cli-table

```sh
$ go get -v github.com/benweidig/cli-table
```

## Inspired by

This library was created after using [https://github.com/gosuri/uitable](https://github.com/gosuri/uitable). First I
wanted to fork it and change it to my needs but then I decided I wanted to go another way and startet my own library.

## License

MIT. See [LICENSE](LICENSE).
