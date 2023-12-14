// Copyright The Karbour Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/KusionStack/karbour/pkg/core/handler"
	"github.com/KusionStack/karbour/pkg/core/manager/cluster"
	"github.com/KusionStack/karbour/pkg/multicluster"
	"github.com/KusionStack/karbour/pkg/util/ctxutil"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"k8s.io/apiserver/pkg/server"
)

// Get returns an HTTP handler function that reads a cluster
// detail. It utilizes a ClusterManager to execute the logic.
//
//	@Summary		Get returns a cluster resource by name.
//	@Description	This endpoint returns a cluster resource by name.
//	@Tags			cluster
//	@Produce		json
//	@Success		200	{object}	unstructured.Unstructured	"Unstructured object"
//	@Failure		400	{string}	string						"Bad Request"
//	@Failure		401	{string}	string						"Unauthorized"
//	@Failure		404	{string}	string						"Not Found"
//	@Failure		405	{string}	string						"Method Not Allowed"
//	@Failure		429	{string}	string						"Too Many Requests"
//	@Failure		500	{string}	string						"Internal Server Error"
//	@Router			/api/v1/cluster/{clusterName} [get]
func Get(clusterMgr *cluster.ClusterManager, c *server.CompletedConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cluster := chi.URLParam(r, "clusterName")
		// Extract the context and logger from the request.
		ctx := r.Context()
		logger := ctxutil.GetLogger(ctx)
		logger.Info("Getting cluster...")

		client, err := multicluster.BuildMultiClusterClient(r.Context(), c.LoopbackClientConfig, "")
		if err != nil {
			render.Render(w, r, handler.FailureResponse(r.Context(), err))
			return
		}
		clusterUnstructured, err := clusterMgr.GetCluster(r.Context(), client, cluster)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(r.Context(), err))
			return
		}
		render.JSON(w, r, handler.SuccessResponse(ctx, clusterUnstructured))
	}
}

// Create returns an HTTP handler function that creates a cluster
// resource. It utilizes a ClusterManager to execute the logic.
//
//	@Summary		Create creates a cluster resource.
//	@Description	This endpoint creates a new cluster resource using the payload.
//	@Tags			cluster
//	@Accept			plain
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ClusterPayload				true	"cluster to create (either plain text or JSON format)"
//	@Success		200		{object}	unstructured.Unstructured	"Unstructured object"
//	@Failure		400		{string}	string						"Bad Request"
//	@Failure		401		{string}	string						"Unauthorized"
//	@Failure		404		{string}	string						"Not Found"
//	@Failure		405		{string}	string						"Method Not Allowed"
//	@Failure		429		{string}	string						"Too Many Requests"
//	@Failure		500		{string}	string						"Internal Server Error"
//	@Router			/api/v1/cluster [post]
func Create(clusterMgr *cluster.ClusterManager, c *server.CompletedConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the context and logger from the request.
		ctx := r.Context()
		logger := ctxutil.GetLogger(ctx)
		logger.Info("Creating cluster...")

		// Decode the request body into the payload.
		payload := &ClusterPayload{}
		if err := payload.Decode(r); err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}

		client, _ := multicluster.BuildMultiClusterClient(r.Context(), c.LoopbackClientConfig, "")
		clusterCreated, err := clusterMgr.CreateCluster(r.Context(), client, payload.ClusterName, payload.ClusterDisplayName, payload.ClusterDescription, payload.ClusterKubeConfig)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		render.JSON(w, r, handler.SuccessResponse(ctx, clusterCreated))
	}
}

// Create returns an HTTP handler function that updates a cluster
// resource. It utilizes a ClusterManager to execute the logic.
//
//	@Summary		UpdateMetadata updates the cluster metadata by name.
//	@Description	This endpoint updates the display name and description of an existing cluster resource.
//	@Tags			cluster
//	@Accept			plain
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ClusterPayload				true	"cluster to update (either plain text or JSON format)"
//	@Success		200		{object}	unstructured.Unstructured	"Unstructured object"
//	@Failure		400		{string}	string						"Bad Request"
//	@Failure		401		{string}	string						"Unauthorized"
//	@Failure		404		{string}	string						"Not Found"
//	@Failure		405		{string}	string						"Method Not Allowed"
//	@Failure		429		{string}	string						"Too Many Requests"
//	@Failure		500		{string}	string						"Internal Server Error"
//	@Router			/api/v1/cluster/{clusterName}  [put]
func UpdateMetadata(clusterMgr *cluster.ClusterManager, c *server.CompletedConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the context and logger from the request.
		ctx := r.Context()
		logger := ctxutil.GetLogger(ctx)
		logger.Info("Updating cluster metadata...")
		cluster := chi.URLParam(r, "clusterName")

		// Decode the request body into the payload.
		payload := &ClusterPayload{}
		if err := payload.Decode(r); err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}

		client, _ := multicluster.BuildMultiClusterClient(r.Context(), c.LoopbackClientConfig, "")
		clusterUpdated, err := clusterMgr.UpdateMetadata(r.Context(), client, cluster, payload.ClusterDisplayName, payload.ClusterDescription)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		render.JSON(w, r, handler.SuccessResponse(ctx, clusterUpdated))
	}
}

