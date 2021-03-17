# RELE.AI - Coding Challenge : David Marom
Converts an ungrouped YAML format to a grouped format and vice versa

## Backend

Requirements: Go lang has to be installed on the machine: https://golang.org/doc/install \
Run: `go run main.go` \
Port: 8081 

Routes:
  - ```/group```
  - ```/ungroup```

### Test routes using postman
Use the postman export attached: `rele.postman_collection.json`

Or create your own:
Send a GET request to http://localhost:8081/group (or ungroup) \
Put the YAML you wich to group / ungroup in the request body

### Config file for the group route

The server will use the `group-config.yaml` file as a dictionary for the grouping process \
Use this config file to define your own grouping rules \
Template: \
key:the group you want to put it in

Example: I want the key-pair `aaa:???` to appear under "type: ccc", I will add `aaa:ccc` to the `group-config.yaml` file

### Examples for grouped an ungrouped YAML files:
Grouped:
```yaml
type: DataFileWrite
payload:
    content: "foo"
    destination: "/tmp"
---
type: DataFileTimeout
payload:
    timeout: 60
```

Ungrouped:
```yaml
type: DataFile
payload:
    content: "foo"
    destination: "/tmp"
    timeout: 60
```
