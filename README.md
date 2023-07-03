# Server-Side-GO
A minimalistic example of an API built using GO, Mux and MongoDB Atlas.


## Structure
The API in itself it quite simple, basically it performs simple CRUD. Go is the language of choice, MongoDB has been connected using mongo drivers.

To understand better how to integrate mongodb follow the setup according to the [official documentation](https://www.mongodb.com/docs/drivers/go/current/quick-start/)


## Get Started
Add imports for os and github.com/joho/godotenv
Generate MongoDB URI via MongoDB Atlas, follow docs to do so.
Connect as follows;

```
if err := godotenv.Load(); err != nil {
	exceptions.Handle("Error loading .env file")
}

uri := os.Getenv("MONGODB_URI")
if uri == "" {
	exceptions.Handle("MONGO_URI not found")
}

client, err := mongo.Connect(ctxTODO, options.Client().ApplyURI(uri))
if err != nil {
	exceptions.Handle("Error connecting to MongoDB")
}

collection = client.Database("fluttergo-fullstack").Collection("movies")
```


## Author
Me :) reachout if you need assistance with anything. 
