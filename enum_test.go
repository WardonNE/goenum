package goenum

import "testing"

// enum based on basic type
type imageType string

// implement Stringify interface
func (i imageType) String() string {
	return string(i)
}

const (
	imageTypePNG  = "png"
	imageTypeJPEG = "jpeg"
)

var imageTypes = Enum[imageType]{
	imageTypePNG,
	imageTypeJPEG,
}

func TestEnum(t *testing.T) {
	all := imageTypes.Values()
	for index, item := range all {
		if item != imageTypes[index] {
			t.Errorf("element %d excepted to be %v, but %v got", index, imageTypes[index], item)
		}
	}

	index := imageTypes.Index(imageTypePNG)
	if index != 0 {
		t.Errorf("imageTypes.Index(imageTypePNG) excepted to be 0, but %d got", index)
	}

	index = imageTypes.Index(imageType("gif"))
	if index != -1 {
		t.Errorf("imageTypes.Index(imageType(\"gif\")) excepted to be -1, but %d got", index)
	}

	exists := imageTypes.Exists(imageTypePNG)
	if !exists {
		t.Errorf("imageTypes.Exists(imageTypePNG) excepted to be true, but %v got", exists)
	}

	exists = imageTypes.Exists(imageType("gif"))
	if exists {
		t.Errorf("imageTypes.Exists(imageType(\"gif\")) excepted to be false, but %v got", exists)
	}

	item, err := imageTypes.FromString("png")
	if err != nil {
		t.Errorf("imageTypes.FromString(\"png\") excepted to be returned nil error, but %v got", err)
	} else {
		if item != imageTypePNG {
			t.Errorf("imageTypes.FromString(\"png\") excepted to be %v, but %v got", imageTypePNG, item)
		}
	}

	item, err = imageTypes.FromString("gif")
	if err == nil {
		t.Errorf("imageTypes.FromString(\"gif\") excepted to be returned error %v, but nil got", err)
	} else {
		var empty imageType
		if item != empty {
			t.Errorf("imageTypes.FromString(\"gif\") excepted to be empty, but %v got", item)
		}
	}
}
