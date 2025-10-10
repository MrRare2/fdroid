package apps

import (
	"io/fs"
	"os"
	"path/filepath"
)

// FindIcon searches for an icon file in common locations within a directory.
func FindIcon(dir string) (string, bool) {
	var foundPath string

	// Common locations for icons
	searchPaths := []string{
		filepath.Join(dir, "fastlane/metadata/android/en-US/images/icon.png"),
		filepath.Join(dir, "app/src/main/res/mipmap-hdpi/ic_launcher.png"),
		filepath.Join(dir, "app/src/main/res/mipmap-mhdpi/ic_launcher.png"),
		filepath.Join(dir, "app/src/main/res/mipmap-xhdpi/ic_launcher.png"),
		filepath.Join(dir, "app/src/main/res/mipmap-xxhdpi/ic_launcher.png"),
		filepath.Join(dir, "app/src/main/res/mipmap-xxxhdpi/ic_launcher.png"),
		filepath.Join(dir, "icon.png"),
	}

	for _, sp := range searchPaths {
		if _, err := os.Stat(sp); err == nil {
			return sp, true
		}
	}

	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || foundPath != "" {
			return err
		}
		if !d.IsDir() && d.Name() == "icon.png" {
			foundPath = path
		}
		return nil
	})

	return foundPath, foundPath != ""
}
