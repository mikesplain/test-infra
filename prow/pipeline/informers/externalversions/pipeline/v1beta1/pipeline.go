/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	time "time"

	pipelinev1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "k8s.io/test-infra/prow/pipeline/clientset/versioned"
	internalinterfaces "k8s.io/test-infra/prow/pipeline/informers/externalversions/internalinterfaces"
	v1beta1 "k8s.io/test-infra/prow/pipeline/listers/pipeline/v1beta1"
)

// PipelineInformer provides access to a shared informer and lister for
// Pipelines.
type PipelineInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.PipelineLister
}

type pipelineInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPipelineInformer constructs a new informer for Pipeline type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPipelineInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPipelineInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPipelineInformer constructs a new informer for Pipeline type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPipelineInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TektonV1beta1().Pipelines(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TektonV1beta1().Pipelines(namespace).Watch(options)
			},
		},
		&pipelinev1beta1.Pipeline{},
		resyncPeriod,
		indexers,
	)
}

func (f *pipelineInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPipelineInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pipelineInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&pipelinev1beta1.Pipeline{}, f.defaultInformer)
}

func (f *pipelineInformer) Lister() v1beta1.PipelineLister {
	return v1beta1.NewPipelineLister(f.Informer().GetIndexer())
}
