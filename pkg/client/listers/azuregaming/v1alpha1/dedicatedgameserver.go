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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/dgkanatsios/azuregameserversscalingkubernetes/pkg/apis/azuregaming/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DedicatedGameServerLister helps list DedicatedGameServers.
type DedicatedGameServerLister interface {
	// List lists all DedicatedGameServers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DedicatedGameServer, err error)
	// DedicatedGameServers returns an object that can list and get DedicatedGameServers.
	DedicatedGameServers(namespace string) DedicatedGameServerNamespaceLister
	DedicatedGameServerListerExpansion
}

// dedicatedGameServerLister implements the DedicatedGameServerLister interface.
type dedicatedGameServerLister struct {
	indexer cache.Indexer
}

// NewDedicatedGameServerLister returns a new DedicatedGameServerLister.
func NewDedicatedGameServerLister(indexer cache.Indexer) DedicatedGameServerLister {
	return &dedicatedGameServerLister{indexer: indexer}
}

// List lists all DedicatedGameServers in the indexer.
func (s *dedicatedGameServerLister) List(selector labels.Selector) (ret []*v1alpha1.DedicatedGameServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DedicatedGameServer))
	})
	return ret, err
}

// DedicatedGameServers returns an object that can list and get DedicatedGameServers.
func (s *dedicatedGameServerLister) DedicatedGameServers(namespace string) DedicatedGameServerNamespaceLister {
	return dedicatedGameServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DedicatedGameServerNamespaceLister helps list and get DedicatedGameServers.
type DedicatedGameServerNamespaceLister interface {
	// List lists all DedicatedGameServers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DedicatedGameServer, err error)
	// Get retrieves the DedicatedGameServer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DedicatedGameServer, error)
	DedicatedGameServerNamespaceListerExpansion
}

// dedicatedGameServerNamespaceLister implements the DedicatedGameServerNamespaceLister
// interface.
type dedicatedGameServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DedicatedGameServers in the indexer for a given namespace.
func (s dedicatedGameServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DedicatedGameServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DedicatedGameServer))
	})
	return ret, err
}

// Get retrieves the DedicatedGameServer from the indexer for a given namespace and name.
func (s dedicatedGameServerNamespaceLister) Get(name string) (*v1alpha1.DedicatedGameServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dedicatedgameserver"), name)
	}
	return obj.(*v1alpha1.DedicatedGameServer), nil
}
