// This file was automatically generated by informer-gen

package v1

import (
	user_v1 "github.com/openshift/origin/pkg/user/apis/user/v1"
	clientset "github.com/openshift/origin/pkg/user/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/user/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/user/generated/listers/user/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// IdentityInformer provides access to a shared informer and lister for
// Identities.
type IdentityInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.IdentityLister
}

type identityInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newIdentityInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.UserV1().Identities().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.UserV1().Identities().Watch(options)
			},
		},
		&user_v1.Identity{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *identityInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&user_v1.Identity{}, newIdentityInformer)
}

func (f *identityInformer) Lister() v1.IdentityLister {
	return v1.NewIdentityLister(f.Informer().GetIndexer())
}
