package main

type Directory struct {
	Name     string
	Children []*Directory
	Parent   *Directory
	Files    map[string]int // Mao from File name to size
}

func New() *Directory {
	return &Directory{
		Name: "Root",
	}
}

func (d *Directory) InsertDirectory(name string, parent *Directory) *Directory {
	child := &Directory{Name: name}
	child.Parent = parent
	d.Children = append(d.Children, child)
	return d
}

func (d *Directory) InsertFile(name string, size int) {
	if d.Files == nil {
		d.Files = make(map[string]int)
	}
	d.Files[name] = size
}

func (d *Directory) Size() int {
	size := 0
	for _, child := range d.Children {
		size += child.Size()
	}
	for _, fileSize := range d.Files {
		size += fileSize
	}
	return size
}

func (d *Directory) ContainsChild(name string) (bool, *Directory) {
	for _, child := range d.Children {
		if child.Name == name {
			return true, child
		}
	}
	return false, nil
}
