# Introduction

A simple way of the implementation of enumeration in Golang

## Getting Started
```
go get -v -u github.com/wardonne/goenum
```
## Usage
```golang
// enum based on basic type
type ImageType string

// implement Stringify interface
func (i ImageType) String() string {
    return string(i)
}

const (
    ImageTypePNG = "png"
    ImageTypeJPEG = "jpeg"
)

var ImageTypes = goenum.Enum[ImageType]{
    ImageTypePNG,
    ImageTypeJPEG,
}

func main() {
    // get all image types
    all := ImageTypes.Values()

    // get index of ImageTypePNG
    index := ImageTypes.Index(ImageTypePNG)

    // check if ImageTypePNG in ImageTypes
    exists := ImageTypes.Exists(ImageTypePNG)

    // get enum by extension
    if imageType, err := ImageTypes.FromString("png"); err != nil {
        // error handle
    } else {
        // do sth
    }
}

```

## Extend Usage
```golang
// struct enum
// use private field so that enum value can't be updated
type ImageType struct {
    extension string 
    mimeType string
}

// implement Stringify interface
func (i ImageType) String() string {
    return fmt.Sprintf("extension: %s, mimetype: %s", i.extension, i.mimeType)
}

// make extension readable
func (i ImageType) Extension() string {
    return i.extension
}

// make mimeType readable
func (i ImageType) MimeType() string {
    return i.mimeType
}

// extends goenum.Enum
type ImageTypeEnum Enum[ImageType]

// get enum by extension
func (i ImageTypeEnum) FromExtension(extension string) (ImageType, error) {
    var empty ImageType
    for _, enum := range this {
        if enum.extension == extension {
            return enum, nil
        }
    }
    return empty, goenum.ErrInvalidEnumValue
}

var (
    ImageTypePNG = ImageType{"png", "image/png"}
    ImageTypeJPEG = ImageType{"jpeg", "image/jpeg"}
)

var ImageTypes = ImageTypeEnum{
    ImageTypePNG,
    ImageTypeJPEG,
}

func main() {
    // get all image types
    all := ImageTypes.Values()

    // get index of ImageTypePNG
    index := ImageTypes.Index(ImageTypePNG)

    // check if ImageTypePNG in ImageTypes
    exists := ImageTypes.Exists(ImageTypePNG)

    // get enum by extension
    if imageType, err := ImageTypes.FromExtension("png"); err != nil {
        // error handle
    } else {
        // do sth
    }
}

```