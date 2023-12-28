
# Strong Password Recommendation steps Backend





## Deployment

To deploy this project run

1. Clone this repository

  ```bash
   git clone https://github.com/your-username/agnos-backend.git
   cd agnos-backend
   cd app
   ```
2. Build and run the Docker containers:

```bash
docker-compose up  
```


The server is located at http://localhost:8080

## API Endpoints
```bash 
/api/strong_password_steps
```

Example: /api/strong_password_steps

Method: POST

Request Body:
```bash
{
  "init_password": "A3a"
}
```
Response:
```bash
{
  "num_of_steps": 3
}
```

## Running test
```bash
go test test/unit_test.go
```

## Access to database
1. Login to PgAdmin: http://localhost:5050

```
Email: admin@admin.com

Password: root
```

2. Add New Server by adding these following field:
```
Name: test_db
Hostname/address: postgres
Username: root
Password: root
```
