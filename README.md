# Unity Package Extractor

This is a command-line tool for extracting files from Unity package files (`.unitypackage`).

This took took inspiration from https://github.com/Cobertos/unitypackage_extractor. All I did was rewrite it into Go and further optimize it. The original work deserves credit, however this implementation is unquestionably faster.


## TODO
- Add a GUI using either Fyne, Wails, or Gio

## Requirements

- Go 1.16 or later
- Tar (command-line utility)

## Installation

### From releases
Navigate to the releases, and download our latest pre-built executable file for your operating system.

### From source
To install Unity Package Extractor, please follow these steps:

Clone the repository to your local machine using the following command:
```bash
git clone https://github.com/Minagoroshi/unitypkgextractor.git
````
Navigate to the project directory:
```bash 
cd unitypkgextractor
```
Build the project by running the following command:
```bash
go build .
````
You can now use Unity Package Extractor by running the built executable file.
Note: Make sure you have Go installed on your machine before proceeding with the installation. If you do not have Go installed, you can download it from the official website: https://golang.org/dl/

## Usage
To extract files from a Unity package, either:

1. Drag and Drop the Unity package file onto the executable file


2. Run the following command:

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
