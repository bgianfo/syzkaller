# SerenityOS Support

[SerenityOS](https://github.com/SerenityOS/serenity/) support is *incomplete*.

SerenityOS does not support Go at the moment.
Until that happens running on SerenityOS is challening.

However, `syz-stress` can be run as follows:

```shell
make TARGETOS=linux stress
make TARGETOS=serenityos TARGETARCH=386 SOURCEDIR=~/src/serenity executor
scp -P 5555 -i serenityos_id_rsa -o IdentitiesOnly=yes bin/serenityos_386/syz-executor  root@localhost:/
bin/linux_amd64/syz-stress -os=serenityos -ipc=pipe -procs=8 -executor "/usr/bin/ssh -p 5555 -i serenityos_id_rsa -o IdentitiesOnly=yes root@localhost /syz-executor"
```
