package bytetinkerer

import (
	"io"
	"io/ioutil"
)

type FilteredBytes struct {
	Bytes []byte
}

type BytesHandlerFunc func(*FilteredBytes) error

//From will create new Filtered List out of any io.Reader and will wrap handlers around it
func From(stream io.Reader, handlers ...BytesHandlerFunc) (*FilteredBytes, error) {
	data, err := ioutil.ReadAll(stream)
	if err != nil {
		return nil, err
	}
	fb := FilteredBytes{
		Bytes: data,
	}
	for _, h := range handlers {
		err := h(&fb)
		if err != nil {
			return nil, err
		}
	}
	return &fb, nil
}

//Stamp will generate specific string as a stamp to data
func Stamp(withValue string) BytesHandlerFunc {
	return func(fb *FilteredBytes) error {
		fb.Bytes = append(fb.Bytes, []byte("\n"+withValue)...)
		return nil
	}
}

//Extract will update our newely created lists without noise
func Extract(target []byte) BytesHandlerFunc {
	return func(fb *FilteredBytes) error {
		newList := make([]byte, len(target))
		var j int = 0
		for i := 0; i < len(fb.Bytes); i++ {
			if target[j] == fb.Bytes[i] {
				newList = append(newList, target[j])
				j++
			} else {
				newList = newList[:0]
				j = 0
			}
			if j == len(target) {
				fb.Bytes = newList
				break
			}
		}
		return nil
	}
}

//Remove fuction will delete []byte that is matching our target from our newely created list
func Remove(target []byte) BytesHandlerFunc {
	return func(fb *FilteredBytes) error {
		newList := make([]byte, len(fb.Bytes))
		var j int = 0
		for i := 0; i < len(fb.Bytes); i++ {
			if target[j] != fb.Bytes[i] {
				newList = append(newList, fb.Bytes[i])
				j = 0
			} else {
				if len(newList) > 0 {
					newList = newList[:len(newList)-1]
				}
				j++
			}
			if j == len(target) {
				newList = append(newList, fb.Bytes[i:]...)
				fb.Bytes = newList
				break
			}
		}
		return nil
	}
}

//ConvertToString is our helper function that will convert the data to a readable string
func (fb *FilteredBytes) ConvertToString() string {
	return string(fb.Bytes)
}
