# Delete Entities from Google Datastore
Deleting entities from Google Datastore with simple gin gonic server in go.
## How to use
go get -u github.com/rakeshkumargupt/entity_delete_datastore

# Example

Sample Payload: 
  {
	"namespace" : "NS1",
	"kind" : "Knd1",
	"names" : ["abc","cde", "fgh"]
 }

# Dependencies
1. Install Go
2. Google Datastore

# References
For more information visit these Google Developers links:

https://cloud.google.com/appengine/docs/standard/go/datastore/creating-entities

https://cloud.google.com/appengine/docs/standard/go/datastore/reference#Delete
 



