package rustsnitch


type Group struct {
	Name string
	Count int
	Desc string
	Members map[string]bool  //TODO: change to player data type?
}