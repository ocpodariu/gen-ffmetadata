## gen-ffmetadata

This tool allows you to write your video metadata as a YAML document and then generate an [FFmpeg Metadata file](https://ffmpeg.org/ffmpeg-formats.html#Metadata-1) from it.

### Workflow example
Let's assume you want to add chapter markers to your video `awesome_compilation.mp4`. You can do it like this:

1. Write your video metadata and chapter details in `compilation.yaml`, following the format in `sample/video.yaml`.
2. Generate the FFmpeg Metadata with this tool:
```
gen-ffmetadata -templates/mp4.tpl compilation.yaml
```
3. Map the FFmpeg Metadata file `compilation.metadata` to your video:
```
ffmpeg -i awesome_compilation.mp4 -i compilation.metadata -map_metadata 1 -map_chapters 1 -codec copy compilation_with_metadata.mp4
```

### Supported video formats
For now it only outputs metadata for Quicktime/MP4 videos with the `templates/mp4.tpl` template.

If you want to expand it for other video formats, you need to create a new template in the `templates/` directory. These links will help you write a template for any video format supported by FFmpeg:

- [List of metadata keys accepted by FFmpeg](https://wiki.multimedia.cx/index.php?title=FFmpeg_Metadata)
- [Golang text/template](https://golang.org/pkg/text/template/)
- [Golang Template Cheatsheet](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet)


### How to install
Get [Git](https://git-scm.com/) and [Go](https://golang.org/doc/install) in case you don't have them already, clone the repository and build the app.
```
git clone https://github.com/ocpodariu/gen-ffmetadata.git
cd gen-ffmetadata
go build
```

### How to use
```
./gen-ffmetadata [OPTIONS] METADATA_FILE
```

#### Flags
`-out` (default: same name as the input filename, but with `.metadata` extension) - Change the name of the output FFmpeg Metadata file:
```
./gen-ffmetadata -out video14.meta METADATA_FILE
```

`-template` (default: "metadata.tpl") - Use a custom template for the output FFmpeg Metadata file:
```
./gen-ffmetadata -template mkv.tpl METADATA_FILE
```
