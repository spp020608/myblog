package dto


type RegisterParam struct {
	Username string `json:"username" form:"username" binding:"required,min=1,max=20"`
	Password string `json:"password" form:"password"binding:"required,min=1,max=128"`
	Tel      string `json:"tel" form:"tel" binding:"alphanum,len=11"`
}
