
1. export PATH=/opt/homebrew/bin:$PATH    
2. goose -dir database/migrations postgres "postgresql://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable" up