// Delete returns an HTTP handler function that deletes a cluster
// resource. It utilizes a ClusterManager to execute the logic.
//
//	@Summary		Delete removes a cluster resource by name.
//	@Description	This endpoint deletes the cluster resource by name.
//	@Tags			cluster
//	@Produce		json
//	@Success		200	{string}	string	"Operation status"
//	@Failure		400	{string}	string	"Bad Request"
//	@Failure		401	{string}	string	"Unauthorized"
//	@Failure		404	{string}	string	"Not Found"
//	@Failure		405	{string}	string	"Method Not Allowed"
//	@Failure		429	{string}	string	"Too Many Requests"
//	@Failure		500	{string}	string	"Internal Server Error"
//	@Router			/api/v1/cluster/{clusterName} [delete]
func Delete(clusterMgr *cluster.ClusterManager, c *server.CompletedConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the context and logger from the request.
		ctx := r.Context()
		logger := ctxutil.GetLogger(ctx)
		cluster := chi.URLParam(r, "clusterName")
		logger.Info("Deleting cluster...", "cluster", cluster)

		client, _ := multicluster.BuildMultiClusterClient(r.Context(), c.LoopbackClientConfig, "")
		err := clusterMgr.DeleteCluster(r.Context(), client, cluster)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		render.JSON(w, r, handler.SuccessResponse(ctx, "Cluster deleted"))
	}
}

// GetYAML returns an HTTP handler function that returns a cluster
// YAML. It utilizes a ClusterManager to execute the logic.
//
//	@Summary		GetYAML returns a cluster YAML by name.
//	@Description	This endpoint returns a cluster YAML by name.
//	@Tags			cluster
//	@Produce		json
//	@Success		200	{array}		byte	"Byte array"
//	@Failure		400	{string}	string	"Bad Request"
//	@Failure		401	{string}	string	"Unauthorized"
//	@Failure		404	{string}	string	"Not Found"
//	@Failure		405	{string}	string	"Method Not Allowed"
//	@Failure		429	{string}	string	"Too Many Requests"
//	@Failure		500	{string}	string	"Internal Server Error"
//	@Router			/api/v1/cluster/{clusterName}/yaml [get]
func GetYAML(clusterMgr *cluster.ClusterManager, c *server.CompletedConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the context and logger from the request.
		ctx := r.Context()
		cluster := chi.URLParam(r, "clusterName")
		client, err := multicluster.BuildMultiClusterClient(r.Context(), c.LoopbackClientConfig, "")
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		result, err := clusterMgr.GetYAMLForCluster(r.Context(), client, cluster)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		render.JSON(w, r, handler.SuccessResponse(ctx, string(result)))
	}
}

// GetTopology returns an HTTP handler function that returns a cluster
// topology. It utilizes a ClusterManager to execute the logic.
//
//	@Summary		GetTopology returns the topology of a cluster by name.
//	@Description	This endpoint returns the topology of a cluster by name.
//	@Tags			cluster
//	@Produce		json
//	@Success		200	{object}	map[string]cluster.ClusterTopology	"map from string to cluster.ClusterTopology"
//	@Failure		400	{string}	string								"Bad Request"
//	@Failure		401	{string}	string								"Unauthorized"
//	@Failure		404	{string}	string								"Not Found"
//	@Failure		405	{string}	string								"Method Not Allowed"
//	@Failure		429	{string}	string								"Too Many Requests"
//	@Failure		500	{string}	string								"Internal Server Error"
//	@Router			/api/v1/cluster/{clusterName}/topology [get]
func GetTopology(clusterMgr *cluster.ClusterManager, c *server.CompletedConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the context and logger from the request.
		ctx := r.Context()
		cluster := chi.URLParam(r, "clusterName")
		client, err := multicluster.BuildMultiClusterClient(r.Context(), c.LoopbackClientConfig, cluster)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		topologyMap, err := clusterMgr.GetTopologyForCluster(r.Context(), client, cluster)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		render.JSON(w, r, handler.SuccessResponse(ctx, topologyMap))
	}
}

