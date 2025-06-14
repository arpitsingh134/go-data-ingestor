# Go Data Ingestion Service

It fetches log data from a public API, applies transformation, and stores it in a **MongoDB** database running inside a Docker container.

It also includes:
- Dockerized setup
- Unit testing
- Mongo Express UI for easy database inspection
- MongoDB shell access via Docker

---

## 🧱 Project Structure

```bash
.
├── cmd/
│   └── main.go                  # Entry point
├── config/
│   └── config.go                # MongoDB and app config
├── internal/
│   ├── fetcher/                 # API fetching logic	
│   ├── transformer/             # Data transformation
│   ├── repository/              # MongoDB interaction
│   └── models/                  # Data model definitions
├── test/                        # Unit tests
├── go.mod
├── Dockerfile
├── docker-compose.yml
└── README.md
````

---

## 🚀 Getting Started

### 🔧 Prerequisites

* [Git](https://git-scm.com/)
* [Docker](https://www.docker.com/)
* [Docker Compose](https://docs.docker.com/compose/)

### 🧪 Run the Application

```bash
git clone https://github.com/your-username/go-data-ingestor.git
cd go-data-ingestor
docker-compose up --build
```

---

## ✅ Run Tests

```bash
go test ./test/...
```

---

## 🗄️ MongoDB Details

* MongoDB URI: `mongodb://localhost:27017`
* Database: `logdb`
* Collection: `posts`

---

## 📊 Viewing Ingested Data

### 🔹 Option 1: Using Mongo Express (Browser UI)

1. Run the app using Docker Compose:

   ```bash
   docker-compose up --build
   ```

2. Open your browser:

   ```
   http://localhost:8081
   ```

3. You’ll see the Mongo Express dashboard.

4. Navigate to:

   * Database: `logdb`
   * Collection: `posts`

5. Click to view the ingested JSON documents.

---

### 🔹 Option 2: Using Mongo Shell in Docker

1. Open terminal access to MongoDB:

   ```bash
   docker exec -it go-data-ingestor-mongo mongosh
   ```

2. Switch to your database:

   ```js
   use logdb
   ```

3. View documents:

   ```js
   db.posts.find().pretty()
   ```

---

## 🔁 API Endpoint Used

* Source: [`https://jsonplaceholder.typicode.com/posts`](https://jsonplaceholder.typicode.com/posts)

---

## 🔄 Transformation Logic

Each record from the API is transformed to include:

* `ingested_at`: UTC timestamp at time of ingestion
* `source`: `"placeholder_api"`

---

## 🧪 Trade-offs & Improvements

* Focused on local-first
* Could extend to cloud-native storage (e.g., AWS S3 or DynamoDB)
* Additional observability and retry logic could be added
* Uses simple MongoDB schema for flexibility and fast prototyping

---

## 📬 Questions?

Feel free to reach out or open issues if you get stuck or want to contribute!

