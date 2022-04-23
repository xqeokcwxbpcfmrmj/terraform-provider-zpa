package inspection_custom_controls

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/willguibr/terraform-provider-zpa/gozscaler/common"
)

const (
	mgmtConfig             = "/mgmtconfig/v1/admin/customers/"
	customControlsEndpoint = "/inspectionControls/custom"
)

type InspectionCustomControl struct {
	ID                               string                   `json:"id,omitempty"`
	Action                           string                   `json:"action,omitempty"`
	ActionValue                      string                   `json:"actionValue,omitempty"`
	AssociatedInspectionProfileNames []AssociatedProfileNames `json:"associatedInspectionProfileNames,omitempty"`
	Rules                            []common.Rules           `json:"rules,omitempty"`
	ControlNumber                    string                   `json:"controlNumber,omitempty"`
	ControlRuleJson                  string                   `json:"controlRuleJson,omitempty"`
	CreationTime                     string                   `json:"creationTime,omitempty"`
	DefaultAction                    string                   `json:"defaultAction,omitempty"`
	DefaultActionValue               string                   `json:"defaultActionValue,omitempty"`
	Description                      string                   `json:"description,omitempty"`
	ModifiedBy                       string                   `json:"modifiedBy,omitempty"`
	ModifiedTime                     string                   `json:"modifiedTime,omitempty"`
	Name                             string                   `json:"name,omitempty"`
	ParanoiaLevel                    string                   `json:"paranoiaLevel,omitempty"`
	Severity                         string                   `json:"severity,omitempty"`
	Type                             string                   `json:"type,omitempty"`
	Version                          string                   `json:"version,omitempty"`
}

type AssociatedProfileNames struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (service *Service) Get(customID string) (*InspectionCustomControl, *http.Response, error) {
	v := new(InspectionCustomControl)
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+customControlsEndpoint, customID)
	resp, err := service.Client.NewRequestDo("GET", relativeURL, nil, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) GetByName(controlName string) (*InspectionCustomControl, *http.Response, error) {
	var v struct {
		List []InspectionCustomControl `json:"list"`
	}

	relativeURL := mgmtConfig + service.Client.Config.CustomerID + customControlsEndpoint
	resp, err := service.Client.NewRequestDo("GET", relativeURL, common.Pagination{PageSize: common.DefaultPageSize, Search: controlName}, nil, &v)
	if err != nil {
		return nil, nil, err
	}
	for _, control := range v.List {
		if strings.EqualFold(control.Name, controlName) {
			return &control, resp, nil
		}
	}
	return nil, resp, fmt.Errorf("no inspection profile named '%s' was found", controlName)
}

func (service *Service) Create(customControls InspectionCustomControl) (*InspectionCustomControl, *http.Response, error) {
	v := new(InspectionCustomControl)
	resp, err := service.Client.NewRequestDo("POST", mgmtConfig+service.Client.Config.CustomerID+customControlsEndpoint, nil, customControls, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

func (service *Service) Update(customID string, customControls *InspectionCustomControl) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+customControlsEndpoint, customID)
	resp, err := service.Client.NewRequestDo("PUT", relativeURL, nil, customControls, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (service *Service) Delete(customID string) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s", mgmtConfig+service.Client.Config.CustomerID+customControlsEndpoint, customID)
	resp, err := service.Client.NewRequestDo("DELETE", relativeURL, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