// GetDetail returns an HTTP handler function that returns details of a
// cluster. It utilizes a ClusterManager to execute the logic.
//
//	@Summary		GetDetail returns the details of a cluster by name.
//	@Description	This endpoint returns the details of a cluster by name.
//	@Tags			cluster
//	@Produce		json
//	@Success		200	{object}	cluster.ClusterDetail	"Cluster detail"
//	@Failure		400	{string}	string					"Bad Request"
//	@Failure		401	{string}	string					"Unauthorized"
//	@Failure		404	{string}	string					"Not Found"
//	@Failure		405	{string}	string					"Method Not Allowed"
//	@Failure		429	{string}	string					"Too Many Requests"
//	@Failure		500	{string}	string					"Internal Server Error"
//	@Router			/api/v1/cluster/{clusterName}/detail [get]
func GetDetail(clusterMgr *cluster.ClusterManager, c *server.CompletedConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the context and logger from the request.
		ctx := r.Context()
		cluster := chi.URLParam(r, "clusterName")
		client, err := multicluster.BuildMultiClusterClient(r.Context(), c.LoopbackClientConfig, cluster)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		clusterDetail, err := clusterMgr.GetDetailsForCluster(r.Context(), client, cluster)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		render.JSON(w, r, handler.SuccessResponse(ctx, clusterDetail))
	}
}

// GetNamespace returns an HTTP handler function that reads a namespace.
// It utilizes a ClusterManager to execute the logic.
//
//	@Summary		GetNamespace returns the namespace by name.
//	@Description	This endpoint returns the namespace by name.
//	@Tags			cluster
//	@Produce		json
//	@Success		200	{object}	v1.Namespace	"v1.Namespace"
//	@Failure		400	{string}	string			"Bad Request"
//	@Failure		401	{string}	string			"Unauthorized"
//	@Failure		404	{string}	string			"Not Found"
//	@Failure		405	{string}	string			"Method Not Allowed"
//	@Failure		429	{string}	string			"Too Many Requests"
//	@Failure		500	{string}	string			"Internal Server Error"
//	@Router			/api/v1/cluster/{clusterName}/namespace/{namespaceName} [get]
func GetNamespace(clusterMgr *cluster.ClusterManager, c *server.CompletedConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the context and logger from the request.
		ctx := r.Context()
		cluster := chi.URLParam(r, "clusterName")
		namespace := chi.URLParam(r, "namespaceName")
		client, err := multicluster.BuildMultiClusterClient(r.Context(), c.LoopbackClientConfig, cluster)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		namespaceObj, err := clusterMgr.GetNamespaceForCluster(r.Context(), client, cluster, namespace)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		render.JSON(w, r, handler.SuccessResponse(ctx, namespaceObj))
	}
}

// GetNamespaceTopology returns an HTTP handler function that returns the
// topology of a namespace. It utilizes a ClusterManager to execute the logic.
//
//	@Summary		GetNamespaceTopology returns the topology of a namespace by name.
//	@Description	This endpoint returns the the topology of a namespace by name.
//	@Tags			cluster
//	@Produce		json
//	@Success		200	{object}	map[string]cluster.ClusterTopology	"map from string to cluster.ClusterTopology"
//	@Failure		400	{string}	string								"Bad Request"
//	@Failure		401	{string}	string								"Unauthorized"
//	@Failure		404	{string}	string								"Not Found"
//	@Failure		405	{string}	string								"Method Not Allowed"
//	@Failure		429	{string}	string								"Too Many Requests"
//	@Failure		500	{string}	string								"Internal Server Error"
//	@Router			/api/v1/cluster/{clusterName}/namespace/{namespaceName}/topology [get]
func GetNamespaceTopology(clusterMgr *cluster.ClusterManager, c *server.CompletedConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the context and logger from the request.
		ctx := r.Context()
		cluster := chi.URLParam(r, "clusterName")
		namespace := chi.URLParam(r, "namespaceName")
		client, err := multicluster.BuildMultiClusterClient(r.Context(), c.LoopbackClientConfig, cluster)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		topologyMap, err := clusterMgr.GetTopologyForClusterNamespace(r.Context(), client, cluster, namespace)
		if err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}
		render.JSON(w, r, handler.SuccessResponse(ctx, topologyMap))
	}
}

