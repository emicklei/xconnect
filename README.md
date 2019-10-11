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

- description of what a service can provide by connecting to it, the *listen* part
- decscription of what a service needs to consume or produce to using a connection, the *connect* part
- metadata about the service such as version, ownership and custom labels.

Not all application configuration is related to connectivity ; the section for connectivity can be part of it.
So instead of keeping connection related information separate from the rest of the configuration, with the inevitable effect of becoming out of date, the xconnect section should be integrated in the complete configuration.
The actual format of the xconnect data is free-form YAML, meaning that users are free to add their own service metadata if desired.

## Sprint Boot application configration

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

## Inspiration

- https://dzone.com/articles/cataloguing-microservices