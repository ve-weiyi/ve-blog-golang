# Go 抽象语法树(AST)详解

## 前言

作为一名代码开发程序员，我常常思考一个程序文件，比如 `hello.go` 是如何被编译器理解的.
平常在编写程序时，IDE又是如何提供代码提示的。在这奥妙无穷的背后，AST（Abstract Syntax
Tree）抽象语法树功不可没，它站在每一行程序的身后，默默无闻地工作，为繁荣的互联网世界立下了汗马功劳。

## 编译过程

编译过程是将高级语言源代码转换为目标机器代码或其他中间表示的过程。它通常包括以下几个主要阶段：

### 1. 词法分析（Lexical Analysis）

- **输入**：源代码
- **输出**：词法单元（tokens）
- **任务**：将源代码分解为基本的词法单元，如关键字、标识符、运算符等

### 2. 语法分析（Syntax Analysis）

- **输入**：词法单元序列
- **输出**：语法树（Abstract Syntax Tree，AST）或其他中间表示
- **任务**：根据语法规则检查词法单元序列的结构，构建语法树表示源代码的语法结构

### 3. 语义分析（Semantic Analysis）

- **输入**：语法树或中间表示
- **输出**：语义信息
- **任务**：检查源代码的语义正确性，捕捉并处理语法上合法但语义上错误的结构，生成符号表等

### 4. 中间代码生成（Intermediate Code Generation）

- **输入**：语法树或中间表示
- **输出**：中间代码
- **任务**：将语法树转换为一种中间表示，以便进行后续优化和目标代码生成

### 5. 优化（Optimization）

- **输入**：中间代码
- **输出**：优化后的中间代码
- **任务**：对中间代码进行各种优化，提高程序性能、减小代码体积等

### 6. 目标代码生成（Code Generation）

- **输入**：优化后的中间代码
- **输出**：目标机器代码或汇编代码
- **任务**：将中间代码转换为目标机器代码，根据目标平台生成可执行文件

### 7. 链接（Linking）

- **输入**：目标机器代码或汇编代码，可能包括库文件
- **输出**：可执行文件
- **任务**：将多个目标文件链接在一起，解决符号引用、地址重定位等问题，生成最终的可执行文件

## AST 抽象语法树

AST 指的是 Abstract Syntax Tree（抽象语法树），它是编程语言中源代码语法结构的一种抽象表示形式。AST
通常是一个树形结构，用于表示程序的语法层次关系，从而方便进行语法分析和处理。

在编程语言中，源代码是由一系列的词法单元（token）组成的，而 AST 则将这些词法单元按照语法规则组织成一个树形结构。每个节点代表源代码中的一个语法结构，例如表达式、语句、函数声明等。

### 1. Token（词法单元）

在编程语言中，Token（词法单元）是源代码中的最小语法单元，是编译器或解释器在进行词法分析时识别和处理的基本单位。Token
是源代码经过词法分析器处理后得到的标识符，它代表了源代码中的不可分割的词法结构。

不同类型的编程语言有不同的 Token 类型，一般包括以下几类：

1. **关键字（Keywords）**
    - 表示编程语言的保留字，具有特殊含义
    - 例如：在 C 语言中，`if`、`else`、`while` 等就是关键字

2. **标识符（Identifiers）**
    - 表示变量、函数、类等的名称
    - 标识符需要遵循一定的命名规则
    - 例如：`variableName` 是一个标识符

3. **常量（Literals）**
    - 表示固定值的词法单元，包括整数、浮点数、字符串等
    - 例如：`42`、`3.14`、`"Hello, World!"` 都是常量

4. **运算符（Operators）**
    - 表示执行操作的符号，如加法、减法、乘法等
    - 例如：`+`、`-`、`*` 是运算符

5. **分隔符（Delimiters）**
    - 表示源代码结构的分隔符，如括号、分号、逗号等
    - 例如：`(`、`)`、`;` 是分隔符

6. **注释（Comments）**
    - 表示注释内容，编译器或解释器通常会忽略它们
    - 例如：`// This is a comment`

Token 的产生是由词法分析器（Lexer 或 Scanner）负责的，它扫描源代码，将字符序列组合成有意义的 Token。Token
提供给语法分析器（Parser）使用，用于构建 AST（抽象语法树）等进一步的语法分析和语义分析。

下面是一个简单的示例，展示了一个小型程序的源代码和对应的一些 Token：

```go
// 源代码
x = 10 + 5

// 对应的 Token
Token(IDENT, "x")
Token(ASSIGN, "=")
Token(INT, "10")
Token(PLUS, "+")
Token(INT, "5")
```

