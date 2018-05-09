# Kubernetes HeatScheduler
Nerdalize is building a different cloud. Rather than putting thousands of servers into a datacenter, we're distributing them over homes. There, homeowners make good use of the residual heat: to heat their home in winter and their shower water in summer.

The HeatScheduler is a scheduler for Kubernetes that takes boiler temperatures into account. Most important is that it shows how to run an extra scheduler with your own scheduling logic on any Kubernetes cluster.

## How can I deploy this?
1. Change you Docker image name in the `Makefile`.
2. Build and push the Docker image to a container registry: `make push`.
3. Deploy the policy config: `kubectl create -f deployments/scheduler-policy-config.yaml`.
4. Deploy the scheduler: `kubectl create -f deployments/heat-scheduler.yaml`.
5. Deploy a pod that uses `schedulerName: heat-scheduler`. See `deployments/pod-example.yaml` for an example.

## How can I customize this?
Feel free to use this code for your own custom Kubernetes scheduler. Including your own business logic can be done by changing the `selectNode` function in `util.go`.

## How does it work?
* In the `deployments` folder you find a Kubernetes Deployment that launches a new [kube-scheduler](https://kubernetes.io/docs/reference/generated/kube-scheduler/). More information about this can be found in the [Kubernetes docs](https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/).
* In `deployments/scheduler-policy-config.yaml` you can change the predicates that the scheduler should use by default, before passing the pod to the extender. A list of these predicates can be found [here](https://github.com/kubernetes/kubernetes/blob/master/pkg/scheduler/algorithm/predicates/predicates.go#L50).
* We currently only use the `filter` endpoint, which returns a filtered subset of the original list of Nodes. It is also possible to return a weighted list of Nodes that can be used to indicate Node priorities by using the `prioritize` endpoint. See the official [docs](https://github.com/kubernetes/kubernetes/blob/release-1.5/docs/design/scheduler_extender.md) for more information.
