saveUsers:
	@echo "Saving users..."
	API_URL=https://run.mocky.io/v3/03d2a7bd-f12f-4275-9e9a-84e41f9c2aae,https://run.mocky.io/v3/87931203-8086-43ef-ba16-4c8903d8fa88 go run cmd/save_users/main.go

searchUsers:
	@echo "Searching users..."
	go run cmd/search_users/main.go --tags enim