package mp3tag

import (
	"io/ioutil"
	"github.com/eugeis/gee/lg"
	"github.com/mikkyang/id3-go"
	"strings"
	"os"
	"path/filepath"
)

var Log = lg.NewLogger("mp3tag ")

func FileNameToTitle(folder string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		Log.Err("io access to '%v' not possible '%v'", folder, err)
		return
	}

	for _, f := range files {
		if mp3File, err := id3.Open(folder + "/" + f.Name()); err == nil {
			title := fileNameOnly(f.Name())
			mp3File.SetTitle(title)
			Log.Info("Set title '%v' to '%v'", title, f.Name())
			mp3File.Close()
		} else {
			Log.Err("id3 access to '%v' not possible '%v'", f.Name(), err)
		}
	}
}

func FileNamePrefixToTitle(folder string, separator string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		Log.Err("io access to '%v' not possible '%v'", folder, err)
		return
	}

	for _, f := range files {
		if mp3File, err := id3.Open(folder + "/" + f.Name()); err == nil {
			if titleParts := splitToParts(f.Name(), separator); len(titleParts) > 0 {
				mp3File.SetTitle(titleParts[0])
				Log.Info("Set title '%v' to '%v'", titleParts[0], f.Name())
			}
			mp3File.Close()
		} else {
			Log.Err("id3 access to '%v' not possible '%v'", f.Name(), err)
		}
	}
}

func FileNamePrefixToFileName(folder string, separator string) {
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		Log.Err("io access to '%v' not possible '%v'", folder, err)
		return
	}

	for _, f := range files {
		if titleParts := splitToParts(f.Name(), separator); len(titleParts) > 0 {
			newFileName := folder + "/" + titleParts[0] + filepath.Ext(f.Name())
			os.Rename(folder+"/"+f.Name(), newFileName)
			Log.Info("Change file name '%v' to '%v'", f.Name(), newFileName)
		} else {
			Log.Err("id3 access to '%v' not possible '%v'", f.Name(), err)
		}
	}
}

func splitToParts(fileName string, separator string) []string {
	return strings.Split(fileNameOnly(fileName), separator)
}
func fileNameOnly(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
