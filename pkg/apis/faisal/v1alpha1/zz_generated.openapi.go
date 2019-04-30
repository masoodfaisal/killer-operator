// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/example-inc/killer-operator/pkg/apis/faisal/v1alpha1.PodKiller":       schema_pkg_apis_faisal_v1alpha1_PodKiller(ref),
		"github.com/example-inc/killer-operator/pkg/apis/faisal/v1alpha1.PodKillerSpec":   schema_pkg_apis_faisal_v1alpha1_PodKillerSpec(ref),
		"github.com/example-inc/killer-operator/pkg/apis/faisal/v1alpha1.PodKillerStatus": schema_pkg_apis_faisal_v1alpha1_PodKillerStatus(ref),
	}
}

func schema_pkg_apis_faisal_v1alpha1_PodKiller(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodKiller is the Schema for the podkillers API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/example-inc/killer-operator/pkg/apis/faisal/v1alpha1.PodKillerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/example-inc/killer-operator/pkg/apis/faisal/v1alpha1.PodKillerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/example-inc/killer-operator/pkg/apis/faisal/v1alpha1.PodKillerSpec", "github.com/example-inc/killer-operator/pkg/apis/faisal/v1alpha1.PodKillerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_faisal_v1alpha1_PodKillerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodKillerSpec defines the desired state of PodKiller",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_faisal_v1alpha1_PodKillerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodKillerStatus defines the observed state of PodKiller",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}