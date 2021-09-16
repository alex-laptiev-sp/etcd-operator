/*
Copyright 2019 The etcd-operator Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta2

import (
	time "time"

	etcdv1beta2 "github.com/alex-laptiev-sp/etcd-operator/pkg/apis/etcd/v1beta2"
	versioned "github.com/alex-laptiev-sp/etcd-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/alex-laptiev-sp/etcd-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta2 "github.com/alex-laptiev-sp/etcd-operator/pkg/generated/listers/etcd/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EtcdClusterInformer provides access to a shared informer and lister for
// EtcdClusters.
type EtcdClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta2.EtcdClusterLister
}

type etcdClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEtcdClusterInformer constructs a new informer for EtcdCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEtcdClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEtcdClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEtcdClusterInformer constructs a new informer for EtcdCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEtcdClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EtcdV1beta2().EtcdClusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EtcdV1beta2().EtcdClusters(namespace).Watch(options)
			},
		},
		&etcdv1beta2.EtcdCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *etcdClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEtcdClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *etcdClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&etcdv1beta2.EtcdCluster{}, f.defaultInformer)
}

func (f *etcdClusterInformer) Lister() v1beta2.EtcdClusterLister {
	return v1beta2.NewEtcdClusterLister(f.Informer().GetIndexer())
}
