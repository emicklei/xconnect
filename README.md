# xconnect - declaring microservices connectivity

[![Build Status](https://travis-ci.org/emicklei/xconnect.png)](https://travis-ci.org/emicklei/xconnect)

Xconnect is a structure in YAML that describes how a service can accept connections and what connections a service needs to operate.
With this information, a static overview of the landscape of intercconnected services can be constructed.

[specification in YAML](https://raw.githubusercontent.com/emicklei/xconnect/master/spec-xconnect.yaml)

## how does it work

Every application/service uses some kind of configuration to specifiy what other services is connects to.
This can be a database, another service, a filesystem etc.
By standardising a part of that configuration, a tool can extract connectivity information from that configuration.
This is called the xconnect section.

## xonnect section

The xconnect section of a configuration (now YAML only), consists of 3 parts:

- description of what a service can provide by connecting to it, the *listen* part.
- description of what a service needs to consume or produce to using a connection, the *connect* part
- metadata about the service such as version, ownership and custom labels.

Not all application configuration is related to connectivity ; the section for connectivity can be part of it.
So instead of keeping connection related information separate from the rest of the configuration, with the inevitable effect of becoming out of date, the xconnect section should be integrated in the complete configuration.
The actual format of the xconnect data is free-form YAML, meaning that users are free to add their own service metadata if desired.

## Use as Go package

This example uses *gopkg.in/yaml.v2* for parsing the configuration.

    content, err := ioutil.ReadFile("your-app.yaml")
    var doc xonnect.Document
    err := yaml.Unmarshal(content, &doc)

## Sprint Boot application configration

    xconnect:
      meta: 
      listen:
      connect:
        some-db:
          url: jdbc:postgresql://localhost:5432/postgres?reWriteBatchedInserts=true
 
    spring:
      datasource:
        # use a reference to the actual value, within this document
        url: ${xconnect.connect.some-db.url}

To extract the xconnect section using the command line tool:

    xconnect -input application.yml -target file://xconnect-from-application.yml

or POST the extracted information to an HTTP endpoint

    xconnect -input application.yml -target http://some-service/v1/xconnect

## Kubernetes configration (ConfigMap)

    apiVersion: v1
    data:
        xconnect:
            meta: 
                ...
            listen:
                ...
            connect:
                ...
            
    kind: ConfigMap
    metadata:
        creationTimestamp: null
        name: some
        namespace: somewhere

To extract the xconnect section using the command line tool:

    xconnect -input configmap.yml -k8s -target file://xconnect-from-configmap.yml

or using the Go package:

    var k8s xconnect.K8SConfiguration
    if err = yaml.Unmarshal(yamlContentBytes, &k8s); err != nil {
        return
    }
    cfg, err := k8s.ExtractConfig()
    ...

## Getting the extra fields

See xconnect_test.go

## Inspiration

- https://dzone.com/articles/cataloguing-microservices


© 2019, ernestmicklei.com. MIT License. Contributions welcome.