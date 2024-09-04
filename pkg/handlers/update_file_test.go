package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/artmoskvin/hide/pkg/files"
	"github.com/artmoskvin/hide/pkg/handlers"
	"github.com/artmoskvin/hide/pkg/model"
	"github.com/artmoskvin/hide/pkg/project"
	project_mocks "github.com/artmoskvin/hide/pkg/project/mocks"
)

func TestUpdateFileHandler_Success(t *testing.T) {
	tests := []struct {
		name    string
		payload handlers.UpdateFileRequest
		want    model.File
	}{
		{
			name: "Update file with udiff",
			payload: handlers.UpdateFileRequest{
				Type: handlers.Udiff,
				Udiff: &handlers.UdiffRequest{
					Patch: "--- test.txt\n+++ test.txt\n@@ -1,3 +1,3 @@\n line1\n-line2\n+line20\n line3\n",
				},
			},
			want: model.File{
				Path: "test.txt",
				Lines: []model.Line{
					{Number: 1, Content: "line1"},
					{Number: 2, Content: "line20"},
					{Number: 3, Content: "line3"},
				},
			},
		},
		{
			name: "Update file with linediff",
			payload: handlers.UpdateFileRequest{
				Type: handlers.LineDiff,
				LineDiff: &handlers.LineDiffRequest{
					StartLine: 1,
					EndLine:   3,
					Content:   "line11\nline12\n",
				},
			},
			want: model.File{
				Path: "test.txt",
				Lines: []model.Line{
					{Number: 1, Content: "line11"},
					{Number: 2, Content: "line12"},
					{Number: 3, Content: "line3"},
				},
			},
		},
		{
			name: "Update file with overwrite",
			payload: handlers.UpdateFileRequest{
				Type: handlers.Overwrite,
				Overwrite: &handlers.OverwriteRequest{
					Content: "line1\nline2\nline3\n",
				},
			},
			want: model.File{
				Path: "test.txt",
				Lines: []model.Line{
					{Number: 1, Content: "line1"},
					{Number: 2, Content: "line2"},
					{Number: 3, Content: "line3"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockManager := &project_mocks.MockProjectManager{
				ApplyPatchFunc: func(ctx context.Context, projectId string, path, patch string) (*model.File, error) {
					return &tt.want, nil
				},
				UpdateLinesFunc: func(ctx context.Context, projectId string, path string, lineDiff files.LineDiffChunk) (*model.File, error) {
					return &tt.want, nil
				},
				UpdateFileFunc: func(ctx context.Context, projectId string, path, content string) (*model.File, error) {
					return &tt.want, nil
				},
			}

			handler := handlers.UpdateFileHandler{ProjectManager: mockManager}
			router := handlers.NewRouter().WithUpdateFileHandler(handler).Build()

			payload, _ := json.Marshal(tt.payload)
			request, _ := http.NewRequest(http.MethodPut, "/projects/123/files/test.txt", bytes.NewBuffer(payload))
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			if response.Code != http.StatusOK {
				t.Errorf("want status %d, got %d", http.StatusOK, response.Code)
			}

			var actual model.File
			if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}

			if !actual.Equals(&tt.want) {
				t.Errorf("want file %+v, got %+v", tt.want, actual)
			}
		})
	}
}

func TestUpdateFileHandler_RespondsWithBadRequest_IfRequestIsUnparsable(t *testing.T) {
	handler := handlers.UpdateFileHandler{}
	router := handlers.NewRouter().WithUpdateFileHandler(handler).Build()

	request, _ := http.NewRequest(http.MethodPut, "/projects/123/files/test.txt", bytes.NewBuffer([]byte("invalid json")))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Errorf("want status %d, got %d", http.StatusBadRequest, response.Code)
	}

	if !strings.Contains(strings.ToLower(response.Body.String()), "failed parsing request body") {
		t.Errorf("want error message 'Failed parsing request body', got %s", response.Body.String())
	}
}

func TestUpdateFileHandler_RespondsWithBadRequest_IfRequestIsInvalid(t *testing.T) {
	tests := []struct {
		name    string
		payload handlers.UpdateFileRequest
		message string
	}{
		{
			name: "Update type is missing",
			payload: handlers.UpdateFileRequest{
				Udiff: &handlers.UdiffRequest{
					Patch: "--- test.txt\n+++ test.txt\n@@ -1,3 +1,3 @@\n line1\n-line2\n+line20\n line3\n",
				},
			},
			message: "type must be provided",
		},
		{
			name: "Update type is invalid",
			payload: handlers.UpdateFileRequest{
				Type: "invalid",
				Udiff: &handlers.UdiffRequest{
					Patch: "--- test.txt\n+++ test.txt\n@@ -1,3 +1,3 @@\n line1\n-line2\n+line20\n line3\n",
				},
			},
			message: "invalid type: invalid",
		},
		{
			name: "Udiff is missing when type is udiff",
			payload: handlers.UpdateFileRequest{
				Type: handlers.Udiff,
				LineDiff: &handlers.LineDiffRequest{
					StartLine: 1,
					EndLine:   3,
					Content:   "line11\nline12\n",
				},
			},
			message: "udiff must be provided",
		},
		{
			name: "LineDiff is missing when type is linediff",
			payload: handlers.UpdateFileRequest{
				Type: handlers.LineDiff,
				Udiff: &handlers.UdiffRequest{
					Patch: "--- test.txt\n+++ test.txt\n@@ -1,3 +1,3 @@\n line1\n-line2\n+line20\n line3\n",
				},
			},
			message: "linediff must be provided",
		},
		{
			name: "Overwrite is missing when type is overwrite",
			payload: handlers.UpdateFileRequest{
				Type: handlers.Overwrite,
				Udiff: &handlers.UdiffRequest{
					Patch: "--- test.txt\n+++ test.txt\n@@ -1,3 +1,3 @@\n line1\n-line2\n+line20\n line3\n",
				},
			},
			message: "overwrite must be provided",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := handlers.UpdateFileHandler{}
			router := handlers.NewRouter().WithUpdateFileHandler(handler).Build()

			body, _ := json.Marshal(tt.payload)
			request, _ := http.NewRequest(http.MethodPut, "/projects/123/files/test.txt", bytes.NewBuffer(body))
			response := httptest.NewRecorder()

			router.ServeHTTP(response, request)

			if response.Code != http.StatusBadRequest {
				t.Errorf("want status %d, got %d", http.StatusBadRequest, response.Code)
			}

			if !strings.Contains(strings.ToLower(response.Body.String()), tt.message) {
				t.Errorf("want error message '%s', got %s", tt.message, response.Body.String())
			}
		})
	}
}

