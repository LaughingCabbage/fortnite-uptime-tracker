package main

type ServerStatus struct {
	ServiceInstanceID  string       `json:"serviceInstanceId"`
	Status             string       `json:"status"`
	Message            string       `json:"message"`
	MaintenenceURL     string       `json:"maintenenceUrl"`
	OverrideCatalogIDs []string     `json:"overrideCatalogIds"`
	AllowedActions     []string     `json:"allowedActions"`
	Banned             bool         `json:"banned"`
	LauncherInfo       LauncherInfo `json:"launcherInfoDTO"`
}

type LauncherInfo struct {
	AppName       string `json:"appName"`
	CatalogItemID string `json:"catalogItemId"`
	Namespace     string `json:"namespace"`
}
