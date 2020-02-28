go install -buildmode=shared std
go install -buildmode=shared -linkshared ./a
go build -o main -linkshared .
