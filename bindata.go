// Code generated by go-bindata.
// sources:
// assets/glauth.css
// assets/glauth.js
// assets/index.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsGlauthCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x41\x6e\xc3\x20\x10\x45\xf7\x9c\x62\x24\xaf\x63\xd9\x0e\x4e\xc1\x39\x0d\x03\x63\x9b\xd6\x01\x04\xa4\xad\x5a\xf5\xee\x95\x1d\x39\x2a\xc9\xa6\xeb\xff\xfe\x7f\x33\xac\x7e\xbd\x5e\xd0\xe7\xe8\x1d\x7c\x33\x00\xed\x17\x1f\x07\xa8\x74\x73\x94\x1d\x9e\x19\x00\x2a\xfd\x36\x45\x7f\x75\xe6\xb0\x87\xa4\xc7\x66\x6c\xcf\xec\x87\xb1\x2a\x65\x95\x53\x50\x8e\x96\xa2\x7f\xe4\x5c\xf6\xf4\x84\xd4\x17\xca\xd1\xea\x07\x54\x0a\x83\x7f\xd0\x22\x95\xd8\x4b\x3c\xad\x87\x04\x65\x8c\x75\xd3\x00\x22\x7c\xae\x74\x95\x28\xbe\x53\xfc\x2f\xce\xe6\x6e\x25\x77\xb0\x93\xa2\x41\x79\xb3\x66\x65\x17\xc8\x73\x31\xf4\x32\x0a\x2d\xcc\xe6\xb9\xc5\x66\x8b\xef\xab\x7c\x5d\x7d\xfe\x17\x60\xf4\x2e\x1f\x92\xfd\xa2\x01\xda\x5d\xbd\x4f\xd4\x3e\x14\x12\x41\x9c\x2b\x73\x6f\x7d\x90\x9d\xe6\x3c\x00\xfa\xa5\x30\xd7\x3e\xd4\x68\x9d\x29\xba\xed\x49\x35\xa2\x7f\xc4\x12\xa9\xa8\xcb\x4f\x5a\x85\x5a\xea\xed\x8e\xdf\x00\x00\x00\xff\xff\x91\x33\xcc\x6e\xee\x01\x00\x00")

func assetsGlauthCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsGlauthCss,
		"assets/glauth.css",
	)
}

func assetsGlauthCss() (*asset, error) {
	bytes, err := assetsGlauthCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/glauth.css", size: 494, mode: os.FileMode(436), modTime: time.Unix(1524616076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsGlauthJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\xd1\x8a\xdb\x30\x10\x7c\xf7\x57\x6c\xdd\xc0\xc9\x34\x95\x43\x5f\x43\xfa\x76\x70\x29\xf4\x0a\x6d\x7e\x40\xb1\xd6\x89\xa8\x2c\x99\xf5\xc6\x69\x39\xfc\xef\x45\x96\x9d\xf8\x72\x77\x70\xc5\x2f\x49\xd8\x9d\x99\x8c\x35\x63\xb5\x8a\xc0\x38\x46\x6a\x95\x85\x0d\x7c\x59\xad\x56\xeb\x24\x59\x88\xf2\xe4\x0a\x36\xde\x89\x0c\x9e\x12\x80\x53\xad\x15\xe3\x77\x64\x32\x45\x23\xb2\x75\x02\x90\xe7\x71\xb8\x53\xc6\x86\x49\x97\xad\x93\x24\xc8\xb1\x32\x76\xab\x61\x03\x41\x69\xd4\xb9\x55\xe8\x55\x03\xba\x8a\x13\xd8\xc0\x42\x1e\x90\xbf\xfd\xfa\xf1\x28\x20\xcd\x35\xee\x4f\x87\xbc\x55\xd4\xa4\x4b\xb8\x98\xd1\x91\x17\x99\x8c\x55\x6d\x15\x23\x6c\xe0\x41\x39\x6d\x71\xaf\xa8\x91\x85\xaf\x6a\x63\x51\xc0\x42\xa4\x1f\x1b\x56\xdc\xd4\xca\xa1\xfd\x3c\xa2\xd3\x4c\x1e\xb9\xb2\x22\x83\xfe\x29\xa2\x54\x98\x84\xdf\x9b\x8b\xa8\xd0\xc3\xfa\xb9\xcc\xc8\x0e\x1f\x13\xfe\x7b\xad\xcc\x77\x31\xc3\x00\x52\x8b\x34\xd3\x42\xd4\x78\x61\xa2\xcb\x12\x00\x59\x86\x26\xdc\x14\xe7\xad\x03\x4c\xef\x89\x3c\x7d\x48\x27\x74\x65\xcf\xea\x6f\xf3\x52\xa0\x41\xde\x99\x0a\xfd\x89\xc5\xb3\x12\x2d\xc7\xde\x0e\x1a\xeb\xa4\xbb\xad\x5b\xec\xe6\xa5\x6b\xa1\x99\xb7\x45\x0b\xb3\x3c\x85\x4f\x63\x6d\x5f\x6b\x9b\x96\xa5\xa7\x7b\x55\x1c\xaf\xe6\x70\xdc\x01\x98\x12\x04\xca\xad\x86\xaf\x83\xc6\x75\x05\xd7\x97\x21\x20\xd6\xc3\xb8\xeb\xbf\xbb\xff\xcb\x2f\x28\xcd\x08\xaf\xcf\x4a\x1e\x8d\x46\x91\xc9\x9a\xb0\x46\xa7\x77\x7e\xd0\x4d\x33\x59\x2a\x8d\x5b\x27\xc2\x05\xf0\xae\x48\x07\xda\xec\x30\x43\x44\xaf\x25\x99\x4c\xce\x81\xf0\x60\x1a\x46\x7a\x40\x5b\x23\x89\xbb\xc0\xfb\x89\x56\xb1\x69\xf1\x6e\x12\x58\xe1\x1d\xe3\x1f\x5e\xc2\xde\xfa\xe2\x77\xfc\xcb\x90\xce\xd9\x38\xed\xcf\xb2\xf2\x15\x3a\x1e\x9d\xe4\x39\x21\x9f\xc8\x41\x1c\x8f\xe4\x4c\x96\xe4\xab\x47\x7f\x16\xc3\xd1\xbd\x85\xf2\x54\x29\x8e\xa0\x0e\x6d\x83\x4f\x53\xf4\x00\xeb\x97\xc3\xd5\xf8\x2f\x00\x00\xff\xff\x1d\x12\x43\xad\x6a\x05\x00\x00")

