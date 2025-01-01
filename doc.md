goose create <migration_name> <sql|go>
goose -v turso libsql://dickens-test-immanu3l.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MzQyMjIyNzgsImlkIjoiMzhkNTFlMWQtNWJjZi00NzFiLTk0ZGItZDZiZDkyNzI2ZDgzIn0.9IJpzSjHZbU2blBB7a_bmvgf5LNmjMYeAZbpK1OvsvS55xT48ToafFji2kISLVPd4EM_mYS8E0d6CYIV7oPvCgdown
go build -buildmode=c-archive 
tailwindcss -i ./main.css -o ../../bin/style.css
tailwindcss -i ./main.css -o ../../bin/style.css --watch
