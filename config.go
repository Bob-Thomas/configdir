// configdir provides access to configuration folder in each platforms.
//
// System wide configuration folders:
//
//   - Windows: %PROGRAMDATA% (C:\ProgramData)
//   - Linux/BSDs: ${XDG_CONFIG_DIRS} (/etc/xdg)
//   - MacOSX: "/Library/Application Support"
//
// User wide configuration folders:
//
//   - Windows: %APPDATA% (C:\Users\<User>\AppData\Roaming)
//   - Linux/BSDs: ${XDG_CONFIG_HOME} (${HOME}/.config)
//   - MacOSX: "${HOME}/Library/Application Support"
//
// User wide cache folders:
//
//   - Windows: %LOCALAPPDATA% (C:\Users\<User>\AppData\Local)
//   - Linux/BSDs: ${XDG_CACHE_HOME} (${HOME}/.cache)
//   - MacOSX: "${HOME}/Library/Caches"
//
// configdir returns paths inside the above folders.

package configdir

import "path/filepath"

func dir(root, vendorName, applicationName string) string {
	if vendorName != "" && hasVendorName {
		return filepath.Join(root, vendorName, applicationName)
	}
	return filepath.Join(root, applicationName)
}

func CacheDir(vendorName, applicationName string) string {
	return dir(cacheFolder, vendorName, applicationName)
}

func SystemSettingsDir(vendorName, applicationName string) string {
	return dir(systemSettingFolders[0], vendorName, applicationName)
}

func SettingsDir(vendorName, applicationName string) string {
	return dir(globalSettingFolder, vendorName, applicationName)
}