func TestUpdateFileHandler_RespondsWithInternalServerError_IfFileManagerFails(t *testing.T) {
	mockManager := &project_mocks.MockProjectManager{
		ApplyPatchFunc: func(ctx context.Context, projectId string, path, patch string) (*model.File, error) {
			return nil, errors.New("file manager error")
		},
	}

	handler := handlers.UpdateFileHandler{ProjectManager: mockManager}
	router := handlers.NewRouter().WithUpdateFileHandler(handler).Build()

	body, _ := json.Marshal(handlers.UpdateFileRequest{
		Type: handlers.Udiff,
		Udiff: &handlers.UdiffRequest{
			Patch: "--- test.txt\n+++ test.txt\n@@ -1,3 +1,3 @@\n line1\n-line2\n+line20\n line3\n",
		},
	})

	request, _ := http.NewRequest(http.MethodPut, "/projects/123/files/test.txt", bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusInternalServerError {
		t.Errorf("want status %d, got %d", http.StatusInternalServerError, response.Code)
	}

	if !strings.Contains(strings.ToLower(response.Body.String()), "failed to update file") {
		t.Errorf("want error message 'Failed to update file', got %s", response.Body.String())
	}
}

func TestUpdateFileHandler_RespondsWithNotFound_IfProjectNotFound(t *testing.T) {
	mockManager := &project_mocks.MockProjectManager{
		ApplyPatchFunc: func(ctx context.Context, projectId string, path, patch string) (*model.File, error) {
			return nil, project.NewProjectNotFoundError(projectId)
		},
	}

	handler := handlers.UpdateFileHandler{ProjectManager: mockManager}
	router := handlers.NewRouter().WithUpdateFileHandler(handler).Build()

	body, _ := json.Marshal(handlers.UpdateFileRequest{
		Type: handlers.Udiff,
		Udiff: &handlers.UdiffRequest{
			Patch: "--- test.txt\n+++ test.txt\n@@ -1,3 +1,3 @@\n line1\n-line2\n+line20\n line3\n",
		},
	})

	request, _ := http.NewRequest(http.MethodPut, "/projects/123/files/test.txt", bytes.NewBuffer(body))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("want status %d, got %d", http.StatusNotFound, response.Code)
	}

	if !strings.Contains(strings.ToLower(response.Body.String()), "not found") {
		t.Errorf("want error message 'Project not found', got %s", response.Body.String())
	}
}

func TestPathStartingWithSlash(t *testing.T) {
	t.Run("Update file with invalid path should return 400", func(t *testing.T) {
		// Setup
		mockManager := &project_mocks.MockProjectManager{
			UpdateLinesFunc: func(ctx context.Context, projectId string, path string, lineDiff files.LineDiffChunk) (*model.File, error) {
				return nil, nil
			},
		}

		handler := handlers.UpdateFileHandler{ProjectManager: mockManager}
		router := handlers.NewRouter().WithUpdateFileHandler(handler).Build()

		payload, _ := json.Marshal(handlers.UpdateFileRequest{
			Type: handlers.LineDiff,
			LineDiff: &handlers.LineDiffRequest{
				StartLine: 1,
				EndLine:   3,
				Content:   "line11\nline12\n",
			},
		})

		request, _ := http.NewRequest(http.MethodPut, "/projects/123/files//test.txt", bytes.NewBuffer(payload))
		response := httptest.NewRecorder()

		// Execute
		router.ServeHTTP(response, request)

		// Verify
		if response.Code != http.StatusBadRequest {
			t.Errorf("want status %d, got %d", http.StatusBadRequest, response.Code)
			t.Errorf("Body: %s", response.Body.String())
		}
	})
}

func TestEmptyPath(t *testing.T) {
	t.Run("Update file with invalid path should return 400", func(t *testing.T) {
		// Setup
		mockManager := &project_mocks.MockProjectManager{
			UpdateLinesFunc: func(ctx context.Context, projectId string, path string, lineDiff files.LineDiffChunk) (*model.File, error) {
				return nil, nil
			},
		}

		handler := handlers.UpdateFileHandler{ProjectManager: mockManager}
		router := handlers.NewRouter().WithUpdateFileHandler(handler).Build()

		payload, _ := json.Marshal(handlers.UpdateFileRequest{
			Type: handlers.LineDiff,
			LineDiff: &handlers.LineDiffRequest{
				StartLine: 1,
				EndLine:   3,
				Content:   "line11\nline12\n",
			},
		})

		request, _ := http.NewRequest(http.MethodPut, "/projects/123/files/", bytes.NewBuffer(payload))
		response := httptest.NewRecorder()

		// Execute
		router.ServeHTTP(response, request)

		// Verify
		if response.Code != http.StatusBadRequest {
			t.Errorf("want status %d, got %d", http.StatusBadRequest, response.Code)
			t.Errorf("Body: %s", response.Body.String())
		}
	})
}
