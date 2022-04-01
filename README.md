# Avoid Palindrome Pod Labels Policy 

This repository contains a Kubewarden policy written in Go to reject Pods with at least one label wiyh palindrome key.

More information about Kubewarden and its policies can be found the [official documentation](https://docs.kubewarden.io)

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
