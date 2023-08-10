# Go Csv Project

## To get the users from API and save them in a csv file
```bash
make saveUsers
```
#### OR
```bash
export API_URL=API_URL # API URL must seperated with comma(,)
go run cmd/save_users/main.go
```

## To search users by tag from saved csv
```bash
make searchUsers
```
#### OR
```bash
go run cmd/search_users/main.go --tags tags # tags must seperated with comma(,)
```