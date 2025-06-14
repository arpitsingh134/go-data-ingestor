package config

type Config struct {
    APIUrl     string
    MongoURI   string
    Database   string
    Collection string
}

func Load() Config {
    return Config{
        APIUrl:     "https://jsonplaceholder.typicode.com/posts",
        MongoURI:   "mongodb://mongo:27017",
        Database:   "logdb",
        Collection: "posts",
    }
}
