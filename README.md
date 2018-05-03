# Kubernetes HeatScheduler
Nerdalize is building a different cloud. Rather than putting thousands of servers into a datacenter, we're distributing them over homes. There, homeowners make good use of the residual heat: to heat their home in winter and their shower water in summer.

The HeatScheduler is a scheduler for Kubernetes that takes boiler temperatures into account. Most important is that it shows how to run an extra scheduler with your own scheduling logic on any Kubernetes cluster.

## How does it work?
* In the `deployments` folder you find a Kubernetes Deployment that launches a new [kube-scheduler](https://kubernetes.io/docs/reference/generated/kube-scheduler/). More information about this can be found in the [Kubernetes docs](https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/).
* In `deployments/scheduler-policy-config.yaml` you can change the predicates that the scheduler should use by default, before passing the pod to the extender. A list of these predicates can be found [here](https://github.com/kubernetes/kubernetes/blob/master/pkg/scheduler/algorithm/predicates/predicates.go#L50).
* The custom scheduling logic is defined in `handler.go`.

## How can I deploy this?
1. Change you Docker image name in the `Makefile`.
2. `make push`
3. `kubectl create -f deployments/scheduler-policy-config.yaml`
4. `kubectl create -f deployments/heat-scheduler.yaml`
5. Deploy a pod that uses `schedulerName: heat-scheduler`. See `deployments/pod-example.yaml` for an example.
