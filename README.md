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
- metadata about the service such as version, opex (ownership) and custom tags.

Not all application configuration is related to connectivity ; the section for connectivity can be part of it (embedding).
So instead of keeping connection related information separate from the rest of the configuration, with the inevitable effect of becoming out of date, the xconnect section should be integrated in the complete configuration.
The actual format of the xconnect data is free-form YAML, meaning that users are free to add their own service metadata if desired.

## Example

    doc-field: doc-value

    xconnect:
      extra-field: extra-value
      
      meta:
        # name for discovery
        name: account-service
        # tagged version of the implementation
        version: v1.2.3
        # team that owns the code and operates it
        opex: team accounts
        tags:
          - account
          - registration
          - search    
      listen:
        api:
          host: localhost
          protocol: grpc
          port: 9443
        web:
          host: localhost
          protocol: http
          secure: true
          port: 443
      connect:
        some-db:
          kind: db
          url: jdbc:postgresql://localhost:5432/postgres?reWriteBatchedInserts=true
        some-cache:
          host: #REDIS_IP
          port: 6379
          kind: db
        variant-publish:
          kind: gcp.pubsub
          resource: VariantToAssortment_Push_v1-topic          
        variant-pull:
          kind: gcp.pubsub:
          resource: subscription: Variant_v1-subscription
            test:
              topic: Variant_v1-topic

## Use as Go package

  import (
    "github.com/emicklei/xconnect"
  )

This example uses *gopkg.in/yaml.v2* for parsing the configuration.

    content, err := ioutil.ReadFile("your-app.yaml")
    var doc xconnect.Document
    err := yaml.Unmarshal(content, &doc)

    version := doc.XConnect.Meta.Version
    // alternative: doc.FindString("xconnect/meta/version")
    
    webPort := doc.Xconnect.Listen["web"].Port
    // alternative: doc.FindInt("xconnect/listen/web/port")
    
    variantPullTestTopic := doc.FindString("xconnect/connect/variant-pull/resource/test/topic")

## Sprint Boot application configration

A Spring configuration needs it own root element in a YAML file.
To use an xconnect section in this file, relevant property values should be referenced using the ${..} notation supported by Spring.

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

### extract

To extract the xconnect section using the command line tool:

    xconnect -input application.yml -target file://xconnect-from-application.yml

## Kubernetes configration (ConfigMap)

A Kubernetes configuration as its defined structure in a YAML file.
To use an xconnect section in this file, its content must be inside the `data` field.

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

### extract

To extract the xconnect section using the command line tool:

    xconnect -input configmap.yml -k8s -target file://xconnect-from-configmap.yml

## view

  xconnect -dot | dot -Tpng  > graph.png && open graph.png

## Getting the extra fields

See xconnect_test.go

## Inspiration

- https://dzone.com/articles/cataloguing-microservices


© 2019+, [ernestmicklei.com](http://ernestmicklei.com). MIT License. Contributions welcome.
