// Code generated by vfsgen; DO NOT EDIT.

package config

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

// internalAssets statically implements the virtual filesystem provided to vfsgen.
var internalAssets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 5, 20, 7, 54, 48, 362978723, time.UTC),
		},
		"/zsys.conf": &vfsgen۰CompressedFileInfo{
			name:             "zsys.conf",
			modTime:          time.Date(2020, 5, 20, 9, 10, 27, 291604420, time.UTC),
			uncompressedSize: 1094,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xcf\x8a\xdb\x4c\x10\xc4\xef\x7a\x8a\x82\xbd\x7c\xdf\x21\xc1\x7f\xb2\x1b\xd0\x2d\xb0\xb7\x64\x43\x0e\x81\x90\x63\x5b\x2a\x59\x83\x35\x33\x4a\x4f\xcb\xe0\xb7\x0f\x33\x96\xb3\xc6\xab\x40\xac\x93\x34\x3d\xf5\xeb\xea\x56\x77\xef\x92\x45\x3d\xd5\x15\xf0\x80\xcf\xe4\x08\x31\x0c\x94\x64\x08\x98\x83\x60\x30\x3d\x61\xa4\x62\x0a\xce\x10\x3b\x98\xf3\x84\xeb\xc0\x10\xa7\x7d\x5f\x4e\x7a\x7a\x88\x12\xa3\x32\x31\x58\x01\x7e\xef\x89\xa8\x2d\x15\x4d\x0c\xad\x33\x17\x43\xbe\x88\xdd\xd4\x1c\x68\x48\x26\x6a\x90\xd0\x82\xa1\x45\x2b\xc6\x84\xff\x3a\x8d\x1e\x3e\x26\x83\xb2\x61\x30\x58\x44\x1c\x5a\x26\xfb\xbf\x02\xf6\x4d\x11\x49\x67\xd4\x1a\xeb\x0a\x38\x90\xe3\x20\xc9\x6a\x6c\x56\x78\xc0\x8b\x0b\xce\x4f\x1e\x61\xf2\x3b\x6a\x76\x36\x63\x92\x15\xbe\xc5\xa2\x78\x5f\xfc\x01\x78\x87\x20\x9e\x35\xae\x9f\x4f\x3b\x67\x2a\x7a\x2a\xa1\xb9\xb8\xd9\xf3\x45\x86\xf9\x3b\x5d\x29\xbf\xfe\x49\x39\xc7\x10\x8f\xd4\x22\x76\xc1\xa8\x47\x19\x6e\xe5\x03\xc3\xde\xfa\x33\xe3\x4b\x79\xcf\x72\x4a\xd3\x5f\x7a\xe4\x02\x5a\x39\xa5\x57\x61\x12\x3f\x0e\x4c\x23\xf5\x7c\xa3\xbe\xca\xdb\x8a\x49\xca\x89\xe7\x2a\xb3\xfa\x0a\x56\xfa\xa7\xd3\xc0\x94\xff\xf7\x6b\xed\xdf\x94\x47\x17\xa7\xf4\x2c\xa7\xea\xa6\xb8\x75\xb5\x64\xf7\x72\xfa\xd6\xcb\x76\x11\xfc\x83\x3c\xdc\x92\x1f\xef\x24\xaf\x17\xc9\x2f\x31\x58\x7f\x8b\xfe\xb0\x88\xfe\x78\x27\xfa\x27\x45\xdf\xb4\x63\xb9\x1f\xdb\xd5\x9d\xec\x4d\x86\xa7\x7f\x6b\xf6\xf6\xe9\xf1\xaf\xf4\x4d\xb5\x67\xa0\xca\x70\x5e\xe0\x32\xfc\x32\xa0\x53\x12\x69\x94\x86\x50\xfe\x9a\x9c\xb2\xc5\x8e\x5d\x54\xc2\xe4\xe0\xc2\x1e\x82\x14\x64\x4c\x7d\xcc\x43\xe1\x5d\xc8\x8a\x31\xc6\xa1\x88\xf2\x2a\x15\xde\xb3\xd0\xe7\x95\x75\x9e\x71\x2a\xb3\x98\x98\x37\x39\x3b\x9f\x0f\x6b\x3c\xad\xaa\xdf\x01\x00\x00\xff\xff\x95\xe2\x67\xc7\x46\x04\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/zsys.conf"].(os.FileInfo),
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
