# kluff-go-sdk

## Generate GRPC code
- add yours proto [file] and [folder] in apps-core/api/proto/[folder]/[file]
- append your folder name of proto in proto_folder.txt that is in root
- come to root and run this command in terminal type
```console
chmod +x generate_protos.sh
```
- on root run below command this will create pb and grpc file in pkg/api/[folder]/
```console
./generate_protos.sh
```
- run this command in terminal before any step if path issue
```console
export PATH="$PATH:$(go env GOPATH)/bin"
```