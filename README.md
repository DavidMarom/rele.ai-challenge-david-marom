# RELE.AI - Coding Challenge : David Marom
Converts an ungrouped YAML format to a grouped format and vice versa

## Backend
Requirements: Go lang has to be installed on local machine: https://golang.org/doc/install \
Run: `go run main.go` \
Port: 8081 

Routes:
  - ```/group```
  - ```/ungroup```

### Test routes using postman
Use the postman export attached: `rele.postman_collection.json`

Or create your own:
Send a GET request to http://localhost:8081/group (or ungroup) \
Put the YAML format in the request body

### Config file for the group route

The server will use the `group-config.yaml` file as a dictionary for the grouping process \
Template: \
<key>:<the group you want to put it in>

## Frontend
