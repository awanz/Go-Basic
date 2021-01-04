About Go Language

# Struktur Directory
├── bin
├── pkg
└── src
    └── github.com
        └── awanz
            └── project_name

# Command
- go build file_name.go
- go run file_name.go

# Make variable
var <variable_name> <data_type> , example:
var name string 
and for assignment value you can 
var name string = "awan"
if you want fast make variable can make like
name := awan

# Constant variable
const phi = 3.14

# read input user
3 way you can read input user
1) Scan() : you can for scan number
fmt.Scan(&salary)
2) Scanf() : you can for scan string
fmt.Scanf(&name)
3) Scanln() : you can for scan last number
fmt.Scanln(&number)

# How to read input by user?
import library `fmt`, `os` and `bufio`, and example code like
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)

# Slice is like array but no array
slice[low:high] from low to high
slice[:high] from zero to high
slice[low:] from low to unlimited

# Map
map[KeyTypeData]ValueTypeData
variableMap[key_name]=Value
delete(variableMap, key)