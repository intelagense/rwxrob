# Kubernetes Notes and Cheatsheet

Kubernetes is mandatory learning for everything I do for work and play.
I might get the CKAD and CAD eventually, but having the skills is more
important than the paper.

## Python API Documentation

<https://github.com/kubernetes-client/python/tree/master/kubernetes/docs>

## Essential Elements of a Home Lab

1. Linux system with 16+ cores
1. Kind
1. Local Registry (probably in Docker, port 5000)

## Later Elements

1. Load Balancer
1. Block Device Volume Mounts

## Stuff to Learn

I need to learn the following in order of priority:

1. Docker (including Dockerfiles)
1. Kubernetes (1.18 for now)
1. Seldon.io - AI models at scale in production
1. Prometheus - metrics collector in Go
1. ElasticSearch - just a log search engine

## Getting Started

1. `minikube` - fast setup but runs in a vm so not ideal
1. `microk8s` - requires snaps (less than ideal) but real
1. `kind` - full local k8s

## Re: Podman

Just say *HELL no* --- especially as an "alternative" to Kubernetes.

## Documentation

Apparently you can use this document on the exam itself, so this is
*mandatory* learning if for nothing more than the structure of where
stuff is so I can find it quickly:

* <https://kubernetes.io/docs/home/>

Other resources:

* <https://github.com/kelseyhightower/envconfig>
* <https://github.com/cncf/sig-security/blob/master/security-whitepaper/cloud-native-security-whitepaper.md>
* k8s slack
* <https://operatorhub.io>
* <https://github.com/NVIDIA/k8s-device-plugin> (access to GPU)
* <https://www.envoyproxy.io/docs/envoy/latest/start/start>
* <https://layer5.io/service-mesh-landscape>
* <https://github.com/dgkanatsios/CKAD-exercises>
* <https://kubernetes.io/docs/reference/kubectl/cheatsheet/>
* <https://kube.academy/lessons/training-and-preparation-resources>
* <https://github.com/cncf/curriculum/blob/master/CKA_Curriculum_v1.20.pdf>
* <https://kube.academy/>
* <https://github.com/kelseyhightower/kubernetes-the-hard-way/projects>
* <https://k3d.io/>
* <https://github.com/kubernetes-sigs/kind>
* <https://www.alexellis.io/>
* [qmacro videos](https://www.youtube.com/playlist?list=PLfctWmgNyOIf9rXaZp9RSM2YVxAPGGthe)

## Task Cheatsheet

|Task|Command|
|-|-|
Watch everything while changing|`k get pods -A -w`
