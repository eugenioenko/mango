## Mango

Mango is a minimal expression parser and interpreter

## Usage

```
mango [options] [file]

    --version	prints mango version
    --help		prints this message
```

Example: `mango ./example.mgo`

## Building

```
go build
```

## Adding Expressions and Statement

- Update AST definitions in `./tools/ast.js`
- Run `node ./tools/ast.js` to generate syntactical nodes definitions
- Update `parser.go` to parse the new node
- Update `interpreter.go` to interpret the new node
