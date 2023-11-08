// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package bkgql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// GetPipelinePipeline includes the requested fields of the GraphQL type Pipeline.
// The GraphQL type's documentation follows.
//
// A pipeline
type GetPipelinePipeline struct {
	Id string `json:"id"`
	// The repository for this pipeline
	Repository GetPipelinePipelineRepository `json:"repository"`
	Steps      GetPipelinePipelineSteps      `json:"steps"`
}

// GetId returns GetPipelinePipeline.Id, and is useful for accessing the field via an interface.
func (v *GetPipelinePipeline) GetId() string { return v.Id }

// GetRepository returns GetPipelinePipeline.Repository, and is useful for accessing the field via an interface.
func (v *GetPipelinePipeline) GetRepository() GetPipelinePipelineRepository { return v.Repository }

// GetSteps returns GetPipelinePipeline.Steps, and is useful for accessing the field via an interface.
func (v *GetPipelinePipeline) GetSteps() GetPipelinePipelineSteps { return v.Steps }

// GetPipelinePipelineRepository includes the requested fields of the GraphQL type Repository.
// The GraphQL type's documentation follows.
//
// A repository associated with a pipeline
type GetPipelinePipelineRepository struct {
	// The git URL for this repository
	Url string `json:"url"`
}

// GetUrl returns GetPipelinePipelineRepository.Url, and is useful for accessing the field via an interface.
func (v *GetPipelinePipelineRepository) GetUrl() string { return v.Url }

// GetPipelinePipelineSteps includes the requested fields of the GraphQL type PipelineSteps.
// The GraphQL type's documentation follows.
//
// Steps defined on a pipeline
type GetPipelinePipelineSteps struct {
	// A YAML representation of the pipeline steps
	Yaml string `json:"yaml"`
}

// GetYaml returns GetPipelinePipelineSteps.Yaml, and is useful for accessing the field via an interface.
func (v *GetPipelinePipelineSteps) GetYaml() string { return v.Yaml }

// GetPipelineResponse is returned by GetPipeline on success.
type GetPipelineResponse struct {
	// Find a pipeline
	Pipeline GetPipelinePipeline `json:"pipeline"`
}

// GetPipeline returns GetPipelineResponse.Pipeline, and is useful for accessing the field via an interface.
func (v *GetPipelineResponse) GetPipeline() GetPipelinePipeline { return v.Pipeline }

// UpdatePipelinePipelineUpdatePipelineUpdatePayload includes the requested fields of the GraphQL type PipelineUpdatePayload.
// The GraphQL type's documentation follows.
//
// Autogenerated return type of PipelineUpdate.
type UpdatePipelinePipelineUpdatePipelineUpdatePayload struct {
	// A unique identifier for the client performing the mutation.
	ClientMutationId string                                                    `json:"clientMutationId"`
	Pipeline         UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipeline `json:"pipeline"`
}

// GetClientMutationId returns UpdatePipelinePipelineUpdatePipelineUpdatePayload.ClientMutationId, and is useful for accessing the field via an interface.
func (v *UpdatePipelinePipelineUpdatePipelineUpdatePayload) GetClientMutationId() string {
	return v.ClientMutationId
}

// GetPipeline returns UpdatePipelinePipelineUpdatePipelineUpdatePayload.Pipeline, and is useful for accessing the field via an interface.
func (v *UpdatePipelinePipelineUpdatePipelineUpdatePayload) GetPipeline() UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipeline {
	return v.Pipeline
}

// UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipeline includes the requested fields of the GraphQL type Pipeline.
// The GraphQL type's documentation follows.
//
// A pipeline
type UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipeline struct {
	Steps UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipelineSteps `json:"steps"`
}

// GetSteps returns UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipeline.Steps, and is useful for accessing the field via an interface.
func (v *UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipeline) GetSteps() UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipelineSteps {
	return v.Steps
}

// UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipelineSteps includes the requested fields of the GraphQL type PipelineSteps.
// The GraphQL type's documentation follows.
//
// Steps defined on a pipeline
type UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipelineSteps struct {
	// A YAML representation of the pipeline steps
	Yaml string `json:"yaml"`
}

// GetYaml returns UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipelineSteps.Yaml, and is useful for accessing the field via an interface.
func (v *UpdatePipelinePipelineUpdatePipelineUpdatePayloadPipelineSteps) GetYaml() string {
	return v.Yaml
}

// UpdatePipelineResponse is returned by UpdatePipeline on success.
type UpdatePipelineResponse struct {
	// Change the settings for a pipeline.
	PipelineUpdate UpdatePipelinePipelineUpdatePipelineUpdatePayload `json:"pipelineUpdate"`
}

// GetPipelineUpdate returns UpdatePipelineResponse.PipelineUpdate, and is useful for accessing the field via an interface.
func (v *UpdatePipelineResponse) GetPipelineUpdate() UpdatePipelinePipelineUpdatePipelineUpdatePayload {
	return v.PipelineUpdate
}

// __GetPipelineInput is used internally by genqlient
type __GetPipelineInput struct {
	OrgPipelineSlug string `json:"orgPipelineSlug"`
}

// GetOrgPipelineSlug returns __GetPipelineInput.OrgPipelineSlug, and is useful for accessing the field via an interface.
func (v *__GetPipelineInput) GetOrgPipelineSlug() string { return v.OrgPipelineSlug }

// __UpdatePipelineInput is used internally by genqlient
type __UpdatePipelineInput struct {
	Id   string `json:"id"`
	Yaml string `json:"yaml"`
}

// GetId returns __UpdatePipelineInput.Id, and is useful for accessing the field via an interface.
func (v *__UpdatePipelineInput) GetId() string { return v.Id }

// GetYaml returns __UpdatePipelineInput.Yaml, and is useful for accessing the field via an interface.
func (v *__UpdatePipelineInput) GetYaml() string { return v.Yaml }

// The query or mutation executed by GetPipeline.
const GetPipeline_Operation = `
query GetPipeline ($orgPipelineSlug: ID!) {
	pipeline(slug: $orgPipelineSlug) {
		id
		repository {
			url
		}
		steps {
			yaml
		}
	}
}
`

func GetPipeline(
	ctx context.Context,
	client graphql.Client,
	orgPipelineSlug string,
) (*GetPipelineResponse, error) {
	req := &graphql.Request{
		OpName: "GetPipeline",
		Query:  GetPipeline_Operation,
		Variables: &__GetPipelineInput{
			OrgPipelineSlug: orgPipelineSlug,
		},
	}
	var err error

	var data GetPipelineResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

// The query or mutation executed by UpdatePipeline.
const UpdatePipeline_Operation = `
mutation UpdatePipeline ($id: ID!, $yaml: String!) {
	pipelineUpdate(input: {id:$id,steps:{yaml:$yaml}}) {
		clientMutationId
		pipeline {
			steps {
				yaml
			}
		}
	}
}
`

func UpdatePipeline(
	ctx context.Context,
	client graphql.Client,
	id string,
	yaml string,
) (*UpdatePipelineResponse, error) {
	req := &graphql.Request{
		OpName: "UpdatePipeline",
		Query:  UpdatePipeline_Operation,
		Variables: &__UpdatePipelineInput{
			Id:   id,
			Yaml: yaml,
		},
	}
	var err error

	var data UpdatePipelineResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
