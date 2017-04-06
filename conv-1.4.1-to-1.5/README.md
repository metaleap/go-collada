# convert
--
    import "github.com/metaleap/go-collada/conv-1.4.1-to-1.5"

Provides in-memory conversion of Collada 1.4.1 documents to version 1.5.

## Usage

```go
var (
	//	If true, conversion-logic is always run against the given input document's entire node tree;
	//	if false, conversion is only performed if the input "COLLADA" root element's "version" attribute is not "1.5" or higher.
	Force = false

	//	Set this to an image file format (BMP, JPG, PNG etc.) for Collada 1.4.1 documents
	//	that have binary images hex-encoded in <image><data> elements *without* a format attribute.
	HexFormat = ""

	//	The default logging function used by this package. Simply does a log.Printf(format, fmtArgs...)
	//	Set this to nil to disable logging.
	Log = func(format string, fmtArgs ...interface{}) {
		log.Printf(format, fmtArgs...)
	}

	//	The <shader>/<compiler> element introduced in Collada 1.5 requires a "platform" attribute, but there was no
	//	such equivalent in Collada 1.4.1. In the context of importing such 1.4.1 documents, probably any arbitrary
	//	value will do for this new attribute, but this field still gives you the choice to set this to a preferred
	//	value for all your <shader>s.
	ShaderCompilerPlatform = "PC"

	//	Conversion mode: true for strict mode, or false for lax mode.
	//	If strict, obsoleted elements and attributes are removed or rewritten so that the remaining document is (in theory) strictly 1.5-conformant and will (ideally) validate against the Collada 1.5 XML schema definition.
	//	If lax, obsoleted elements and attributes are not removed, for use-cases where your 1.5 loader consuming the conversion result is known to simply ignore or discard them quietly.
	//	Note in practice there won't be any noticeable difference in performance or output for approximately 95% of "common use-case" Collada documents...
	Strict = true
)
```

#### func  ConvertBytes

```go
func ConvertBytes(srcFile []byte) (dstFile []byte, err error)
```
Converts the specified Collada 1.4.1 document to Collada 1.5.

#### func  ConvertDoc

```go
func ConvertDoc(srcFile []byte) (doc *xmlx.Document, err error)
```
Converts the specified Collada 1.4.1 document to Collada 1.5.

--
**godocdown** http://github.com/robertkrimen/godocdown
