package system

type MeunResultCrm struct {
	Data    []*MeunCrm `json:"data"`
	Success bool       `json:"success"`
}

type MetaCrm struct {
	//Auths []string `fmt:"auths"`
	//Roles []string `fmt:"roles"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Rank  int64  `json:"rank"`
}

type MeunCrm struct {
	Path     string      `json:"path"`
	Name     string      `json:"name"`
	Meta     MetaCrm     `json:"meta"`
	Children *[]*MeunCrm `json:"children,omitempty"`
}

type SysCrmMenusResponse struct {
	Menus *[]*MeunCrm `json:"menus"`
}
