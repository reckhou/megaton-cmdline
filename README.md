megaton-cmdline
===============

Command line tool of megaton.

##Tools & Usage:

Tools are under `tools` folder. For detailed usage, run each tool without any parameters.

###DatExtractor

`datExtractor` is used to extract all files packed inside a `dat` file.

    ./datExtractor <datFilePath> <options...>
    
###PatchGenerator

`patchGenerator` is used to compare differences bewteen 2 `dat` files and generate a new patch file.

`newVersionCode` indecates patch file's version code. The patch file's version code MUST be higher than base `dat` file's to make sure it can be applied correctly by client.

    ./patchGen <datWithLowVersion> <datWithHighVersion> <patchFilePath> <newVersionCode> <option>...
