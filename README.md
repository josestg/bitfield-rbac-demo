# BitField RBAC Demo

### Run the server

```bash
SECRET=myjwtsecret go run main.go
```

### Demo

1. Create a token with permissions `SeeUser` (0), `AddUsers` (1), and `DelUsers` (2)

```bash
curl -X POST --location "http://localhost:8080/token" \
    -H "Content-Type: application/json" \
    -d "{
          \"permissions\": [0, 1, 2]
        }"
```

2. Create a user with the token

```bash
curl -X POST --location "http://localhost:8080/users" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU5MDU0MzgsInBlcm1pc3Npb25zIjpbMCwxLDJdfQ.pV3T00JgFwFexM-JHTePftuS3UMO3kMA2NHZYNivwkI" \
    -d "{
          \"username\": \"test\",
          \"password\": \"test\"
        }"
```

> Should return `200 OK`

3. Try access API that requires permissions that the token does not have

```bash
curl -X GET --location "http://localhost:8080/emails" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU5MDU2MzMsInBlcm1pc3Npb25zIjpbMCwxLDIsNTIsNTNdfQ.LEEuDKjFNdOkyCVzvurZq6foQmhLtnjY2IQwQSM0D3o"
```

> Should return `403 Forbidden` with message `You don't have permission`.