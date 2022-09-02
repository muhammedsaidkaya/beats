package add_docker_metadata

import (
 "fmt"
 "github.com/elastic/beats/v7/libbeat/common"
 "github.com/elastic/beats/v7/libbeat/common/docker"
 "strings"
)

func picus(d *addDockerMetadata, container *docker.Container, meta *common.MapStr, environment string) {
 Block{
  Try: generateCustomMetadata,
  Catch: func(e Exception) {
   d.log.Debugf("Caught %v\n", e)
  },
  Finally: func(environmentName string) {
   //Environment
   meta.Put("container.environment.name", environmentName)
  },
 }.Do(d, container, meta, environment)
}

func generateCustomMetadata(d *addDockerMetadata, container *docker.Container, meta *common.MapStr, environment string) {

 cacheValue, found := customCache.GetCache(container.ID)
 if found == false {
  d.log.Debugf("Running custom script", container.ID)

  //Get image tags from ECR
  err, tagsObject, _ := executeScript(fmt.Sprintf("/etc/filebeat/get_image_tag.sh %v", container.Image))

  if err != nil {
   d.log.Errorf("Error while running custom script: %v", err)
  } else {

   //Tags
   tagsMap := jsonToMap(tagsObject)
   if val, ok := tagsMap["result"]; ok {
    //do something here
    tags := val.([]interface{})
    _, err = meta.Put("container.environment.tags", tags)

    if err != nil {
     d.log.Errorf("Error while putting custom labels: %v", err)
    } else {
     //Sha
     sha := strings.Split(strings.Split(tags[0].(string), environment)[1], "-")[1]
     d.log.Debugf("Running custom script - Sha: %v", sha)
     meta.Put("container.environment.sha", sha)

     //Caching
     cacheValue := common.MapStr{
      "tags": tags,
      "sha":  sha,
     }
     customCache.SetCache(container.ID, cacheValue)
     d.log.Debugf("Put object to the cache %v", cacheValue)
    }
   }

  }
 } else {
  d.log.Debugf("Getting from cache cid=%v", container.ID)
  meta.Put("container.environment", cacheValue)
 }
}
