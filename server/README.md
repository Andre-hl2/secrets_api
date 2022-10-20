# Server

This is the directory with the source code for the project.

## Code Organization
The project is structure as this:

### API
This package is where the API logic layer is kept, its subdivided between two layers

#### - Controllers
This is where the business logic of the application resides, it does not worry about input validation, serializing/deserialing data, or handling the read/write of data from source. This part of the application should have only the single responsiblity of dealing with how the application should work.

#### - Handlers
This is the layer where the validate, serialize and deserialize, and correctly receive and respond to http requests. Currently we're using Go standard `http` package and only handling `REST` calls. 

But by having the handlers and controllers dettached, our application can easily be refactored in the future to have different interfaces with the external world, like JRPC handlers, or messaging queues consumer/producers for event driven handlers.

### Config

This package is responsible for reading configuration options when running the application, it currently only reads values from the `.env` file. But this would be where the application would also handle cloud specific settings (e.g.: AWS, Google Cloud, etc.)

### Storage

This package is where everyhing data specific is kept, other packages will use this one when reading/writing data, but will never worry about how that's done.

The storage code is divided as this:

#### - Entities

The representation of all data types the application will have (User, Secret, etc). We can have different versions of the same entity for different purposes, for instance:
- one struct that will be saved in our data persistance layer with all the fields, like `id`, `created_at`, `updated_at`, etc. 
- one struct that will be used to represent a new instance of that entity, a `NewUser` that does not have yet persistence specific fields

#### - Stores

The application needs to store the entities into `Data sources` so we can read/write different entities (Users, Secrets, etc). This project uses a `Store` interface to make it simple to swap data sources implementations if needed.

Currently the project has `MemoryStore` and a `PostgresStore` implementations, each one of them requiring different setup (external services, env vars, etc), and having different persistance. But both implementing the store interface, making so we can easy change which one we want to use.

By having this abstraction layer, the application does not need to worry about data source specific logic (like DB connection pools), making the code easier to read and to maintain.