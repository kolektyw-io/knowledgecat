# KnowledgeCat Files structure

## Data folder
KnowledgeCat consists of data folder, which contains sub-folders called namespaces that allow documentation to
be split into respective topics.

### Basic structure of data folder

```
   <<datadir>>
        |
        ----- NAMESPACE1
        |          |
        |          ----namespace.json
        |          ----0001_Intro.md
        |
        ------ NAMESPACE2
                    |
                    ----- namespace.json
                    ----- 0002_Some_another_file.md
```

To be recognized, namespace has to have in its directory `namespace.json` file which sets order of topics shown on left pane.

### Structure of namespace.json

```
    {
  "description" : "Default KnowledgeCat Server namespace",
  "name" : "KnowledgeCat Server",
  "mnemonic" : "KCATSRV",
  "entries" : [
    {
      "file" : "0001_Intro.md",
      "title": "Home",
      "children" : []
    },
    {
      "file" : "0002_Installation.md",
      "title" : "Installation",
      "children": []
    },
    {
      "file" : "0003_File_structure.md",
      "title" : "KnowledgeCat helpfiles structure",
      "children" : []
    }
  ]
}
```


- `descriprion` contains descriptive text regarding namespace 
- `name` contains text shown as namespace name 
- `mnemonic` is used to indicate folder location (i.e. when namespace needs to be changed because it overlaps with another its mnemonic shall be altered) 
- `entries` array of dicts `{"file":file location, "title": title of document, "children": array of subarticles - not yet used}`   