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
		"github.com/apicurito-operator/pkg/apis/apicuritos/v1alpha1.Apicurito":       schema_pkg_apis_apicuritos_v1alpha1_Apicurito(ref),
		"github.com/apicurito-operator/pkg/apis/apicuritos/v1alpha1.ApicuritoSpec":   schema_pkg_apis_apicuritos_v1alpha1_ApicuritoSpec(ref),
		"github.com/apicurito-operator/pkg/apis/apicuritos/v1alpha1.ApicuritoStatus": schema_pkg_apis_apicuritos_v1alpha1_ApicuritoStatus(ref),
	}
}

func schema_pkg_apis_apicuritos_v1alpha1_Apicurito(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Apicurito is the Schema for the apicuritos API",
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
							Ref: ref("github.com/apicurito-operator/pkg/apis/apicuritos/v1alpha1.ApicuritoSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/apicurito-operator/pkg/apis/apicuritos/v1alpha1.ApicuritoStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/apicurito-operator/pkg/apis/apicuritos/v1alpha1.ApicuritoSpec", "github.com/apicurito-operator/pkg/apis/apicuritos/v1alpha1.ApicuritoStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_apicuritos_v1alpha1_ApicuritoSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApicuritoSpec defines the desired state of Apicurito",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_apicuritos_v1alpha1_ApicuritoStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ApicuritoStatus defines the observed state of Apicurito",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
