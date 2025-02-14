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

package v1beta1

import (
	v1beta1 "github.com/kubescape/storage/pkg/apis/softwarecomposition/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SBOMSummaryLister helps list SBOMSummaries.
// All objects returned here must be treated as read-only.
type SBOMSummaryLister interface {
	// List lists all SBOMSummaries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.SBOMSummary, err error)
	// SBOMSummaries returns an object that can list and get SBOMSummaries.
	SBOMSummaries(namespace string) SBOMSummaryNamespaceLister
	SBOMSummaryListerExpansion
}

// sBOMSummaryLister implements the SBOMSummaryLister interface.
type sBOMSummaryLister struct {
	indexer cache.Indexer
}

// NewSBOMSummaryLister returns a new SBOMSummaryLister.
func NewSBOMSummaryLister(indexer cache.Indexer) SBOMSummaryLister {
	return &sBOMSummaryLister{indexer: indexer}
}

// List lists all SBOMSummaries in the indexer.
func (s *sBOMSummaryLister) List(selector labels.Selector) (ret []*v1beta1.SBOMSummary, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.SBOMSummary))
	})
	return ret, err
}

// SBOMSummaries returns an object that can list and get SBOMSummaries.
func (s *sBOMSummaryLister) SBOMSummaries(namespace string) SBOMSummaryNamespaceLister {
	return sBOMSummaryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SBOMSummaryNamespaceLister helps list and get SBOMSummaries.
// All objects returned here must be treated as read-only.
type SBOMSummaryNamespaceLister interface {
	// List lists all SBOMSummaries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.SBOMSummary, err error)
	// Get retrieves the SBOMSummary from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.SBOMSummary, error)
	SBOMSummaryNamespaceListerExpansion
}

// sBOMSummaryNamespaceLister implements the SBOMSummaryNamespaceLister
// interface.
type sBOMSummaryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SBOMSummaries in the indexer for a given namespace.
func (s sBOMSummaryNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.SBOMSummary, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.SBOMSummary))
	})
	return ret, err
}

// Get retrieves the SBOMSummary from the indexer for a given namespace and name.
func (s sBOMSummaryNamespaceLister) Get(name string) (*v1beta1.SBOMSummary, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("sbomsummary"), name)
	}
	return obj.(*v1beta1.SBOMSummary), nil
}
