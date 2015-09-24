package carton

import (
	log "github.com/Sirupsen/logrus"
	"github.com/megamsys/gulp/carton/bind"
	"github.com/megamsys/gulp/provision"
	"gopkg.in/yaml.v2"
	"github.com/megamsys/gulp/repository"
)

type BoxLevel int

const (
	// BoxAny indicates that there is atleast one box to deploy or delete.
	BoxAny BoxLevel = iota

	// BoxZero indicates that there are no boxes to deploy or delete.
	BoxZero
)

type Carton struct {
	Id         string //assemblyid
	Name       string
	CartonsId  string
	Tosca      string
//	Compute    provision.BoxCompute
	Repo       repository.Repo
	DomainName string
	Provider   string
	Envs       []bind.EnvVar
	Boxes      *[]provision.Box
}


func (a *Carton) String() string {
	if d, err := yaml.Marshal(a); err != nil {
		return err.Error()
	} else {
		return string(d)
	}
}

//If there are boxes, then it set the enum BoxSome or its BoxZero
func (c *Carton) lvl() provision.BoxLevel {
	if len(*c.Boxes) > 0 {
		return provision.BoxSome
	} else {
		return provision.BoxNone
	}
}

//Converts a carton to a box, if there are no boxes below.
func (c *Carton) toBox() error { //assemblies id.
	switch c.lvl() {
	case provision.BoxNone:
		c.Boxes = &[]provision.Box{provision.Box{
			CartonId:   c.Id,    //this isn't needed.
			Id:         c.Id,    //assembly id sent in ContextMap
			CartonsId:  c.CartonsId,      //assembliesId,
			Level:      c.lvl(), //based on the level, we decide to use the Box-Id as ComponentId or AssemblyId
			Name:       c.Name,
			DomainName: c.DomainName,
	//		Compute:    c.Compute,
			Repo:       c.Repo,
			Provider:   c.Provider,
			Tosca:      c.Tosca,
		},
		}
	}
	return nil
}

/*func (c *Carton) Create() error {
	for _, box := range *c.Boxes {
		err := Create(&DeployOpts{B: &box})
		if err != nil {
			log.Errorf("Unable to deploy box", err)
		}
	}
	return nil
}*/

// moves the state to the desired state
// changing the boxes state to StatusStateup.
func (c *Carton) Stateup() error {
	for _, box := range *c.Boxes {
		err := Deploy(&DeployOpts{B: &box})
		if err != nil {
			log.Errorf("Unable to deploy box", err)
		}
	}
	return nil
}

// moves the state down to the desired state
// changing the boxes state to StatusStatedown.
func (c *Carton) Statedown() error {
	return nil
}

func (c *Carton) Delete() error {
	return nil
}

func (c *Carton) Start() error {
	return nil
}

func (c *Carton) Stop() error {
	return nil
}

func (c *Carton) Restart() error {
	return nil
}

// GetTosca returns the tosca type  of the carton.
func (c *Carton) GetTosca() string {
	return c.Tosca
}


// Envs returns a map representing the apps environment variables.
func (c *Carton) GetEnvs() []bind.EnvVar {
	return c.Envs
}
