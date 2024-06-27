
run:
	go run ./cmd/cli/main.go

parser:
	cd ./internal/parser && antlr4 -Dlanguage=Go -visitor -no-listener -package parser SimpleSql.g4

parser-pigeon:
	pigeon -o ./internal/parser_pigeon/parser.go ./internal/parser_pigeon/simpledb_sql.peg