func assetsGlauthJsBytes() ([]byte, error) {
	return bindataRead(
		_assetsGlauthJs,
		"assets/glauth.js",
	)
}

func assetsGlauthJs() (*asset, error) {
	bytes, err := assetsGlauthJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/glauth.js", size: 1386, mode: os.FileMode(436), modTime: time.Unix(1526347681, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x6d\x6f\xdb\x36\x10\xfe\x9e\x5f\x71\xd1\xbe\x4e\x62\xed\xae\x5b\x17\xc8\x06\xba\x34\x5b\x03\xb4\x6b\x90\x76\xd8\x86\xa2\x28\x68\xf1\x6c\xd1\xa1\x48\x85\x77\x72\x62\x08\xfa\xef\x03\x2d\xbf\xc9\x73\x12\xc7\xd8\xbe\xd8\x22\xf9\xdc\xcb\xf3\xf0\x8e\xa2\xd2\xd3\xb7\x1f\xcf\x3f\xff\x7d\x75\x01\x39\x17\x66\x78\x92\x86\x3f\x30\xd2\x4e\x06\x11\xda\x68\x78\x02\x90\xe6\x28\x55\x78\x00\x48\x0b\x64\x09\x59\x2e\x3d\x21\x0f\xa2\x8a\xc7\xf1\xeb\x68\x7b\x29\x67\x2e\x63\xbc\xad\xf4\x6c\x10\xfd\x15\xff\xf1\x26\x3e\x77\x45\x29\x59\x8f\x0c\x46\x90\x39\xcb\x68\x79\x10\x5d\x5e\x0c\x50\x4d\xb0\x63\x69\x65\x81\x83\x68\xa6\xf1\xae\x74\x9e\xb7\xc0\x77\x5a\x71\x3e\x50\x38\xd3\x19\xc6\x8b\xc1\xf7\xa0\xad\x66\x2d\x4d\x4c\x99\x34\x38\xe8\xed\x71\xa4\x90\x32\xaf\x4b\xd6\xce\x6e\xf9\xda\x03\x94\x15\xe7\xce\x77\x31\x2d\x88\x35\x1b\x1c\xfe\xf6\xfe\x4d\xc5\x39\x7c\x62\xc9\x15\xa5\xa2\x9d\x5c\x22\x8c\xb6\x37\xe0\xd1\x0c\x22\xe2\xb9\x41\xca\x11\x39\x82\xdc\xe3\x78\x10\x05\x25\xe8\x4c\x88\x42\xde\x67\xca\x26\x23\xe7\x98\xd8\xcb\x32\x0c\x32\x57\x88\xf5\x84\x78\x99\xf4\x93\x17\x22\x23\xda\xcc\x25\x85\xb6\x49\x46\xb4\xca\xf7\x7f\x0b\x14\x73\x8e\x05\x3e\x27\x9c\x24\x42\x26\x31\x31\x41\xb8\xa5\x51\x6b\x75\x1a\xc7\xf0\xee\xf3\x87\xf7\xaf\x80\x72\x5d\x80\xb4\x0a\xae\x91\x4a\x67\x55\x32\x25\xb8\xbc\x78\x0d\x54\x95\x61\x73\xc1\x8d\x97\x40\x34\x58\xa0\x65\x5a\x80\x0b\x54\x5a\xc2\x6d\x85\x5e\x23\x41\x1c\x0f\xd7\x6e\xbf\xe8\x31\x18\x86\xcb\x0b\xf8\xf9\x6b\x3b\x0b\x90\xb6\xfb\x0b\xe4\xb3\x8d\x06\x8e\x28\x59\xea\x10\xa8\x87\x5a\x7e\x45\xb9\x9e\x89\x97\xc9\x4f\x49\x7f\x33\x5e\x10\x9e\x52\x34\x4c\x45\xeb\xe6\x39\x5e\x7d\x4b\x4a\xf4\x92\x1f\x92\xfe\x6a\xf4\x80\xc7\xf4\xf4\x0b\x5a\xa5\xc7\x5f\x5b\x3a\xa9\x68\x7b\x29\x3c\x8e\x9c\x9a\x83\x77\x26\x14\xab\xcb\xaa\xa0\xc3\x5a\x4a\xa5\x67\x90\x19\x49\x34\x88\x42\x5d\x4a\x6d\xd1\xc3\x62\xab\x62\xca\xdd\x5d\x26\x09\xa3\xa5\x6d\x21\xb5\x5d\xd9\x2d\x37\xe1\x83\xd4\x16\xa6\x55\x31\x72\xec\x9d\x85\xb1\xf3\x20\xa1\xf4\xba\x90\x7e\x0e\x85\xf4\x37\xc8\xda\x4e\xa0\x40\x22\x39\x41\x70\x1e\x32\x69\x0c\xb0\x03\x99\x85\x86\x59\x6b\xdf\xcd\x64\xed\x31\x5a\xad\xae\x11\x5a\x85\x52\x91\x4c\xa5\xb4\x68\xa2\x95\xc5\x62\x04\x8b\xdf\x58\xe1\x58\x56\x86\xa1\xac\x8c\x89\xbd\x9e\xe4\xdc\x75\x23\x94\x9e\x6d\x26\xd2\xbc\xb7\xec\xbc\x54\xe4\xbd\xad\xf9\x72\xf8\x39\xd7\x04\xec\x9c\x01\x4d\xa0\x2d\x28\x9c\xa1\x71\x65\xd0\x0f\x62\xb0\x8e\xe1\xb6\xd2\x8c\xe0\x51\xaa\xf9\x82\x7c\xe9\x9d\xaa\x16\xcc\x4e\x53\x51\x6e\xb4\x6a\x43\xee\x61\xea\xdd\xdd\x56\x72\xdd\xdd\x30\xf1\x3d\xc5\xbd\x3e\x84\x27\x2a\xe2\x1f\xbb\x2c\xf2\xfe\x70\xa1\x43\x2a\xf2\xfe\xc3\x2a\x45\x4f\x48\xba\xa3\x45\x27\xcf\x67\xe7\x33\x92\xd9\x0d\x5a\x05\x84\x7e\x86\xfe\x91\xcc\xda\xf5\xa3\x72\xfb\x77\xa2\xed\x73\x5b\x8f\x62\x53\xc3\xa1\xb2\x4e\xb6\x1b\xad\x5b\x39\x31\x63\x51\x1a\xc9\x18\x01\xcf\x4b\x1c\x44\x8c\xf7\x2c\xee\xe3\x5c\x5a\x65\x70\x24\x3d\x6d\x10\xfb\x37\xa8\x75\x13\x7a\xab\x93\x76\x08\x58\x4a\xbb\x42\x15\xc8\x5e\x67\xd1\xb0\xf4\xee\x7e\x9e\x04\xda\xda\xd9\xb3\x54\x04\xcc\x10\xea\xba\x33\xdf\x34\xe9\xc8\xef\xa5\xbc\xe9\xf5\xfd\x9c\x8e\xa1\xf3\x70\x9e\xdf\xc6\x7e\xf1\x8e\x52\xc9\x48\x5b\xf5\xcd\xe3\x2d\xed\xa6\xbc\x07\xd2\xc9\xfe\x60\xe7\x54\x65\x19\x12\xe1\x53\x11\xd6\xb8\xe3\xc2\xa0\xf7\xce\x3f\x15\xa3\x05\x3d\x3b\x00\xa1\xf4\x59\xfe\x84\x4c\x5b\xa0\x63\x03\x1c\x20\xd5\x2e\xf2\xd8\x50\x4f\xc9\xd5\x81\x3d\x3b\x48\x66\xdc\xa3\x2c\xda\xf5\x2d\xb7\x8f\x95\x7f\x7b\x98\x1c\xd3\x00\x75\xfd\x1d\xca\x2c\x87\x36\xfa\xf2\xf4\x4a\x96\x0e\x9b\xe6\xe4\xd1\x86\xae\xeb\x4f\x59\x78\x4d\x36\xcd\x99\x10\x75\xfd\xce\x11\x87\x6b\x5e\xd3\x9c\xd5\xf5\x95\xf3\xdc\x34\x1b\x7a\x57\xda\x4e\x9a\x06\xea\xba\xbd\xdb\x35\x0d\x6c\x09\x56\xd7\x22\x64\xb1\x8c\xf7\x08\x53\x96\xfa\xa8\x63\x2b\xe5\x70\x48\x6d\x8e\x95\x15\x6d\xce\x75\x87\x25\xfb\x15\xc7\xdd\x03\x8d\xd5\x6a\xc5\x95\x50\xd7\x1f\xcb\xa6\x09\x02\x84\xff\x54\xb0\x3a\x0c\xfd\x8b\xb6\xea\xed\xef\xcf\xb1\xf8\xb4\x28\xb2\x6b\xbc\x4d\x7e\xd5\x86\xd1\x3f\x61\x1b\x4c\xae\x91\x2a\xc3\xe7\x4e\xe1\x01\xe0\x95\xff\x60\x42\x07\xe0\xcf\x8d\x46\xcb\x97\x57\x07\x40\x95\x64\xbc\x46\x23\x59\xcf\x10\xfe\xcc\xd1\xee\xda\xa4\x82\xfd\xf6\x96\x6c\xd7\x40\x77\xcb\x1e\xa8\x88\xce\xbd\x51\x4e\xe5\x7d\x32\x71\x6e\x62\x50\x96\x9a\x16\x97\xc7\x30\x27\x8c\x1e\x91\x98\x86\x9b\xee\x5c\xf4\x92\x5e\x2f\xe9\x2d\x47\x0f\xdd\x22\xf7\x79\x3f\xf4\xbe\x3f\xdd\xfd\xae\x38\xd0\x7f\xa6\xec\x94\x42\xd3\x57\x6a\x6c\xa4\xc7\x9d\xf4\x37\xb5\x9d\x4c\x49\xf4\x93\x17\xc9\x8b\xed\xb9\xff\x2e\x50\xe1\xc2\xd5\xae\x0d\xf2\x3a\x79\xb9\x1c\xc7\x77\x9a\xf3\xd8\xb8\xf0\x01\x78\x50\xb4\xee\x77\xcb\x2e\x38\x15\xed\xd6\xa6\xa2\xfd\x0c\xfe\x27\x00\x00\xff\xff\x08\x9c\x44\x70\x17\x0f\x00\x00")

func assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexHtml,
		"assets/index.html",
	)
}

func assetsIndexHtml() (*asset, error) {
	bytes, err := assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.html", size: 3863, mode: os.FileMode(436), modTime: time.Unix(1524616076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"assets/glauth.css": assetsGlauthCss,
	"assets/glauth.js":  assetsGlauthJs,
	"assets/index.html": assetsIndexHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"assets": {nil, map[string]*bintree{
		"glauth.css": {assetsGlauthCss, map[string]*bintree{}},
		"glauth.js":  {assetsGlauthJs, map[string]*bintree{}},
		"index.html": {assetsIndexHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
