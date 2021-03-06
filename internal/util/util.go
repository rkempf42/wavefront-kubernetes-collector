// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"time"

	kube_api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/wait"
	kube_client "k8s.io/client-go/kubernetes"
	v1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

func GetNodeLister(kubeClient *kube_client.Clientset) (v1listers.NodeLister, *cache.Reflector, error) {
	lw := cache.NewListWatchFromClient(kubeClient.CoreV1().RESTClient(), "nodes", kube_api.NamespaceAll, fields.Everything())
	store := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
	nodeLister := v1listers.NewNodeLister(store)
	reflector := cache.NewReflector(lw, &kube_api.Node{}, store, time.Hour)
	go reflector.Run(wait.NeverStop)
	return nodeLister, reflector, nil
}

func GetPodLister(kubeClient *kube_client.Clientset) (v1listers.PodLister, error) {
	lw := cache.NewListWatchFromClient(kubeClient.CoreV1().RESTClient(), "pods", kube_api.NamespaceAll, fields.Everything())
	store := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
	podLister := v1listers.NewPodLister(store)
	reflector := cache.NewReflector(lw, &kube_api.Pod{}, store, time.Hour)
	go reflector.Run(wait.NeverStop)
	return podLister, nil
}

func GetServiceLister(kubeClient *kube_client.Clientset) (v1listers.ServiceLister, error) {
	lw := cache.NewListWatchFromClient(kubeClient.CoreV1().RESTClient(), "services", kube_api.NamespaceAll, fields.Everything())
	store := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
	serviceLister := v1listers.NewServiceLister(store)
	reflector := cache.NewReflector(lw, &kube_api.Service{}, store, time.Hour)
	go reflector.Run(wait.NeverStop)
	return serviceLister, nil
}
