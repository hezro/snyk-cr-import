package main

import (
  "bytes"
  "encoding/json"
  "io/ioutil"
  "log"
  "net/http"
  "strconv"
  "time"

  "gopkg.in/alecthomas/kingpin.v2"
)

const base_url = "https://api.snyk.io/api/v1"

// Define flags for the utility.
var (
  snyk_token = kingpin.Flag("token", "Snyk API Token").Required().String() 
  cr_integration_id = kingpin.Flag("crId", "Container Registry ID").Required().String() 
  org_id = kingpin.Flag("orgId", "Snyk Target Organization ID").Required().String() 
  image_name = kingpin.Flag("imageName", "Name of the container image including the tag: hezro/juice-shop:1.1.0").Required().String()
)

// Define the structs for the json fields.
type DataAttributes struct {
  Target Target `json:"target"`
}
type Target struct {
  Name string `json:"name"`
}

// Set up the http client with a 10 second timeout.
func httpClient() * http.Client {
  client := & http.Client {
    Timeout: 10 * time.Second
  }
  return client
}

// Import the container image into Snyk.
func image_import(client * http.Client, snyk_token string, cr_integration_id string, org_id string, image_name string)[] byte {
  url := base_url + "/org/" + org_id + "/integrations/" + cr_integration_id + "/import"

  // Build the json data.
  target := DataAttributes {}
  target.Target.Name = image_name

  json_data, _ := json.Marshal(target)
  body := bytes.NewBuffer(json_data)

  // Send the POST request
  req, err := http.NewRequest("POST", url, body)
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", snyk_token)

  if err != nil {
    log.Fatalf("| Error | Failed: %+v", err)
  }

  // Send the POST request
  response, err := client.Do(req)
  if err != nil {
    log.Fatalf("| Error | Failed sending request to API endpoint. |  %+v \n", err)
  }
  defer response.Body.Close()

  response_body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatalf("| Error | Could not parse response body. |  %+v \n", err)
  }

  status_code := response.StatusCode
  if status_code != 201 {
    log.Printf("| Error |  Import failed to POST | Status Code: %s \n", strconv.Itoa((status_code)))
  }

  return response_body
}

func main() {
  // Parse command line arguments
  kingpin.Parse()

  // Create the http client
  client := httpClient()

  // Start the import
  import_projects: = image_import(client, * snyk_token, * cr_integration_id, * org_id, * image_name)

  // Report if the POST import was successful
  body_len := len(import_projects)
  if body_len == 2 {
    log.Printf("| Success | POST request submitted successfully items from %s to orgId: %s\n", * image_name, * org_id)
  } else {
    log.Printf("| Error |  Import failed to POST | Body: %s \n", import_projects)
  }
}
