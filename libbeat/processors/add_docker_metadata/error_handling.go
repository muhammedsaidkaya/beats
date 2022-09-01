package add_docker_metadata

import (
 "github.com/elastic/beats/v7/libbeat/common"
 "github.com/elastic/beats/v7/libbeat/common/docker"
)

type Block struct {
 Try     func(d *addDockerMetadata, container *docker.Container, meta *common.MapStr, environment string)
 Catch   func(Exception)
 Finally func(environment string)
}

type Exception interface{}

func Throw(up Exception) {
 panic(up)
}

func (tcf Block) Do(d *addDockerMetadata, container *docker.Container, meta *common.MapStr, environment string) {
 if tcf.Finally != nil {

  defer tcf.Finally(environment)
 }
 if tcf.Catch != nil {
  defer func() {
   if r := recover(); r != nil {
    tcf.Catch(r)
   }
  }()
 }
 tcf.Try(d, container, meta, environment)
}
