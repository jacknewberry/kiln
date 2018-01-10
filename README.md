# Kiln

Bakes tiles.

## Subcommands

### `bake`

The `bake` command will "bake" a tile. It takes release and stemcell tarballs,
metadata YAML, and JavaScript migrations as inputs and produces an
OpsMan-compatible tile as its output.

Here is an example command line:
```
$ kiln bake \
    --version 2.0.0 \
    --metadata /path/to/metadata.yml \
    --releases-directory /path/to/releases \
    --stemcell-tarball /path/to/stemcell.tgz \
    --migrations-directory /path/to/migrations \
    --output-file /path/to/cf-2.0.0-build.4.pivotal
```

Refer to the [example-tile](example-tile) for a complete example showing the
different features kiln supports.

#### Options

##### `--bosh-variables-directory`

The `--bosh-variables-directory` flag can be used to include CredHub variable
declarations. You should prefer the use of variables rather than Ops Manager
secrets. Each `.yml` file in the directory should define a top-level `variables`
key.

This flag can be specified multiple times if you have organized your
variables into subdirectories for development convenience.

Example [variables](example-tile/variables) directory.

Note that currently you do not use a template helper function to include a
variable, all variables in the directories specified will be included in the
metadata.

##### `--embed`

The `--embed` flag is for embedding any extra files or directories into the
tile. There are very few reasons a tile developer should want to do this, but if
you do, you can include these extra files here. The flag can be specified
multiple times to embed multiple files or directories.

##### `--metadata`

Specify a file path to a tile metadata file for the `--metadata` flag. This
metadata file will contain the contents of your tile configuration as specified
in the OpsManager tile development documentation.

##### `--migrations-directory`

If your tile has JavaScript migrations, then you will need to include the
`--migrations-directory` flag. This flag can be specified multiple times if you
have organized your migrations into subdirectories for development convenience.

##### `--output-file`

The `--output-file` flag takes a path to the location on the filesystem where
your tile will be created. The flag expects a full file name like
`tiles/my-tile-1.2.3-build.4.pivotal`.

##### `--properties-directory`

The `--properties-directory` flag takes a path to a directory that contains one
or more blueprint property files. The flag can also be specified more than once.

To reference a properties file in the directory you can use the `property`
template helper:

```
$ cat /path/to/metadata
---
property_blueprints:
- $( property "rep_password" )
```

Example [properties](example-tile/properties) directory.

##### `--releases-directory`

The `--releases-directory` flag takes a path to a directory that contains one or
many release tarballs. The flag can also be specified more than once. This is
useful if you consume bosh releases as Concourse resources. Each release will
likely show up in the task as a separate directory. For example:

```
$ tree /path/to/releases
|-- first
|   |-- cflinuxfs2-release-1.166.0.tgz
|   `-- consul-release-190.tgz
`-- second
    `-- nats-release-22.tgz
```

To reference a release you can use the `release` template helper:

```
$ cat /path/to/metadata
---
releases:
- $( release "cflinuxfs2" )
- $( release "consul" )
- $( release "nats" )
```

Example kiln command line:

```
$ kiln bake \
    --version 2.0.0 \
    --metadata /path/to/metadata.yml \
    --releases-directory /path/to/releases/first \
    --releases-directory /path/to/releases/second \
    --stemcell-tarball /path/to/stemcell.tgz \
    --output-file /path/to/cf-2.0.0-build.4.pivotal
```

##### `--runtime-configs-directory`

The `--runtime-configs-directory` flag takes a path to a directory that
contains one or more runtime config files. The flag can also be specified
more than once.

To reference a runtime config in the directory you can use the `runtime_config`
template helper:

```
$ cat /path/to/metadata
---
runtime_configs:
- $( runtime_config "first-runtime-config" )
```

Example [runtime-configs](example-tile/runtime-configs) directory.

##### `--stemcell-tarball`

The `--stemcell-tarball` flag takes a path to a stemcell. That stemcell will be
inspected to specify the version and operating system criteria in your tile
metadata.

##### `--stub-releases`

For tile developers looking to get some quick feedback about their tile
metadata, the `--stub-releases` flag will skip including the release tarballs
into the built tile output. This should result in a much smaller file that
should upload much more quickly to OpsManager.

##### `--version`

The `--version` flag takes the version number you want your tile to become. This
version number will show up in the OpsManager UI and will be the version that
your tile "provides" under the `provides_product_versions` metadata key.
