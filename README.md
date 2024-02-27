# Movies Microservices with Go

Application for movie lovers.

## Services

```
                       +----------------------+                            
                       |                      |                            
          +-------------    Movie Service     |-------------+              
          |            |                      |             |              
          |            +----------------------+             |              
+-------------------+                             +-------------------+    
|                   |                             |                   |    
| Metadata Service  |                             |  Rating Service   |    
|                   |                             |                   |    
+-------------------+                             +-------------------+  

```

- **Movie metadata service**: Store and retrieve the movie metadata records by movie IDs.
- **Rating service**: Store ratings for different types of records and retrieve aggregated ratings for records.
- **Movie service**: Provide complete information to the callers about a movie or a set of movies, including the movie metadata and its rating.

## Layers

```
    +--------------+      +----------------+      +-------------+
--->| API Handler  | ---> | Business Logic | ---> | Repository  | ---> DB(postgres)
    +--------------+      +----------------+      +-------------+
                              (controller)
```
