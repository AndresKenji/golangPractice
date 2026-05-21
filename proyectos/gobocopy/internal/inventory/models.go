package inventory

type FileData struct {
	Path    string
	Size    int64
	ModTime string
	Mode    string
	IsDir   bool
}

func (f *FileData) Compare(t *FileData) bool {
	return f.Path == t.Path &&
		f.Size == t.Size &&
		f.ModTime == t.ModTime &&
		f.Mode == t.Mode &&
		f.IsDir == t.IsDir
}
