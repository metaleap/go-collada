# collimp
--
    import "github.com/go3d/go-collada/imp-1.5"

Loads assets from Collada 1.4.1 and 1.5 XML documents into the data structures
provided by the go-collada/dom package. Note, Collada 1.4.1 documents are
insta-bootstrapped to version 1.5 in-memory using the
go-collada/conv-1.4.1-to-1.5 package.

## Usage

#### func  ImportCollada

```go
func ImportCollada(colladaDoc []byte, importBag *ImportBag) (doc *cdom.Document, err error)
```
Imports the specified Collada document, using the import options specified in
importBag.

#### type ImportBag

```go
type ImportBag struct {
}
```

Provides options for importing Collada documents.

#### func  NewImportBag

```go
func NewImportBag() (me *ImportBag)
```
Initializes and returns a newly created ImportBag instance.

--
**godocdown** http://github.com/robertkrimen/godocdown
