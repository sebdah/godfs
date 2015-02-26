# godfs

## data-node

The data node implements the following URIs:

- `GET /files/{id}` - Returns the contents of a file with id `{id}`
- `POST /files/{id}` - Create a file with id `{id}`. The contents of the `POST` is the file data
- `DELETE /files/{id}` - Delete file with id `{id}`
