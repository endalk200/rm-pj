# rm-rf-js-junk

When working on multiple projects there will be forgotten files and directories taking too much space that are not needed anymore. rm-rm-js-junk is a simple go CLI app that walks specified directory and remove all js junk files like `node_modules`, `bower_components`, `dist`, `build` and etc.

## Usage

Remove all js junk files from specified directory

```bash
rm-rf -p /path/to/project
```

Remove all js junk files from current directory

```bash
rm-rf
```
