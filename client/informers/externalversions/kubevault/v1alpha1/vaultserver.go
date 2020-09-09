/*
Copyright The KubeVault Authors.

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

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	kubevaultv1alpha1 "kubevault.dev/operator/apis/kubevault/v1alpha1"
	versioned "kubevault.dev/operator/client/clientset/versioned"
	internalinterfaces "kubevault.dev/operator/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubevault.dev/operator/client/listers/kubevault/v1alpha1"
)

// VaultServerInformer provides access to a shared informer and lister for
// VaultServers.
type VaultServerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VaultServerLister
}

type vaultServerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVaultServerInformer constructs a new informer for VaultServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVaultServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVaultServerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVaultServerInformer constructs a new informer for VaultServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVaultServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubevaultV1alpha1().VaultServers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubevaultV1alpha1().VaultServers(namespace).Watch(context.TODO(), options)
			},
		},
		&kubevaultv1alpha1.VaultServer{},
		resyncPeriod,
		indexers,
	)
}

func (f *vaultServerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVaultServerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vaultServerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubevaultv1alpha1.VaultServer{}, f.defaultInformer)
}

func (f *vaultServerInformer) Lister() v1alpha1.VaultServerLister {
	return v1alpha1.NewVaultServerLister(f.Informer().GetIndexer())
}
