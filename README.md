# **How to Set Up MongoDB in a Golang HTTP Server**

# **1. Install and Import the Driver for Connecting to MongoDB:**

```bash
go get go.mongodb.org/mongo-driver/mongo
```

```go
import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
```

# **2. Connect to MongoDB:**

```go
var (
    Client *mongo.Client
    err    error
)

func Connect() {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    Client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = Client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB")
}
```

## **Create a Collection:**

```go
var collection *mongo.Collection = Client.Database("testdb").Collection("items")
```

# **3. Use the Created Client Object for CRUD Operations:**

## **3.1. Create:**

```go
result, err = collection.InsertOne(context.TODO(), data)
```

## **3.2. Read:**

```go
err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data)
```

## **3.3. Read All:**

```go
tmp, err = collection.Find(context.TODO(), nil)
tmp.All(context.TODO(), data)
```

## **3.4. Update:**

```go
_, err = collection.UpdateByID(context.TODO(), id, data)
```

## **3.5. Update Many:**

```go
_, err = collection.UpdateMany(context.TODO(), bson.M{"_id": ids}, newData)
```

## **3.6. Delete:**

```go
_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": id})
```

## **3.7. Delete Many:**

```go
_, err = collection.DeleteMany(context.TODO(), bson.M{"_id": ids})
```