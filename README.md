## 1. What is completed 

Due to time constraints, these are the deliverables:

- GET /api/products?search=&page=&limit=
    - [X] Backend (Completed)
    - [X] Frontend (With no debounce on search, no filter )

## 2. Setup and run

**With Docker** 

```bash
cp .env.example .env    
docker compose up --build
```

App runs on <http://localhost:5173>. 

**Without Docker** 

Run each part from its folder:

| Part     | Setup                                                       | Run                                            |
| -------- | ----------------------------------------------------------- | ---------------------------------------------- |
| Backend  | `cp .env.example .env` in `backend/`                        | `go run ./cmd/server` (port 8081)              |
| Frontend | `cp .env.example .env` and `npm install` in `frontend/`     | `npm run dev` (port 5173)                      |
| Database | Needs DBeaver  | Run the sql files  |