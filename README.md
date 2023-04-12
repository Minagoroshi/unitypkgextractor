# Unity Package Extractor

This is a command-line tool for extracting files from Unity package files (`.unitypackage`).

## Requirements

- Go 1.16 or later
- Tar (command-line utility)

## Installation

To install, run the following command:

```bash
go get github.com/username/unitypkgextractor
```

## Usage
To extract files from a Unity package, run the following command:

```bash
unitypkgextractor path/to/package.unitypackage [output/directory]
```
Replace path/to/package.unitypackage with the actual path to the Unity package file, and output/directory with the desired output directory. If no output directory is specified, the current working directory is used.


## Example

```bash
unitypkgextractor ~/Downloads/mypkg.unitypackage ~/Documents/MyGame
```

This will extract the contents of mypkg.unitypackage to the directory ~/Documents/mypkg/Assets.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.