// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package lib

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 12, 3, 11, 2, 29, 369645557, time.UTC),
		},
		"/std.js": &vfsgen۰CompressedFileInfo{
			name:             "std.js",
			modTime:          time.Date(2018, 12, 3, 14, 41, 58, 215162776, time.UTC),
			uncompressedSize: 27298,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7c\x5d\x73\xa3\x38\xd6\xf0\xfd\xfb\x2b\x48\x2e\x52\x30\x26\x94\x21\x6e\x4f\xc6\x58\xee\xea\x9e\xe9\xec\x66\xab\x37\x3d\xd5\xc9\xcc\xd6\x6e\x2a\xe5\xc2\x46\x24\x7a\x87\x80\x47\x88\xce\x66\x1d\x3f\xbf\xfd\x29\x49\x48\x08\x21\xf0\x47\x67\x7b\xe6\xe2\xb9\xe9\x8e\x41\xe7\x43\x47\xe7\x1c\x9d\x0f\x89\x2f\x11\xb6\x08\x20\x2f\x2f\xeb\x4d\x48\xbc\x8b\x1c\x3f\x46\x04\xac\xdf\x95\x24\x9f\x0c\xdd\xbf\x5d\x7f\xba\x9a\xf8\xee\x3f\xdf\xfd\xfd\xe3\x24\x70\x3f\x47\x4f\x93\x33\x3a\xea\x1f\x18\x11\xf8\x0e\xdf\x17\x20\x29\xb3\x25\x41\x79\x66\x3b\x6b\xf2\x80\x0a\x6f\xb1\x00\x59\x99\xa6\x61\xf5\x63\xbe\xca\x0b\x30\x6c\x80\x78\x2b\x9c\x93\x9c\x3c\xaf\xa0\x37\x9f\xa3\x0c\x91\x1a\x07\x71\xa1\x44\xc3\x20\x89\xc0\x03\x60\x88\x21\x29\x71\x66\xd1\x07\x4d\x7c\xf7\x90\x7c\xce\x73\xf2\xae\x30\x70\x05\x5d\xec\xac\x39\xa4\x8d\x5f\x5e\x32\xf8\x64\x29\xa0\x4e\xc5\x81\x0d\x3d\x0c\xa3\xf8\x32\x23\x67\x81\x0d\xbd\x55\x5e\x20\x3e\x27\x67\xa0\xfe\x72\xa1\xd3\x35\x93\x55\x44\x1e\x94\x79\x38\x6b\x2a\x55\x08\x2a\xee\xbd\xf9\x3c\x4f\x92\x02\x12\x5b\x99\x9c\x3b\x72\xc4\x9c\xe0\xdb\x7a\x60\x41\x30\xca\xee\xd5\x81\x03\xe8\x12\x67\x42\xa5\xda\x45\xfd\x4b\x94\x96\x70\x5f\xf2\xe3\x57\x23\x4f\xff\x51\x15\x61\x07\xe2\xe7\x06\xe2\xd5\x12\x9c\x37\x89\x3b\x13\xa1\x94\x1e\xd5\xc9\x2e\x1e\x50\x16\xc3\x8c\xe8\x5c\x90\x7e\x2e\xfc\xa1\x64\x83\xf4\xb3\x41\x9c\x89\xa6\xc5\x05\x89\x30\x31\x68\x1c\x71\xd6\x84\xbf\xfc\xb4\xf8\xff\x70\x49\xec\x91\xa6\x34\x51\x1c\xff\xdc\xd4\x16\xa6\xf5\xf4\xf9\x05\x82\x69\xfc\x89\xf3\x39\x74\xa1\x3b\x6c\x83\xfe\xaa\x2d\xb5\x11\xd6\x37\xc3\xde\x34\xd6\x89\x99\x06\x94\xa0\x6c\xca\x81\x8b\xdd\x86\xb8\xdb\x48\x2e\x35\x41\x6b\x1c\x30\x34\x67\x06\xfa\x30\x8b\xcd\xd2\xaa\xb4\x85\x0e\xa8\x24\x56\xeb\xc6\x26\xe4\x6f\x21\xf3\x4e\xd0\x63\xc0\xeb\xab\x4f\x57\x1f\x26\x43\x57\xa2\x9b\xf8\xee\x67\x18\xc5\xec\xcf\xc0\xfd\x31\xca\x96\x30\x65\x3f\xce\xdc\x0b\x94\xc2\xcb\x2c\xc9\xd9\xcf\x91\xfb\x11\x15\x84\xfd\xf9\x86\x62\xfb\x3b\x2c\x8a\xe8\x1e\xee\xea\xc3\x24\xc0\xab\x78\xb0\x1a\x9b\xf4\x5f\x2d\x7e\x48\xdb\x7b\x49\x30\xe9\xbb\x88\xe2\xbb\x48\xc3\x77\xa9\xbf\x5c\xe2\x98\x67\x10\xe1\xfb\xe2\xc6\x60\xc0\x5b\x4c\x67\xd4\x61\x39\xbf\x20\x83\xe9\xf0\x85\xf3\xe8\xba\x75\x33\xf1\x3a\xfe\xab\xcc\xb8\xe0\x34\x1f\xc2\xdd\x57\x4d\x9a\x19\x68\x5b\xde\xba\xed\x06\x0d\xa1\x45\x31\xd3\xb1\xa6\xb4\xd8\x1a\x69\xfa\x3f\x74\xb1\xab\xcc\xd9\x84\x64\x57\x13\xae\x01\x61\x16\x9b\x18\xde\x6e\x3e\x18\xac\x37\x21\xf6\xae\x2f\xff\xf5\xe1\xd3\xc5\xfc\xfa\xaf\x9f\x3e\xdf\x80\xa0\x7e\x70\x79\x75\x03\x46\x21\xf6\x2e\x2e\x3f\x7e\x98\x5f\xfe\xf4\xe1\xea\xe6\xf2\xe2\xf2\xc3\xe7\xf9\xc7\x0f\x57\x7f\xb9\xf9\x2b\x7b\xf5\x21\x5b\xe6\x31\xca\xee\xc1\xfa\x97\x9b\x8b\xf3\xf9\xfb\x7f\xde\x7c\xb8\x9e\xf8\xee\x2f\x37\x17\xfe\x78\x7e\x7d\xf3\xf9\xf2\xea\x2f\x93\x80\xd2\x40\x54\x0b\x01\xd5\x53\xa6\x8f\xef\x30\x8e\x9e\xed\xc0\x09\xb1\x97\xa4\x79\x24\xde\x5d\xf0\xbf\xf9\xdb\x0a\xc8\x5b\x94\x49\x02\xb1\x1c\x3a\x1e\xd5\x43\xc7\xa3\xae\xa1\xa8\xf8\x88\x08\x49\xe1\x87\x2c\x46\x51\xc6\x20\xa8\x02\xfa\x63\x0e\x20\x7e\x9f\xf3\x9f\xb7\xbe\x3b\xbc\x73\x04\xf8\xed\xf0\x0e\x00\xe0\x87\xd8\xfb\x98\x67\xf7\x26\x13\x4e\xf3\x27\x40\x5e\x86\xdc\x82\x1f\xd0\xfd\x03\x80\x2f\xc3\x4d\x05\xe0\x2d\x31\x8c\x88\xee\x8a\x85\x4d\x00\x30\x3c\x39\x81\x00\x0c\xdf\x56\xa3\xff\xf5\xe1\xf3\xa7\x09\xe5\x87\xff\x66\xa3\x25\x2a\x65\x2b\xcd\xab\x19\xab\xe6\x58\x79\x00\xc1\xd3\x6c\x36\x1b\x3a\x03\xc9\xd4\x77\xa3\xe0\x87\xd1\x0f\xe3\xef\x83\x1f\xc6\x06\x84\xf0\xf7\x32\x4a\x9b\xb6\xa5\xb8\x21\x36\x45\x40\xe8\x7f\x27\x27\xf5\x34\x01\x61\xff\x4b\x74\x94\x79\xa0\x30\x3f\x74\xe9\xc6\xe9\xbd\x2f\x51\x1a\x43\xdc\x40\x8e\x12\xfb\x48\x6a\xa5\x3f\x0c\x46\x1b\x98\x16\x50\x68\xe9\x46\x38\x43\xec\xbd\x7f\x26\xf0\x3d\x5b\x09\x2f\x4a\xd3\x7c\x19\x11\x68\x43\x87\xcb\xba\x58\x45\x4b\x08\x20\xff\xf1\x88\xb2\x28\x45\xf7\x19\xf0\xf9\xef\x2f\x24\x5a\xa4\x50\xf1\xd0\xfc\xc1\x1c\x65\xf3\xb2\x80\xa0\x5a\x2e\x54\x5c\xc1\x82\xc0\x18\x24\x51\x5a\x54\x98\x72\x66\x1f\x73\x66\xe0\x62\x1c\x07\x2e\xc0\xed\x5d\xf5\x1b\x2e\x49\x8e\xe7\x59\xf9\x38\x87\x29\x7c\x2c\xc4\xb8\x24\xc7\x4b\x38\x8f\x61\x12\x95\x29\x29\x38\xd6\x4d\x2d\x03\x45\xe2\x6c\xe4\x4f\x72\xa0\xea\x5b\x0c\x88\x88\x19\x49\x1c\x91\x88\x4b\xa7\xad\x08\x56\x25\x44\x33\x64\x54\xd4\x1a\xdf\x0d\xeb\x2d\x9e\x09\x2c\x6c\xc7\x2b\xca\x45\xc4\x8c\x43\xbc\x50\xb7\x8c\xd6\x23\xae\x75\x95\x33\x76\x9c\x2e\x0e\x18\xf1\x6d\xec\x4b\x16\xb8\x3d\x7a\x45\x8a\x96\xf0\x15\xd9\x58\x61\xb8\xd2\xac\x13\x25\x36\x99\x35\x94\xaa\x5a\x14\xa9\x63\x64\x43\x35\x15\x81\xff\x91\x8c\x2c\xa3\x55\xb4\x44\xe4\xd9\x76\x4e\x6b\xdd\x1c\x40\x67\xe0\x9f\x90\x53\x3f\x7c\x7a\x40\x69\xc5\x35\x7b\x33\x45\x03\x32\x80\xdc\x00\x0a\xd0\x46\x12\x2a\x16\x50\xb1\x7c\x8f\xf3\xa7\xda\x1a\x04\x61\xd5\x14\x06\x06\x44\xa7\x05\x37\xa6\x55\x14\xdb\xa8\x4b\x04\x51\xdc\x50\xc0\x24\xc7\x36\xb7\xc4\x61\x08\xa7\x24\x84\x83\x81\x8c\x56\xbc\x27\x1a\x4e\xb1\xed\xeb\x54\x99\x29\xdd\x84\xcc\xc8\xe5\xf8\xb6\x8e\x37\xb0\xd5\xb8\x4e\x81\xcf\xc2\x90\x3e\x6c\xfe\xb8\x17\x9d\x3f\x6e\xe0\x0b\xb6\xe2\x3b\x0b\x7a\xf1\xd1\x98\x49\xc1\x37\xda\x8a\x4f\xf5\xcc\x06\x7c\xe3\x51\x03\xdf\x79\x3f\xbe\x6a\x23\xec\xc6\x58\x0d\xd8\x87\xc7\xd6\xfe\x61\xc4\xb9\x33\x9f\x2c\xe2\x37\xad\x31\x35\x2f\xdb\xa7\x1b\x01\xfb\xa9\xac\x77\x2f\x26\xd3\xfa\x32\x54\x41\x0b\x15\x5d\xeb\x5e\x5c\x26\xc9\x31\x5c\xa3\x16\x2e\x2a\xc3\x5e\x5c\x26\x89\x31\x5c\xe7\x2d\x5c\x54\x76\xdd\xb8\x3a\xd7\xd4\xc4\x99\x5c\xdf\x2d\xf8\x76\xe5\x4e\xae\x6d\x0f\x3e\x11\xa5\x36\x7c\x23\x0d\x62\xa9\x77\x6c\xef\x50\x2f\x2f\xf0\x08\xe0\x8a\x66\xa5\x0d\xf5\x46\x9d\xe6\x84\x12\xdb\x4a\xad\xb1\xec\x7b\x92\xf3\xc7\xfb\xd3\x6b\x2c\xc0\x9e\xf4\xce\x82\xfd\xe9\x35\x16\xa8\x9f\xde\x91\x08\xca\x6c\xec\x34\xe8\x8e\x47\x7b\xd2\x6d\xab\xda\x3e\x33\x15\xba\x77\x00\xcd\x3d\x66\x6b\xa0\xb9\xf7\x3c\x79\x4a\x74\x20\xc9\x2a\x9f\xda\x8f\xe2\x35\xc1\xe5\xd2\x48\x51\x41\x9e\xb1\x00\x73\x47\xcc\x59\x15\x8d\x36\x43\x65\x72\x04\x9a\xa1\xcc\x9a\x3c\xe0\xfc\xc9\xa2\xb1\xf6\x07\x8c\x73\x6c\x1f\x5f\xa4\x11\xe1\x91\x41\x31\xb1\x0a\xc6\x97\xf5\x58\x16\xc4\x5a\x40\xab\x80\x18\x45\x29\xfa\x0f\x8c\x2d\x94\xa5\x28\x83\xde\x71\x27\xfd\x9c\x5c\x69\x2c\xd4\x12\x14\xc1\xf2\x36\xf2\x3c\x80\x96\x64\x23\x8a\x86\x33\x93\xe5\x8c\x21\x3e\xcb\x6e\x2e\xa8\x8c\xda\x9e\x8c\xc7\xe0\xb7\xe4\xae\x29\x0c\x33\x8a\x5c\x53\x86\x56\x58\x69\x0c\xd8\x54\x5c\xcd\x78\xcb\x98\x4c\x2b\xf1\x1a\x5d\xf3\x93\xb3\x20\xf0\x83\xe0\xcd\xe8\xfb\x60\x9b\x88\x96\x51\x46\x65\x41\x69\x58\x3c\xb0\xb5\x16\xf0\x39\xcf\x62\x2b\xb0\xee\xd1\x7d\xc4\x82\x5e\x2a\x20\x1e\x67\xc2\xe9\xd4\x0f\x79\xac\x68\x4e\x89\x90\x13\x16\x5e\x01\xc9\xcf\x22\xf8\x45\xa7\x90\x3e\x92\xf1\x3b\x24\x36\x11\xbf\x5c\xf6\xb2\x92\x47\xd1\xa9\xe2\x2d\x7b\x52\x77\x14\xb5\x2e\x60\xdc\x46\xd5\x25\x3a\x25\x03\x75\x7c\xc7\x92\x29\x15\x95\x36\x51\xa9\x98\x5c\xd6\x6a\x8a\xc7\x72\xbc\x86\x8e\x80\xdb\xbb\x8d\x21\xe7\x23\x61\x77\x64\x5b\x29\x17\xbc\x03\xc3\x4d\x33\x33\x24\xb8\x34\x25\x86\x3b\xe8\xa0\xac\xb5\x98\x8c\x49\x65\xfe\xe5\xe5\x68\x2f\x03\x93\x88\xad\x65\x94\xa6\x30\xb6\x9e\x10\x79\xc8\x4b\x62\x29\x22\x3c\x76\x36\xcd\xcd\x6a\xe8\x84\x4a\xc1\x4e\x70\x1e\x2a\x25\xb4\x86\xb4\x4e\x7d\x26\xad\x10\xce\xc0\xb0\x4a\xfa\x6b\x19\x51\xf1\x9d\x9e\x3a\x6b\xa1\x9e\x03\x65\x30\x7f\xd3\xdc\x98\x9b\xe0\x47\x60\xf8\x96\x9c\x36\x9f\x4d\x86\x5c\xd7\x0b\x10\x84\x1a\xec\x69\x4b\xf8\x9c\xeb\x0c\xd8\x68\x50\x38\xdf\x35\x6b\x56\x1a\x74\xc6\xc7\xe6\x60\xc8\xfe\x8f\x40\x6d\xec\x21\x99\x50\xa6\x2b\x5d\x50\xb2\x7c\x2f\x85\xd9\x3d\x79\xe0\xda\x41\xa1\x12\x53\x62\xa5\x42\xdc\xc2\x3b\xaa\x95\x19\x00\x6a\x91\x93\x73\x90\x38\x75\x46\x55\x02\x8d\xd9\x72\x9a\x85\xe5\x40\x7b\x5a\x2b\x49\x03\x51\x34\x28\x9d\x23\x13\x81\x41\xe9\x38\xeb\x65\x9e\x11\x94\x95\xd0\x22\x9b\x4d\x0e\x74\xe6\x16\x18\x46\xbf\x6d\x36\x28\xb1\xf3\x6a\x6d\x78\xf1\xc4\x34\xaf\xb0\x37\xfb\x71\xf3\x53\xe2\xf0\x6a\x4d\x43\x66\xab\xb2\x78\x68\x5a\xbd\xd3\x89\x48\x23\xe8\x6a\xce\xc2\xd9\x98\x6a\x33\xc2\x7f\x77\x94\x53\x50\x86\x0a\xbd\x5b\x42\x9d\x32\x5f\x41\x24\xca\x44\xcc\x79\x35\x92\x79\x57\x75\x4d\x83\xae\x02\x27\xf3\x3a\xa8\xd2\x8c\x23\xd0\x39\x6c\x8b\xe5\x26\x28\x85\x16\x8a\x61\x46\x50\x82\x20\x96\x5b\x34\xc7\x6b\x1d\x77\xd3\xdf\x08\x1d\x2a\x3a\x89\x9f\xfa\x61\x41\x0d\xb0\x90\x06\x58\xa7\x5b\xc8\x5b\x3e\x44\xf8\xc7\x3c\x86\xef\x88\x5d\x38\xce\x66\xb3\x8b\x38\x9c\x50\x8b\x8f\x48\xbd\xa8\xea\x6e\x53\xab\x47\x87\x2b\xc4\xf0\xf7\x12\x61\xc8\xc3\x26\x6d\x95\x78\xf1\xd9\xa8\x89\x7c\xe9\xf0\xa9\xa6\xf4\x67\x81\x8d\x9d\xb0\x59\x3b\xa9\xed\x01\x0d\xa0\x73\x04\x86\x74\xc1\x8e\x8a\xed\x0b\x02\xd3\xd8\x3a\x1e\xc0\xc1\xb1\x12\x2f\x91\xee\xd0\x84\xfa\x9f\x5f\x59\xf5\xaf\x19\xf7\x21\xc3\x46\xc5\x0d\x44\x2f\x15\xaa\x9a\xd8\xd8\x46\xc9\x77\x22\x48\x64\xef\x10\x7b\xd0\xb9\xb5\xe8\x4c\x68\x4b\x2e\x6c\x4d\x27\xef\xa8\xfd\xa5\x2d\x1b\x18\xaf\x5d\x5f\xb3\x5e\x6f\x2b\x26\xb5\x50\x56\x90\x28\x5b\xc2\x3c\x51\xaa\xe7\x32\x38\x52\x0b\xba\xb7\x77\x55\x87\x61\x58\x15\xc2\xf0\x94\x54\xa6\x54\x99\xa7\x58\x4c\x55\x4b\xf1\x60\xc0\xac\xae\x98\xbe\x79\x13\xfc\x30\x7e\x79\x29\x66\xe0\xcd\xf8\x2c\x18\x3a\x6b\x04\x8a\x1a\x7f\x66\x02\x03\x76\x31\x9d\xfa\x43\x67\x90\x0d\xec\xf1\x9b\x37\x67\xe3\x53\x9b\x61\x61\x0f\x4f\x39\x1a\xea\x0e\xd1\xd4\x0f\xce\x9d\x35\xe4\xfe\x0b\x55\x9e\x8d\xbd\x08\x86\x23\xe5\xcd\x6c\x36\x3e\x39\xf3\x5f\xfc\x1f\x02\x75\x0c\x43\xad\x0e\xf2\x83\x13\xff\xcd\x4b\x10\x8c\xaa\x51\xca\x9b\xf3\x93\xef\x5f\x82\xd1\xd0\xe5\xa3\xc6\x67\x2f\x94\xf2\x46\xc5\xaf\x3f\x93\x0f\x36\x9b\x46\x5e\x2d\xe2\x2d\x45\x19\x6d\xdf\x85\x95\x48\x5d\x7f\x9b\x91\x9e\x02\x31\xd6\x91\x21\x11\x06\x43\x37\x57\x76\x47\x37\x92\xa6\x55\x85\x8d\x21\x9e\x0a\xb0\x90\x0a\x79\x1d\xdd\xe6\x83\xc1\x1d\x80\xb7\xf8\x6e\xa3\x2a\x95\x54\xce\x7e\xbd\x32\xf4\x53\x2a\x2c\x8d\xd6\x89\xec\x81\x74\x04\xe3\x9c\x4b\xca\xe2\x5c\xb4\x51\x45\x09\x78\x0e\x86\x4d\x40\x19\x30\x9b\x3a\x1d\xbc\x73\xa1\x94\x58\xb5\xbe\x10\x71\x34\x36\x94\x19\x31\xfa\x9d\xd9\x06\x63\xae\x13\x56\x70\xdb\x05\x2e\x67\xd3\x89\x41\x59\x65\x43\xb8\x2e\x85\x41\x3a\x11\x08\xa7\xdb\x3f\x83\x6a\xf1\x3b\xb1\x88\x23\x11\x9d\x6d\x24\xa5\xf3\xeb\x4c\xa7\xc1\x68\x36\x0b\x46\xbd\xd8\xd8\xe0\x4e\x74\x9c\xab\x5b\x72\xb7\x8d\x23\xad\x9a\x68\x62\x89\x55\x11\xa7\x53\x7f\x3c\x9b\xf9\xe3\xad\x3c\xf5\x20\x94\x4c\xbd\x34\x7e\x0e\xfc\xbb\xe9\xf4\x7c\x1b\xa3\x5a\x41\x70\x67\xbc\xda\xb3\xe0\x8e\xce\x44\x7b\x78\x76\x47\x45\xbe\x75\x6a\x3d\x2c\x28\x07\x08\x9c\xd9\x6c\xa6\xdb\x56\x6b\x36\x5a\x39\xb2\x61\x66\xbc\xbb\xa9\x63\x75\xf5\x27\x83\x51\x8f\xd1\x09\x96\x77\xa7\xc3\x27\xd8\x20\x24\x1e\x6d\xa5\x64\xaa\xd8\x56\xad\xe6\xdb\x61\x55\x93\x50\xe7\x12\x4a\x67\x56\x35\xb5\x6f\x87\xfd\x8a\x6a\x2a\xe1\x0a\x02\x7a\x0f\xfb\xed\x70\xe2\x1b\x69\x76\x8d\xf7\x27\x06\x1e\x07\x23\x9d\xcb\xf1\xa8\x8f\x4b\x53\x2f\xa7\x3e\xc9\x22\x14\x14\xc0\x7e\x04\xba\x51\xef\x8f\xa1\x5d\x29\x6e\x63\x08\x75\x33\x01\x70\x36\xeb\x36\x40\xc9\xd9\x7f\x05\x71\xbb\xd4\xbc\x2b\xde\x50\xb7\x6c\xfa\xd4\x1f\x87\xba\x6d\xd3\xc7\x3d\xd6\x2d\xa7\xf7\x27\x60\xa3\x5d\x07\x37\x84\xae\x34\x9a\xc9\x9f\x0c\x75\xa5\xc1\xc8\x85\xec\xd8\x41\xb7\xb5\xca\xc9\xf6\x90\x11\x66\xdf\xa6\x53\x3b\x84\xdd\x08\x99\xca\xeb\xd4\x70\x6b\xab\x17\x42\x6d\xcc\xaf\x76\x1d\xbb\xe0\x6f\xcd\x44\xb5\xd7\x5e\xfc\x46\xcf\x61\x96\x6b\xaf\xf3\xe8\xe6\xf2\x1e\x56\xe9\xd4\xa5\xcc\x6b\x4d\x05\xaf\x46\x0c\x31\x6d\x86\x26\x83\xdd\x72\xf0\x2d\xb9\x5c\xcd\x9d\x85\x0a\x8b\xe4\xb9\x55\x3c\xe4\x98\x58\x24\xb7\x96\x79\x46\x22\x94\x59\x51\xa6\x64\xdf\xa2\xbe\x4a\xc0\xf1\xb1\x56\x1a\xec\x62\xa1\xaa\x18\x0e\x00\x4f\x8a\xbc\x04\xe7\x8f\x3f\x56\xb9\x47\x63\x3b\xab\x1a\xda\x1d\x33\x84\x8e\xb3\x69\xd4\x33\x4c\x72\x15\xc7\xd9\xcc\xe9\xf2\x69\xe7\x86\x53\x15\xb3\xea\x84\x18\x3b\x6f\xf5\x27\x03\xc8\x4e\xaa\x76\x52\x66\xe7\xe3\x5a\x87\xcf\xaa\xc3\x8a\x70\xa0\xd1\xa6\x69\xab\xb7\x58\xb0\xfd\x25\xdc\x61\x5e\x85\x9e\x51\x32\xf4\x03\xc3\x7e\xc6\x0b\x00\xe6\xe7\x05\x5d\x36\x9e\x00\x0e\x43\xa2\x14\xd1\x2e\xaf\x6e\x58\x3d\x1e\x00\x50\x1f\x4a\xf3\xea\x33\x69\xc6\xd8\xb6\x3e\xe1\xe2\x92\x01\x72\x36\x3c\x51\xcd\xa6\x88\x4b\x3c\x57\xcb\x87\x4a\x1c\x3b\xc8\xaa\x2c\x35\x9a\xd2\xd4\x70\x9d\x83\xa8\xce\x4c\x93\xbe\xe1\x34\x47\x5c\xe7\xc0\x8e\x4e\xce\x7c\x67\x3a\x1d\xbf\x24\x27\xe3\xb3\x1a\xb6\xec\x85\x1d\x0d\x2b\x58\xff\x0d\x8d\x5a\x83\x17\x9b\x42\x33\x34\x65\x03\xcd\xa2\x03\x0d\x03\xfe\x9e\xc2\x9e\x4b\x58\x8a\xa6\x14\x68\x16\x14\x0d\xaf\x18\x8a\x2c\xb7\x30\xab\x7d\x5e\xa5\xba\xf9\x29\x60\x03\xc3\x8e\x71\x76\x3e\x9b\xd1\x8c\x9c\xe5\xe1\xae\x9d\x9f\xd8\x3e\xcf\xc6\x7d\x67\x50\x25\xe4\x9b\x46\x53\xc2\xac\x3c\x28\x8b\x11\xd6\x7b\x04\x62\x41\x75\xcd\x24\xdd\x6e\x6b\x3e\xff\xa2\x97\x71\xfa\xf0\x34\x0c\x78\x2b\xd2\x79\x0a\xb3\x5d\xe2\xe8\x36\x9d\x3e\x86\x1f\xa2\x62\x8e\x0c\x4e\xb6\xaa\xc8\xfc\x19\xaa\x93\xbb\xb8\x4f\xc6\xab\x52\xb1\x81\xa2\xac\xbd\xab\xef\x14\xd2\xe4\xc7\xeb\xa4\x27\xc5\x65\x77\xdc\x78\x48\xc9\x61\x99\x67\x05\xb1\x10\xe0\x0d\x15\x2f\xc1\x10\xfe\x07\xda\xe2\xe0\xbb\x13\x0a\x44\x56\x61\x17\x6e\x06\x8e\x8f\xdd\x75\xc2\x5e\x4d\x72\x80\xd8\xb9\x78\x97\xdf\x37\x98\x44\x20\xd8\x80\xf5\x86\x55\xe9\x0b\x62\x25\xd5\x29\xc9\xaa\x32\x62\xfb\xc3\x60\xe4\x54\xe4\x4a\x90\x03\x00\x90\xf7\x39\x7a\x7a\x5b\x78\x24\xe7\x96\x64\x3b\x93\xbf\x5d\x7f\xba\xf2\xb8\xe7\x44\xc9\xb3\x5d\x08\x80\x05\x48\x1a\x85\x3a\xbb\x14\x6f\x62\xfd\x4d\xe6\xf4\xdc\x4d\xb0\x13\xc7\x78\x8b\xc0\x4e\xdc\x45\xeb\xcd\xcf\x11\x79\xb0\x13\x37\x6e\xbd\xb8\x79\x5e\x51\x88\xbc\xf5\x82\x5f\x08\xb0\x13\x37\x12\xec\x3d\x80\xae\x83\xff\x94\x15\xf3\x21\xec\xe6\x1b\xe5\xa4\xb5\x9d\x88\xa3\xd4\xf5\xb5\x9d\xf6\x48\x3b\x71\x1f\x04\xf9\x15\x30\x1d\x9a\xa6\x04\x92\xaa\x9b\x60\xaf\x9c\xf0\xd7\xf3\x7f\xe4\xf8\x37\x88\x03\xaf\x80\x59\x6c\x27\xcd\x63\x8b\xb6\xe3\x6c\xa4\x16\x54\x0a\x55\xd8\xc4\x3d\x3e\xa6\x0a\xc4\x1b\x4f\x39\xbb\x94\x90\x7b\x3f\xc1\x04\x62\x0c\xe3\xcf\x90\x7c\x89\x52\x79\x3d\x41\x3c\x9e\xf8\x2e\x33\xc9\x49\x40\x07\x5f\x94\x69\x82\xd2\x47\x98\x11\x7e\x93\x43\x8e\x8e\x48\x44\x47\x66\xf1\xa7\xe4\x9a\x60\x18\x3d\x4e\x82\x0a\xee\x8c\xc2\xd5\xb7\x1a\x76\xbd\xae\xa0\xc2\xbc\xca\x8d\x85\x06\x42\x79\x69\xc1\xc4\x98\x62\x7f\x36\xe4\xf7\x16\x54\xe0\x43\xae\x2e\x74\xcc\x86\x1f\x45\x78\xc5\xbb\x0b\xe2\x60\x5c\x7d\x79\x41\xf6\x2f\xa4\xa3\x61\xe7\x9f\x35\x96\x98\x2e\x1b\x65\xa1\xdf\x29\xf0\x75\xd0\x28\x8e\xaf\xb5\x69\xb4\x2f\xd5\x8c\x47\xec\x46\x10\xd1\xd9\xd0\x91\xc1\x2c\xee\xe0\xa2\xff\xa2\x40\xad\xc6\xbb\xeb\x97\x80\x78\x25\xed\x92\xe8\xa4\x6e\xb5\x59\x32\x6a\x96\x18\x76\x98\x5e\x19\x66\xf1\xc7\x6a\x95\x64\x88\x69\x8d\x41\x06\x66\x8d\x92\x60\x5f\xaf\x4f\x12\x15\xcc\x62\x23\xfd\x5d\x75\xe9\x33\x2c\x56\x79\x56\xec\x7c\xc5\xaa\x0d\xf9\xca\xba\x25\xd1\xb6\x74\xac\xcd\x6a\xaf\xae\x89\xe1\x5f\xa7\x73\x86\x59\x62\xb6\x8b\xfc\xb7\x6f\x64\xe9\xbb\x56\x75\x37\x6b\x07\xce\xbe\xc1\x35\x2d\x03\x17\x0d\x53\x30\x2c\x95\xe9\xe2\x96\x01\x4d\x14\x57\xf3\xd5\xef\x70\xb5\xef\x30\x52\xeb\x30\x4a\x69\x0b\xe6\x5d\x2f\x76\x19\x70\x28\xd6\x66\x9c\xe2\x56\xab\x8b\x48\xb4\x87\xa5\x45\x24\x7a\x2d\xeb\xa2\xa8\x6a\x8b\x6a\xb0\x61\xb6\xa2\x88\x44\x07\x5a\x4e\x93\x6b\xad\x3d\x77\xf8\xc5\xeb\x96\xa5\xc8\x8c\x4f\xbb\x8e\x5c\x5d\x05\x36\x73\xf2\x91\xe5\x53\xaf\x60\xb8\x6a\xba\x69\xb8\x88\x6c\xa6\xde\xba\xf3\xb3\x27\x71\xbd\x2d\xfa\x60\xba\xa7\xe3\xb6\x1e\x3f\x13\xc8\x75\x7b\xd0\x2b\x3a\x52\xdf\xe4\xe9\x99\x9d\xe2\x03\xe8\x0c\xb9\xdd\x37\x15\xaa\x63\xfb\xa3\xc3\xa3\x38\x7e\xaf\x69\x44\xef\xe5\xea\x0a\x8c\xef\x81\x0c\xd2\x70\x02\x44\xd2\x33\x36\xe4\xeb\x36\xbb\x78\x7a\xea\x87\x78\x06\x86\x21\x66\xa7\x75\xea\x33\xf3\xb7\xf8\xae\x2e\x0e\x36\x7b\xe9\xea\x6c\xf7\x60\xc3\x55\xa6\x4e\x9d\x87\x2e\xa7\x6d\x0e\x43\xc9\x36\x76\xf7\x1b\x0a\xd0\x2b\xb9\x0f\x15\xa3\xf4\x22\x46\xde\x8c\xce\x44\x19\x79\x98\x4f\x51\xe9\x33\x09\x9b\x69\xeb\x7a\x37\x6c\x01\xc3\x2c\xee\x02\xdd\xb6\x14\x75\x5e\xb8\xfb\x4a\xd4\x30\xaf\xb4\x10\x0a\x42\xb9\x0e\x26\xc6\x8c\xcb\x50\x0f\x3c\x6c\x15\x8c\xb3\xf9\x63\x23\x71\x23\x4b\xec\x63\x1b\x07\xc4\x67\xe3\xfd\xe2\x33\xad\x50\x20\x03\xb4\x6e\x9e\xf6\xdd\x06\x4d\xdf\xe0\xd8\x1a\x99\x29\xe4\x99\x2d\x18\xd5\x43\xb7\x94\x33\x5d\x9a\x5f\x9f\xa2\x34\x91\xfd\xda\x5e\x13\x53\x4c\xe7\xb3\x98\xce\x24\xd9\x2e\x9c\xdb\x36\x92\x40\x6e\x24\x0a\x30\xcc\xe2\x0e\xb1\x6c\xbf\xac\x1f\x81\x88\x95\x95\x22\x8f\x15\x7f\x76\x75\x06\xd5\xf0\x57\xf1\x03\x02\x57\xed\x8a\x9b\x9c\x18\xac\xbf\x02\x39\xc4\xf0\xdb\x9c\x3f\x76\x7f\xe5\x60\xdf\xa8\x6e\xdb\xf7\x6c\x04\x71\xee\xf5\xb5\x69\x9a\xe2\x0c\x01\x10\xc5\x86\x8f\x31\xf4\x86\x1a\x02\x92\xee\x11\x2d\x42\xfd\x7a\x21\x6a\xcb\xeb\x4d\x5d\x99\x2e\x29\x18\x7f\x01\xc1\x0e\x07\xdc\x64\xcd\x7b\x9b\x97\xb7\xa1\x18\x5b\x00\x54\xf9\x5f\xdb\xa9\x3f\x4c\x60\x3b\x61\x0a\x89\x95\xb1\x7f\xcb\xb0\x78\x42\x64\xf9\x60\xa3\xda\x2d\xda\x8e\xb3\x5e\x46\x05\xb4\xda\x96\xc6\x8a\x9c\x6b\x7b\x1d\xd3\xff\xb3\x0d\x48\x6e\x8b\x3b\x41\x8d\x80\x3a\x35\x08\x2b\x74\x36\x71\xc2\x12\x10\x25\xba\xb5\x9d\xea\xc4\x7b\x07\x05\x5e\x31\x5d\xdb\x6b\xc8\xfe\x30\xd2\xa8\x56\xa2\x49\xa4\xee\x9c\x10\xa1\x80\xb6\xb3\x8d\x98\x52\xac\xb5\xd7\x30\x8b\x6b\x7a\x0c\x2e\xac\xee\xa4\x4d\x5a\xcd\x99\x5f\xb2\xdf\xb2\xfc\x29\xb3\x2a\x4a\x16\x86\x4b\x04\xbf\xc0\xd8\x4a\x70\xfe\x68\xe1\x32\x23\xe8\x11\x1e\xb3\x73\xac\x19\x00\xa0\xcc\x62\x98\xa0\x0c\xc6\xc2\xee\x36\x99\x5d\x2a\x05\xea\x05\xbf\xa0\xe6\x22\x67\x9d\xdc\x92\x3b\xc0\x05\x0c\x5d\x2e\x04\xec\x52\xd6\xd0\xa6\x1e\x1f\xd7\x8d\x2b\xdb\x01\xb3\x56\xf3\x88\x38\xca\xe0\x07\x55\xd3\xa8\x62\x0a\x4d\xda\xa6\x75\xd0\xa9\x35\x69\x9f\x0a\x0c\xbb\xf3\xc4\xd5\xaa\x50\xca\x21\x8a\x5e\x69\x59\x79\xb5\xe8\x4d\x3d\xe2\x6b\x2c\x10\x28\x5d\xf4\x9f\x71\xfe\x88\x0a\xe8\x61\xc8\x0c\xcd\xbc\xf0\xce\xc6\x4c\x4a\x16\xf5\x35\x6a\xe2\x79\x83\xa0\x14\x9a\xd9\x8c\x94\xe3\x6c\x15\x4f\x36\xfb\x9a\x0c\x98\xad\xa5\xf0\x11\x15\x7e\x0c\x53\x48\xa0\x95\xdc\xc2\x3b\xd5\x53\xd3\x65\x5a\xd8\xd0\x45\xde\x02\x65\x74\x49\xc5\x5f\x58\xfe\x15\x53\x55\x83\xff\x5e\xc1\x25\x81\xb1\xa5\xe8\xab\x95\xe4\xd8\x5a\x31\xaa\x28\x41\x30\xb6\xe2\x6a\x02\xc7\x74\xee\xce\x46\x68\xee\x36\x99\x1d\x5f\x44\x28\x85\xb1\x45\x72\x2b\x86\xcb\x3c\xa6\xaa\xcc\x57\xb1\x52\x65\xf8\x7b\x09\x0b\x72\xec\x38\x9b\x4d\xdd\x71\xc1\x70\xf9\xc5\x2e\x79\xb3\x7f\x05\x56\x6c\xc3\x5b\x79\xe2\x1b\x4e\xbb\xee\x79\x35\xc4\xab\x6c\x7b\x0a\x3a\xa9\x9a\x6d\x96\x0c\x9b\x5f\x0d\x78\xc8\xfe\x67\x9c\x45\x89\xf7\xae\xb3\x1d\xb2\xfd\x19\x69\x53\xd7\x93\x97\xfb\x7e\x4a\xad\x2f\xba\xad\x6f\x07\x29\x25\x0c\x85\x36\xdb\x64\x0d\xa2\x36\x15\xf5\x14\xb0\x28\x8e\x7f\xc1\x5b\xeb\x6d\x62\xff\x6d\x02\xde\xe8\x93\x6c\x87\x8b\x67\x81\xac\xd5\x29\xb0\x30\x8b\x8d\x9c\xf6\xef\xe0\xd2\x9e\xd3\xda\x99\x22\xad\x57\xfc\xc6\x0f\xd4\x9d\xb7\xd1\xdf\x25\x4e\xa7\xbc\xa8\xb7\x6c\x09\xc5\x46\xae\x6c\x23\x67\xc0\xcc\x3e\x05\xec\xe8\xc5\xa2\xae\x5e\x2c\x12\xbd\x58\xa9\xf1\x86\x56\x2c\x72\x33\x41\x3b\x37\xb7\x62\x91\x13\x22\xd1\x8a\xcd\xa5\x9c\x1e\x6c\xba\x19\x69\x7d\x59\xd4\xea\xcb\xf2\xf6\xeb\x12\x2c\x99\xdb\x58\x7a\x17\x28\x85\xd7\xcf\x05\x81\x8f\x5a\x03\x56\x7c\xfb\x6d\xe2\xbb\x3f\xb1\x83\x25\x39\x7e\x56\x9b\xaa\x1c\x94\x8e\xd8\xd5\xe3\xd4\x10\xaf\xe2\x71\x14\x74\x75\x20\xd6\x62\xc9\xe0\x71\x6a\xc0\x43\x3c\x8e\x71\x16\x8b\xa8\x80\x59\xf4\xf8\x4d\xa2\x6e\x23\x03\x87\x7c\x41\xf3\x90\x4f\x58\x1a\x89\xa3\x22\x46\x78\x4f\x8f\x77\xae\x78\xbc\xa3\x23\xd5\xe7\x99\xbe\x1e\x59\x7d\x30\x4a\xa1\xce\xd3\xe7\xf6\x72\x9b\x92\x67\x05\x2c\x8a\xe3\xf7\xed\xb5\xea\x75\x7d\x4d\xe8\x5d\x3e\x3e\xe9\x1b\x21\x2f\x9b\x52\x32\x65\xd8\x81\x3b\x80\xee\x80\x4d\xb6\x09\x4e\x73\x62\xd3\x5c\xb7\x64\x3e\x9e\xb4\xdc\xdd\x6d\x54\x82\xbc\x92\x91\xd6\xf8\xea\x90\xb5\xcd\x95\xd1\x4c\xe5\xb8\xc3\xec\xd4\x34\x93\x6f\x6b\xa8\x26\x0e\xbe\x9d\xa5\x9a\xa8\x27\x28\x6d\xd5\xf7\x9b\x17\x4d\xb7\x58\x2b\x7e\xdb\xe7\x46\x25\x0e\x71\x1a\xb1\xbf\x1f\x84\x9d\x01\xf9\x6e\x24\x3b\x1b\x3b\xb0\x7e\x50\x8b\xe8\x7c\xff\x16\x91\xca\x02\xef\xa2\x18\x74\xd6\xec\x6b\x6a\xc0\x83\x9c\x4d\x03\x7c\x3f\x6f\xd3\x00\xbd\x30\xac\x74\x67\x01\x4e\x85\xe5\xa1\x13\x03\xdf\xa1\x8f\x32\xaa\xdb\x39\xa3\xdd\xda\x39\xe2\x93\x32\x3d\x0d\x9d\x96\xf4\xf7\xe0\xc6\x1d\x69\x08\x60\x16\x9b\x17\x6f\x9b\xf3\x54\xbf\x7e\xbb\x7f\x8c\xf3\x6a\x99\x95\x86\xb2\x15\xeb\x6c\xcd\xb0\x9a\x08\xbe\x2e\xe6\xf9\x83\x3e\xdc\xad\x31\xd0\xd8\xfc\x77\x38\x30\xa6\x81\xef\xf9\x0d\x69\x0d\x5a\xd9\x8b\xf7\xcc\x63\x96\x9e\xf8\x82\xf2\xee\xea\x24\x20\x5e\x49\x95\x24\x3a\xa9\x46\x6d\x96\x8c\x2a\x24\x86\x1d\xa6\x3e\x86\x59\x7c\x3b\xd5\x91\xc4\x99\x5a\x18\xe6\x6b\x56\x19\x09\xb6\xb7\xba\x48\x48\x98\xc5\x46\x72\xbb\x78\x1e\x91\x94\xed\x77\x04\xcc\x04\xfb\x8a\x5e\x48\x43\xdc\xf0\x45\x5d\x0c\x77\x7a\xa4\x26\xc0\xe1\x7e\xa9\x73\xb6\xdf\xe6\x30\x58\x3b\x87\xae\xba\x8d\x3b\x71\xf7\x0d\x0e\x84\x19\xf9\x90\x2e\xb4\x73\xd9\x4c\xf5\x23\x23\xaa\x3d\x8f\x85\x75\xc8\x6b\x2b\xf6\xfd\xb2\x2e\x0d\x4b\xe5\xb4\x7b\x26\xbb\xc5\x26\xd3\xa8\x28\xac\x39\x2f\x3e\xe1\x92\x47\xb2\xd2\x78\x98\x2f\x23\xe2\xd3\xc9\x34\xcd\x83\x9b\x0d\x07\x79\xde\x01\x84\x07\xe5\x50\xe9\x1a\xdc\x1f\x5c\xe8\xda\xb6\x53\xda\xa8\x35\x46\xdc\x5b\x68\x54\xbd\xfa\x37\xbd\xaf\xaa\x7e\x35\xa2\x91\xd7\xa8\x80\x89\x0e\xdf\xd6\xe2\x97\xbc\x4b\xb2\xad\xf9\x92\x38\xf5\x35\x92\x7d\x3d\x9f\x5d\xca\x06\xcc\xc2\xd4\x80\x31\xe8\xbf\xac\xb5\x35\xfa\x22\xf5\x12\x84\x8b\x76\x23\x86\x8e\x98\x53\x07\x49\x97\xce\x71\x09\x57\x3c\x5b\xf4\x5e\x0c\x54\x0c\x8d\x1e\xd1\xcc\x53\xf1\xb7\xda\x59\x4a\x63\x67\xd3\xd3\x93\x93\x8d\x12\xd9\xc3\x20\x39\xbb\x3a\x85\xb2\x24\x3f\x56\x7b\x62\x97\x5f\xa1\xdd\xe6\xcd\x9c\x6b\xb5\xbe\x63\xeb\x1a\x6d\xda\x97\xbf\x4a\x93\x65\x50\xf4\x7f\x5a\x1c\xa5\x75\xca\xa5\xab\xb1\x7c\xd1\xd0\x33\xb5\xfd\x5e\x35\xda\xd5\x1c\x9f\x4e\x94\xa6\x93\x29\x24\xec\x0b\x49\x8d\xaf\xec\x00\xbf\x56\xa0\x0a\xca\xc6\x4e\x48\x73\x4a\xc0\x0d\x03\x09\xc3\x40\xb5\x61\x28\x96\xf3\x5c\x5b\x0e\xfc\xf3\x5a\x0c\xdd\x93\x2e\xc0\x3a\xcd\xef\x27\x99\xcb\xaf\xd7\x4d\x90\xcb\xae\xa7\x4f\x0a\x97\x46\x24\x93\xd4\xa5\x83\x27\xf7\x6e\x8c\xf0\xe4\x72\x13\xc2\x7f\xaf\x72\x4c\xac\x8a\xaa\x75\x11\xfe\xbf\xff\x0d\x00\x00\xff\xff\xa3\xdb\x2c\x22\xa2\x6a\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/std.js"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
