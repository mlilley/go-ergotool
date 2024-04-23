package ergotool

import (
	"bufio"
	"os"
	"sort"

	"github.com/labstack/gommon/log"
)

func ReadDoc(filename string) (*Doc, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	doc, err := NewDoc(bufio.NewReader(f))
	return doc, err
}

func WriteDoc(filename string, doc *Doc) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(doc.String())
	return err
}

// func BackupFile(filename string) error {
// 	// get the non-extension portion of the filename
// 	ext := filepath.Ext(filename)
// 	base := strings.TrimSuffix(filename, ext) + "." + time.Now().Format("2006-01-02 03:04:05PM") + ".bak" + ext

// 	// err := os.Rename(filename, filename+".bak")

// 	return nil
// }

func UpdateFootprintLocations(src *Doc, dest *Doc) error {
	srcFootprints, err := src.EnumerateFootprints()
	if err != nil {
		return err
	}

	destFootprints, err := dest.EnumerateFootprints()
	if err != nil {
		return err
	}

	keys := make([]string, 0, len(destFootprints))
	for k, _ := range destFootprints {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		destFp := destFootprints[k]
		if srcFp, has := srcFootprints[destFp.Ref().Value()]; has {
			x := srcFp.At().X()
			y := srcFp.At().Y()
			r := srcFp.At().R()

			destFp.At().SetX(x)
			destFp.At().SetY(y)
			destFp.At().SetR(r)
		} else {
			log.Warnf("failued to update footprint '%s': not found in source", destFp.Ref().Value())
		}
	}

	return nil
}
