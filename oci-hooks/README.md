# OCI Hooks Prototype

Paste this hook in `usr/share/containers/oci/hooks.d`, Here `/path/to/oci-hook-binary` is the path for the binary file built from this go code

```text
{
  "version": "1.0.0",
  "hook": {
    "path": "/path/to/oci-hook-binary"
  },
  "when": {
    "always": true
  },
  "stages": ["createRuntime", "poststop"]

}
```

Now run image through podman by root user
```text
~/sudo podman run sudo podman run docker.io/library/nginx
```

The hook will be invoked and json files will be saved. Example json files are present in this repository as `spec.json`, `state.json` and `container.txt`