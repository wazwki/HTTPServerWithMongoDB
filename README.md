# How to config MongoDB in golang HTTP-server

## Step 1:

### Install and import driver for connect to MongoDB:

```bash
go get go.mongodb.org/mongo-driver/mongo
```
```go
import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
```

## Step 2:

### Connect to MongoDB:

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

### Create a collection:
```go
var collection *mongo.Collection = Client.Database("testdb").Collection("items")
```

## Step 3:

### Use created client-object for CRUD-operations:

#### Create:
```go
result, err = collection.InsertOne(context.TODO(), data)
```

#### Read:
```go
err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&data)
```

#### Read all:
```go
tmp, err = collection.Find(context.TODO(), nil)
tmp.All(context.TODO(), data)
```

#### Update:
```go
_, err = collection.UpdateByID(context.TODO(), id, data)
```

#### Update many:
```go
_, err = collection.UpdateMany(context.TODO(), bson.M{"_id": ids}, newData)
```

#### Delete:
```go
_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": id})
```

#### Delete many:
```go
_, err = collection.DeleteMany(context.TODO(), bson.M{"_id": ids})
```
