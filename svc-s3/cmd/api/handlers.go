package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gitlab.amin.run/general/project/subs-mgmt/svc-s3/internal/instance"
	"gitlab.amin.run/general/project/subs-mgmt/svc-s3/internal/plan"
)

func (app *Config) svcS3(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "you hit the s3 service",
	}

	out, _ := json.MarshalIndent(payload, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}

func (app *Config) svcS32(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "you hit the s32 service",
	}

	out, _ := json.MarshalIndent(payload, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}

// CreatePlanHandler handles the creation of a new plan
func (app *Config) CreatePlanHandler(w http.ResponseWriter, r *http.Request) {
	var plan plan.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdPlan, err := app.PlanService.CreatePlan(context.Background(), &plan)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating plan: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPlan)
}

// GetPlanHandler handles retrieving a plan by its ID
func (app *Config) GetPlanHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	planID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid plan ID", http.StatusBadRequest)
		return
	}

	plan, err := app.PlanService.GetPlan(context.Background(), planID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting plan: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plan)
}

// UpdatePlanHandler handles updating an existing plan
func (app *Config) UpdatePlanHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	planID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid plan ID", http.StatusBadRequest)
		return
	}

	var plan plan.Plan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	plan.ID = planID
	updatedPlan, err := app.PlanService.UpdatePlan(context.Background(), &plan)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating plan: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedPlan)
}

// DeletePlanHandler handles deleting a plan by its ID
func (app *Config) DeletePlanHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	planID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid plan ID", http.StatusBadRequest)
		return
	}

	if err := app.PlanService.DeletePlan(context.Background(), planID); err != nil {
		http.Error(w, fmt.Sprintf("Error deleting plan: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllPlansHandler handles retrieving all plans
func (app *Config) GetAllPlansHandler(w http.ResponseWriter, r *http.Request) {
	plans, err := app.PlanService.GetAllPlans(context.Background())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving plans: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plans)
}



// CreateInstanceHandler handles the creation of a new instance
func (app *Config) CreateInstanceHandler(w http.ResponseWriter, r *http.Request) {
    var instance instance.Instance
    if err := json.NewDecoder(r.Body).Decode(&instance); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    createdInstance, err := app.InstanceService.CreateInstance(context.Background(), &instance)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error creating instance: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(createdInstance)
}

// GetInstanceHandler handles retrieving an instance by its ID
func (app *Config) GetInstanceHandler(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    instanceID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid instance ID", http.StatusBadRequest)
        return
    }

    instance, err := app.InstanceService.GetInstance(context.Background(), instanceID)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error getting instance: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(instance)
}

// UpdateInstanceHandler handles updating an existing instance
func (app *Config) UpdateInstanceHandler(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    instanceID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid instance ID", http.StatusBadRequest)
        return
    }

    var instance instance.Instance
    if err := json.NewDecoder(r.Body).Decode(&instance); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    instance.ID = instanceID
    updatedInstance, err := app.InstanceService.UpdateInstance(context.Background(), &instance)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error updating instance: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedInstance)
}

// DeleteInstanceHandler handles deleting an instance by its ID
func (app *Config) DeleteInstanceHandler(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    instanceID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid instance ID", http.StatusBadRequest)
        return
    }

    if err := app.InstanceService.DeleteInstance(context.Background(), instanceID); err != nil {
        http.Error(w, fmt.Sprintf("Error deleting instance: %v", err), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// GetAllInstancesHandler handles retrieving all instances
func (app *Config) GetAllInstancesHandler(w http.ResponseWriter, r *http.Request) {
    instances, err := app.InstanceService.GetAllInstances(context.Background())
    if err != nil {
        http.Error(w, fmt.Sprintf("Error retrieving instances: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(instances)
}
