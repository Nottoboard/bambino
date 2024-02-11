# Bambino
A fileserver built with [gomek](https://github.com/joegasewicz/gomek)


### Options
This is set via the query params, e.g
```bash
?options={
  "files": ["cat1", "cat2", "cat3"],
  "data": {}, 
  "id": 1, 
  "entity_name: "User"
}
```
- files
- entity_name
- id
- data

### Docker Image
Pull down locksmith from the Docker registry
```
docker pull bandnoticeboard/bambino:v0.0.2
```

Or select a version here - [bandnoticeboard/bambino](https://hub.docker.com/r/bandnoticeboard/bambino)

### Response POST `/files`
This is a response examples if we were to use this request url:
`http://127.0.0.1:4444/files?options={"files": ["cat1", "cat2", "cat3"], "data": {}, "id": 1, "entity_name": "User" }`
```
[
    {
        "id": 1,
        "file_name": "cat1.png",
        "name": "cat1",
        "data": {},
        "entity_name": "User",
        "url": "http://localhost:4444/files/7/cat1",
        "path": "/files/7/cat1",
        "created_on": "2023-04-29 19:26:52.301456 +0100 BST"
    }
    ... etc.
]
```
Each file is stored on the server based on the `entity_name` value set in the `options` query param.
For example, with `"User"` `entity_name` the file would be stored at `files/User/1/image.jpg`

### Response GET `/files?id=1`
Get a file by ID. Response:

``` 
{
    "URL": "http://localhost:4444/uploads/cats/1/cat1.png"
}
```

#### Files
`http://localhost:4444/files`

#### Health
`http://localhost:4444/health`

### Configuration // TODO
`bambinoconf.yaml`