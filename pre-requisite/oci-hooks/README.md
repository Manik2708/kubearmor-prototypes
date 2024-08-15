# LFX PRE-REQUISITE FOR Support Podman for unorchestrated environments

## Video Link
I have created a video regarding the important finding of this task, please see [this](https://www.youtube.com/watch?v=e2RwIX3o2ik&t=3s)

## Steps to follow to apply this Hook

1. Create this hook and paste into `/usr/share/containers/oci/hooks.d`
```text
{
 "version": "1.0.0",
 "hook": {
"path": "/path/to/pre-requisite-oci"
 },
 "when": {
     "always": true
 },
 "stages": ["precreate"]
}
```
Here `/path/to/pre-requisite-oci` is the path for the script present in this directory

2. Make sure that hooks directory is embedded into `usr/share/containers/containers.conf`
```text
hooks_dir = [
  "/usr/share/containers/oci/hooks.d"]
```
3. Change the profile name in `main.go` to an already loaded app-armor profile name.

4. Run `go build`

5. Ensure that the script present here (pre-requisite-oci) has enough permissions. To confirm this, run this in this directory
```text
sudo chmod +x pre-requisite-oci
```
6. Confirm the profile being loaded by hook by running this command
```text
~/sudo podman inspect <container_id> | grep -i appArmor
```

7. Now run the podman container and try to violate the policy.
```text
~/sudo podman run -it nginx #Run a policy violating command
```

## An important Finding

I found that OCI hooks are executing the policy only on some of the PIDs of the container. For example run this command for a running container (after hook is applied) and check whether policy is violated
```text
~/sudo podman exec -it <container_id> /bin/bash
> Run a command that violates the policy
``` 
You will observe that, policy will not work this time. So you would ask why it worked before? Because we ran the command on PID:1. Try to run these commands 
```text
~/sudo podman exec -it <container_id> /bin/bash
> cat /proc/1/attr/current #Observe the result
> cat /proc/self/attr/current #Again observe the result
```
You will observe that in `proc/1/attr/current`, the new profile is loaded but in the second case, the default app-armor profile is loaded. This might be an issue from podman's or open-container's side. I have raised the issue regarding [this](https://github.com/containers/podman/issues/23589)