在这个例子中：

- `Token(IDENT, "x")` 表示一个标识符 Token
- `Token(ASSIGN, "=")` 表示一个赋值操作符 Token
- `Token(INT, "10")` 表示一个整数常量 Token
- 以此类推

### 2. FileSet

FileSet 是用于跟踪源代码文件和位置信息的结构。它维护源代码文件的集合，每个文件都关联有唯一的标识符，并记录了源代码中各个位置的行号、列号等信息。

```go
// 解析 Go 代码文件
fset := token.NewFileSet()
f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
```

让我们看一下 `parser.ParseFile()` 函数源码：

```go
func ParseFile(fset *token.FileSet, filename string, src any, mode Mode) (f *ast.File, err error) {
if fset == nil {
panic("parser.ParseFile: no token.FileSet provided (fset == nil)")
}

// get source
text, err := readSource(filename, src)
if err != nil {
return nil, err
}

var p parser
defer func () {
if e := recover(); e != nil {
// resume same panic if it's not a bailout
bail, ok := e.(bailout)
if !ok {
panic(e)
} else if bail.msg != "" {
p.errors.Add(p.file.Position(bail.pos), bail.msg)
}
}

// set result values
if f == nil {
// source is not a valid Go source file - satisfy
// ParseFile API and return a valid (but) empty
// *ast.File
f = &ast.File{
Name:  new(ast.Ident),
Scope: ast.NewScope(nil),
}
}

p.errors.Sort()
err = p.errors.Err()
}()

// parse source
p.init(fset, filename, text, mode)
f = p.parseFile()

return
}
```

#### 参数说明：

- `fset *token.FileSet`：用于记录位置信息的文件集合
- `filename string`：源文件的文件名
- `src any`：源代码，可以是字符串、字节数组或 io.Reader
- `mode Mode`：控制解析器的模式，例如是否跳过对象解析阶段

#### 返回值：

- `f *ast.File`：解析后的 AST 树，表示整个源文件的语法结构
- `err error`：解析过程中的错误，如果有的话

#### 内部实现：

- `parser` 结构体是用于实际解析的解析器对象
- `readSource` 函数用于获取源代码文本，它可以从文件或其他来源获取
- `p.parseFile()` 调用实际执行解析过程，返回解析得到的 AST

#### 错误处理：

- 如果源代码无法读取，返回的 AST 是空的，但不为 nil
- 如果源代码读取成功但存在语法错误，返回的 AST 包含 ast.Bad* 节点来表示错误的源代码片段
- 多个错误通过 scanner.ErrorList 返回，按源代码位置排序

#### panic 恢复：

- 解析器在遇到错误时可能会触发 panic，但会通过 recover 恢复，确保程序不会崩溃
- 如果 panic 是 bailout 类型且有错误消息，将错误添加到解析器的错误列表中

### 3. AST 类型

在 Go 语言中，AST（抽象语法树）主要由三个大类别的节点构成，它们分别是：

#### 基本节点（Basic Nodes）：

- `ast.Ident`：代表标识符（Identifiers），如变量名、函数名等
- `ast.BasicLit`：代表基本的字面量，如整数、浮点数、字符串等
- `ast.CompositeLit`：代表复合字面量，如数组、切片、字典等

#### 语句节点（Statement Nodes）：

- `ast.ExprStmt`：代表一个表达式语句，即只包含一个表达式的语句
- `ast.AssignStmt`：代表赋值语句，包括简单的赋值和多重赋值
- `ast.IfStmt`：代表条件语句
- `ast.ForStmt`：代表循环语句
- `ast.SwitchStmt`：代表开关语句（switch）
- `ast.RangeStmt`：代表 for range 循环语句

#### 声明节点（Declaration Nodes）：

- `ast.GenDecl`：代表通用声明，用于表示 import、const、var 等声明
- `ast.FuncDecl`：代表函数声明
- `ast.TypeSpec`：代表类型声明

所有节点都实现了 `ast.Node` 接口，返回了 Node 的位置信息：

```go
type Node interface {
Pos() token.Pos // position of first character belonging to the node
End() token.Pos // position of first character immediately after the node
}

// All expression nodes implement the Expr interface.
type Expr interface {
Node
exprNode()
}

// All statement nodes implement the Stmt interface.
type Stmt interface {
Node
stmtNode()
}

// All declaration nodes implement the Decl interface.
type Decl interface {
Node
declNode()
}
```

以下是 ast 包中包含的节点类型，可以用到的时候查看官方文档：[ast package - go/ast - Go Packages](https://pkg.go.dev/go/ast)
