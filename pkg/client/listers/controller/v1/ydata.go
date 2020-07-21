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

package v1

import (
	v1 "github.com/guille-mas/kubernetes-custom-resource/pkg/apis/controller/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// YDataLister helps list YDatas.
// All objects returned here must be treated as read-only.
type YDataLister interface {
	// List lists all YDatas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.YData, err error)
	// YDatas returns an object that can list and get YDatas.
	YDatas(namespace string) YDataNamespaceLister
	YDataListerExpansion
}

// yDataLister implements the YDataLister interface.
type yDataLister struct {
	indexer cache.Indexer
}

// NewYDataLister returns a new YDataLister.
func NewYDataLister(indexer cache.Indexer) YDataLister {
	return &yDataLister{indexer: indexer}
}

// List lists all YDatas in the indexer.
func (s *yDataLister) List(selector labels.Selector) (ret []*v1.YData, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.YData))
	})
	return ret, err
}

// YDatas returns an object that can list and get YDatas.
func (s *yDataLister) YDatas(namespace string) YDataNamespaceLister {
	return yDataNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// YDataNamespaceLister helps list and get YDatas.
// All objects returned here must be treated as read-only.
type YDataNamespaceLister interface {
	// List lists all YDatas in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.YData, err error)
	// Get retrieves the YData from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.YData, error)
	YDataNamespaceListerExpansion
}

// yDataNamespaceLister implements the YDataNamespaceLister
// interface.
type yDataNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all YDatas in the indexer for a given namespace.
func (s yDataNamespaceLister) List(selector labels.Selector) (ret []*v1.YData, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.YData))
	})
	return ret, err
}

// Get retrieves the YData from the indexer for a given namespace and name.
func (s yDataNamespaceLister) Get(name string) (*v1.YData, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("ydata"), name)
	}
	return obj.(*v1.YData), nil
}
