gofmt -d nestStructure.go 
go list -f '{{ .Name }}: {{ .Doc }}'
errcheck ./...
