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


#### Files
`http://localhost:4444/files`

#### Health
`http://localhost:4444/health`