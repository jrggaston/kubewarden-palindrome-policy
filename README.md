# Avoid Palindrome Pod Labels Policy 

This repository contains a Kubewarden policy written in Go to reject Pods with at least one label wiyh palindrome key.

More information about Kubewarden and its policies can be found in the [official documentation](https://docs.kubewarden.io)

The policy allows the creation of this Pod:
```
apiVersion: v1
kind: Pod
metadata:
  name: hello-world
  labels:
    env: production
spec:
  containers:
  - name: nginx
    image: nginx
```

and rejects the creation of this Pod:
```
apiVersion: v1
kind: Pod
metadata:
  name: hello-world
  labels:
    env: production
    level: debug
spec:
  containers:
  - name: nginx
    image: nginx
```

Requirements:
- Keys of one unique character shall not be consider palindrome
- If the key has prefix (<prefix>/<name>), both the key prefix and the key name shall be considered
- The policy shall be case insensitive


