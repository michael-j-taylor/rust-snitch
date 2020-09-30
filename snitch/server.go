package rustsnitch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)


//Server stores a Rust server and the desired groups/ players to track
type Server struct {
	
	//string of format ip:port
	Address string

	//name of server
	Name string

	//slice of tracked groups on server
	Groups []Group

	//slice of map of player: is_online values
	Players []map[string]bool
}


//newServer creates a new Server struct initialized with an address and a name
func newServer(addr, name string) *Server {
	s := new(Server)
	s.Address = addr
	s.Name = name
	return s
}


//TODO: ensure no duplicate groups can exist in Server struct
//addGroup appends a group to the Group field of a Server struct
func (s *Server) addGroup (g *Group) {
	s.Groups = append(s.Groups, *g)
	return
}


//removeGroup removes a group from the Group field of a Server struct
//because the group to be removed is assumed to already exist
//the function takes a string as an argument instead of a Group
func (s *Server) removeGroup(name string) {
	tmp := s.Groups[:0]
	for _, g := range(s.Groups) {
		if g.Name != name {
			tmp = append(tmp, g)  //append group to tmp if it's not being looked for
		}
	}

	s.Groups = tmp
	return
}


//addPlayer adds a <TYPE> to the Players field of a Server struct

//removePlayer removes a <TYPE> from the Players field of a Server struct