// @Summary		Upload kubeConfig file for cluster
// @Description	Uploads a KubeConfig file for cluster, with a maximum size of 2MB, and the valid file extension is "", ".yaml", ".yml", ".json", ".kubeconfig", ".kubeconf".
// @Tags		cluster
// @Accept		multipart/form-data
// @Produce		plain
// @Param		file	formData	file		true	"Upload file with field name 'file'"
// @Success		200		{object}	UploadData	"Returns the content of the uploaded KubeConfig file."
// @Failure		400		{string}	string		"The uploaded file is too large or the request is invalid."
// @Failure		500		{string}	string		"Internal server error."
// @Router		/api/v1/cluster/config/file [post]
func UpdateKubeConfig(w http.ResponseWriter, r *http.Request) {
	// Extract the context and logger from the request.
	ctx := r.Context()
	log := ctxutil.GetLogger(ctx)

	// Begin the update process, logging the start.
	log.Info("Starting get uploaded kubeconfig file in handler...")

	// Limit the size of the request body to prevent overflow.
	const maxUploadSize = 2 << 20
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		render.Render(w, r, handler.FailureResponse(ctx, errors.Wrapf(err, "failed to parse multipart form")))
		return
	}

	// Retrieve the file from the parsed multipart form.
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		render.Render(w, r, handler.FailureResponse(ctx, errors.Wrapf(err, "invalid request")))
		return
	}
	defer file.Close()

	// Check the file extension.
	log.Info("Uploaded filename", "filename", fileHeader.Filename)
	if !isAllowedExtension(fileHeader.Filename) {
		render.Render(w, r, handler.FailureResponse(
			ctx, errors.New("invalid file format, only '', .yaml, .yml, .json, .kubeconfig, .kubeconf are allowed.")))
		return
	}

	// Read the contents of the file.
	buf := make([]byte, maxUploadSize)
	fileSize, err := file.Read(buf)
	if err != nil && err != io.EOF {
		render.Render(w, r, handler.FailureResponse(ctx, errors.Wrapf(err, "error reading file")))
		return
	}

	// Convert the bytes read to a string and return as response.
	data := &UploadData{
		FileName: fileHeader.Filename,
		Content:  string(buf[:fileSize]),
		FileSize: fileSize,
	}
	render.JSON(w, r, handler.SuccessResponse(ctx, data))
}

// ValidateKubeConfig returns an HTTP handler function to validate a KubeConfig.
//
//	@Summary		Validate KubeConfig
//	@Description	Validates the provided KubeConfig using cluster manager methods.
//	@Tags			cluster
//	@Accept			plain
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ValidatePayload	true	"KubeConfig payload to validate"
//	@Success		200		{string}	string			"Verification passed server version"
//	@Failure		400		{object}	string			"Bad Request"
//	@Failure		401		{object}	string			"Unauthorized"
//	@Failure		429		{object}	string			"Too Many Requests"
//	@Failure		404		{object}	string			"Not Found"
//	@Failure		500		{object}	string			"Internal Server Error"
//	@Router			/api/v1/cluster/config/validate [post]
func ValidateKubeConfig(clusterMgr *cluster.ClusterManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the context and logger from the request.
		ctx := r.Context()
		log := ctxutil.GetLogger(ctx)

		// Begin the auditing process, logging the start.
		log.Info("Starting validate kubeconfig in handler...")

		// Decode the request body into the payload.
		payload := &ValidatePayload{}
		if err := payload.Decode(r); err != nil {
			render.Render(w, r, handler.FailureResponse(ctx, err))
			return
		}

		// Log successful decoding of the request body.
		sanitizeConfig, _ := clusterMgr.SanitizeKubeConfigWithYAML(ctx, payload.KubeConfig)
		log.Info("Successfully decoded the request body to payload, and sanitize the kubeconfig in payload",
			"sanitizeKubeConfig", sanitizeConfig)

		// Validate the specified kube config.
		if info, err := clusterMgr.ValidateKubeConfigWithYAML(ctx, payload.KubeConfig); err == nil {
			render.JSON(w, r, handler.SuccessResponse(ctx, info))
		} else {
			render.Render(w, r, handler.FailureResponse(ctx, err))
		}
	}
}

// isAllowedExtension checks if the provided file name has a permitted extension.
func isAllowedExtension(filename string) bool {
	allowedExtensions := []string{"", ".yaml", ".yml", ".json", ".kubeconfig", ".kubeconf"}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			return true
		}
	}
	return false
}