package cmbchina

type DCLISMODResponse struct {
	Ntqmdlstz []struct {
		Busmod string `json:"busmod"`
		Modals string `json:"modals"`
	} `json:"ntqmdlstz"`
}
