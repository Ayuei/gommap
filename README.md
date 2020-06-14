# `gommap`: Golang tools for working with MetaMap

`gommap` is a set of tools for interacting with the National Library of Medicine's [MetaMap](http://metamap.nlm.nih.gov) concept identification and extraction tool. 

## TODO

- [x] `gommap` relies on a couple of hard-coded paths; that obviously needs to change
- [] More configurability
- [] More documentation
- [x] Currently, the `MMOs` struct in `mmo_struct.go` doesn't capture all of the data in MetaMap's output. 
- [x] Refactoring- `mm_server` should only handle network-related stuff, the actual MetaMap I/O should be handled elsewhere.
