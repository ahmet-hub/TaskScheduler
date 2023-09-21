package operations

import (
	"regexp"
)

func GetOperation(argv []string) Op {
	return parseArgs(argv)
}

func parseArgs(argv []string) Op {
	if len(argv) == 0 || (argv[0] == "--help" || argv[0] == "-h") {
		return HelpOp{}
	}

	if len(argv) == 1 && isNumber(argv[0]) {
		return ListOp{}
	}

	if len(argv) == 2 {
		isNumber := isNumber(argv[0])
		switch argv[1] {
		case "-s":
			if isNumber {
				return CreateOp{Id: argv[0]}
			}
		case "-f":
			if isNumber {
				return UpdateOp{Id: argv[0]}
			}
		case "-d":
			if isNumber {
				return DeleteOp{Id: argv[0]}
			}
		}
	}

	if argv[0] == "-ls" {
		operation := ListOp{}
		for i := 1; i < len(argv); i++ {
			arg := argv[i]
			switch arg {
			case "-s", "-f":
				operation.Type = arg
			case "-asc", "-desc":
				operation.Order = arg
			default:
				if isNumber(arg) {
					operation.Search = arg
				} else {
					return UnknownOp{}
				}
			}
		}
		return operation
	}

	return UnknownOp{}
}

func isNumber(value string) bool {
	pattern := `^\d+$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(value)
}
